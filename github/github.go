package github

import (
	"fmt"
	"strings"
)

func GetGithubUrl(repository string, branch string) string {

	repositoryParts := strings.Split(repository, "/")
	if len(repositoryParts) != 2 && (len(repository) != 0 && len(branch) != 0) {
		return "https://github.com/new"
	}

	if len(branch) > 0 && len(repository) > 0 {
		organisation := repositoryParts[0]
		repoName := repositoryParts[1]
		return fmt.Sprintf("https://github.com/%s/%s/tree/%s", organisation, repoName, branch)
	}

	if len(repository) > 0 || len(branch) > 0 {
		return "https://github.com/new"
	}

	return "https://github.com/"
}
