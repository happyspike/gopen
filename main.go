package main

import (
	"strings"

	"github.com/happyspike/gopen/cmd"
	"github.com/happyspike/gopen/github"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	origin, _ := cmd.GitOriginUrl()
	branch, _ := cmd.GitBranch()
	repository := getRepositoryFromOriginUrl(origin)
	open.Start(github.GetGithubUrl(repository, branch))
}

func getRepositoryFromOriginUrl(url string) (repository string) {
	repository = ""
	if strings.Contains(url, "git@github.com:") {
		urlLength := len(url)
		githubLength := len("git@github.com:")
		formatLength := len(".git")

		if urlLength > 0 {
			repository = url[githubLength : urlLength-formatLength]
		}
	}
	return
}
