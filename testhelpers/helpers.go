package testhelpers

import (
	"fmt"
	"testing"

	"github.com/ttacon/chalk"
)

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
