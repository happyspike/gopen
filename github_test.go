package main

import (
	"fmt"
	"testing"

	"github.com/ttacon/chalk"
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

	Given("Valid repository with branch", t, func() bool {
		url := GetGithubUrl("happyspike/open", "foo")
		return ShouldMatch("https://github.com/happyspike/open/tree/foo", url, t)
	})
}

func Given(context string, t *testing.T, test func() bool) {
	fmt.Print("Given " + context)
	if test() == false {
		t.Fail()
	}
}

func ShouldMatch(expected string, actual string, t *testing.T) bool {
	if expected != actual {
		fmt.Printf("\n\t%sFAIL: expected '%s', but received '%s'%s\n", chalk.Red, expected, actual, chalk.Reset)
		return false
	}
	fmt.Printf("\t%sPass%s\n", chalk.Green, chalk.Reset)
	return true
}
