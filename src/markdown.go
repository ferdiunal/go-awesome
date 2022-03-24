package src

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Title(title string) string {
	return fmt.Sprintf("# %s\n", title)
}

func Link(title, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func TableRow(repo *Repo) string {
	return fmt.Sprintf("| %s | %s | %d | %d | %d | %d | %d | %d |\n", Link(repo.Title, repo.Url), repo.Description, repo.StargazersCount, repo.ForksCount, repo.NetworkCount, repo.OpenIssuesCount, repo.SubscribersCount, repo.WatchersCount)
}

func TableHeader() string {
	header := fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n", "Repository", "Description", "Stars", "Forks", "Network", "Open Issues", "Subscribers", "Watchers")
	align := "|:---|:---|:---:|:---:|:---:|:---:|:---:|:---:|\n"
	return fmt.Sprintf("%s%s", header, align)
}

func generateMarkdown(data []byte) {
	os.Remove("README.md")

	_ = ioutil.WriteFile("README.md", data, 0644)
}
