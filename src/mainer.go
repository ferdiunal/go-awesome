package src

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type AwesomeRepositoryInterface interface {
	Run()
	FetchingRepo(category Category)
}

type AwesomeRepository struct {
	URL     string
	Github  *github.Client
	Context context.Context
}

type Category struct {
	Title string
	Url   string
}

type Repo struct {
	Title            string
	Description      string
	Url              string
	ForksCount       int
	NetworkCount     int
	OpenIssuesCount  int
	StargazersCount  int
	SubscribersCount int
	WatchersCount    int
}

type Repositories struct {
	Title  string
	Blocks []*Repo
}

var categories []Category
var getRepositories []*Repositories

func (m *AwesomeRepository) Run() {
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		e.ForEach("ul", func(i int, h *colly.HTMLElement) {
			if i == 0 {
				h.ForEach("li", func(ii int, l *colly.HTMLElement) {
					if ii == 0 {
						l.ForEach("ul", func(iii int, ul *colly.HTMLElement) {
							ul.ForEach("li", func(iiii int, li *colly.HTMLElement) {
								title := strings.Trim(li.DOM.Find("a").Text(), " ")
								url := e.Request.AbsoluteURL(li.DOM.Find("a").AttrOr("href", ""))
								for _, category := range categories {
									if category.Title == title {
										return
									}
								}
								categories = append(categories, Category{
									Title: title,
									Url:   url,
								})

							})
						})
					}
				})
			}
		})
	})

	c.Visit(m.URL)

	c.Wait()

	for _, category := range categories {
		m.FetchingRepo(category)
		time.Sleep(time.Second * 1)
	}

	markdownFileData := ""

	for _, repositories := range getRepositories {
		title := Title(repositories.Title)
		tableHeader := TableHeader()
		tableBody := ""
		for _, repo := range repositories.Blocks {
			tableBody += TableRow(repo)
		}

		markdownFileData += fmt.Sprintf("%s\n%s%s\n\n", title, tableHeader, tableBody)
	}

	generateMarkdown([]byte(markdownFileData))
}

func (m *AwesomeRepository) FetchingRepo(category Category) {
	fmt.Println(category.Title)
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("#content ul", func(e *colly.HTMLElement) {

		repositories := &Repositories{
			Title: category.Title,
		}

		if e.DOM.Find("li").Length() > 0 {
			e.DOM.Find("li").Each(func(i int, s *goquery.Selection) {
				url := s.Find("a").AttrOr("href", "")
				title := s.Find("a").Text()
				if len(url) > 0 {
					repo := m.fetchGithubInfo(title, url)
					if repo != nil {
						repositories.Blocks = append(repositories.Blocks, repo)
					}
				}
				time.Sleep(time.Second * 1)
			})
		}

		getRepositories = append(getRepositories, repositories)
	})

	c.Visit(category.Url)

	c.Wait()
}

func emptyRepo(title, repoUrl string) *Repo {
	return &Repo{
		Title:            title,
		Url:              repoUrl,
		Description:      "",
		ForksCount:       0,
		NetworkCount:     0,
		OpenIssuesCount:  0,
		StargazersCount:  0,
		SubscribersCount: 0,
		WatchersCount:    0,
	}
}

func (m *AwesomeRepository) fetchGithubInfo(title, repoUrl string) *Repo {
	u, _ := url.Parse(repoUrl)

	if u.Host != "github.com" {
		return emptyRepo(title, repoUrl)
	}

	githubRepoUrl := strings.TrimLeft(u.Path, "/")

	repoInfo := strings.Split(githubRepoUrl, "/")

	if len(repoInfo) < 2 {
		fmt.Printf("%s is not a valid github repo url\n", repoInfo)

		return emptyRepo(title, repoUrl)
	}

	repo, _, err := m.Github.Repositories.Get(m.Context, repoInfo[0], repoInfo[1])

	if err != nil {
		fmt.Printf("Error fetching repo: %s \n", err)

		return emptyRepo(title, repoUrl)
	}
	fmt.Println("Fetched:", repoUrl)
	return &Repo{
		Title:            repo.GetFullName(),
		Description:      repo.GetDescription(),
		Url:              repo.GetHTMLURL(),
		ForksCount:       repo.GetForksCount(),
		NetworkCount:     repo.GetNetworkCount(),
		OpenIssuesCount:  repo.GetOpenIssuesCount(),
		StargazersCount:  repo.GetStargazersCount(),
		SubscribersCount: repo.GetSubscribersCount(),
		WatchersCount:    repo.GetWatchersCount(),
	}
}

func NewAwesomeRepository() AwesomeRepositoryInterface {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	githubClient := github.NewClient(tokenClient)
	return &AwesomeRepository{
		URL:     os.Getenv("MINER_URL"),
		Github:  githubClient,
		Context: context,
	}
}
