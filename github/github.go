package github

import (
	"fmt"
	"strings"
)

func GetGithubUrl(repository string, branch string) string {

	repositoryParts := strings.Split(repository, "/")

	if len(repositoryParts) != 2 {
		return "https://github.com/new"
	}

	organisation := repositoryParts[0]
	repoName := repositoryParts[1]

	if len(branch) > 0 && len(repository) > 0 {
		return fmt.Sprintf("https://github.com/%s/%s/tree/%s", organisation, repoName, branch)
	}

	if len(repository) > 0 || len(branch) > 0 {
		return "https://github.com/new"
	}

	return "https://github.com/"
}
