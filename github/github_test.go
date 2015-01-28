package github

import (
	"testing"

	. "github.com/happyspike/gopen/testhelpers"
)

func TestGetGithubUrl(t *testing.T) {
	Given("Valid repository with no branch", t, func() bool {
		url := GetGithubUrl("happyspike/gopen", "")
		return ShouldMatch("https://github.com/new", url, t)
	})

	Given("Invalid repository with branch", t, func() bool {
		url := GetGithubUrl("wrongformat", "master")
		return ShouldMatch("https://github.com/new", url, t)
	})

	Given("No repository with branch", t, func() bool {
		url := GetGithubUrl("", "foo")
		return ShouldMatch("https://github.com/new", url, t)
	})

<<<<<<< HEAD
	Given("No repository with no branch", t, func() bool {
		url := GetGithubUrl("", "")
		return ShouldMatch("https://github.com/", url, t)
	})

=======
>>>>>>> 2f517bd9f967009898a1afdff8a33a388a9d08f8
	Given("Valid repository with branch", t, func() bool {
		url := GetGithubUrl("happyspike/open", "foo")
		return ShouldMatch("https://github.com/happyspike/open/tree/foo", url, t)
	})
}
