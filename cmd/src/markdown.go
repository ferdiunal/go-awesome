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
	return fmt.Sprintf("| %s | %s | %d | \n", Link(repo.Title, repo.Url), repo.Description, repo.StargazersCount)
}

func TableHeader() string {
	header := fmt.Sprintf("| %s | %s | %s |\n", "Repository", "Description", "Stars")
	align := "|:---|:---|:---:|\n"
	return fmt.Sprintf("%s%s", header, align)
}

func generateMarkdown(data []byte) {
	os.Remove("../README.md")

	_ = ioutil.WriteFile("../README.md", data, 0644)
}
