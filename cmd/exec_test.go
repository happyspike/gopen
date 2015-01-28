package cmd

import (
	"fmt"
	"testing"
)

func TestEchoWorks(t *testing.T) {
	echoString := "hello, world"
	cmd := fmt.Sprintf("echo %s", echoString)
	result, err := Exec(cmd)

	if err != nil {
		t.Error("not expecting and error: " + err.Error())
	}

	if result != echoString {
		t.Error(fmt.Sprintf("expected '%s', got '%s'", echoString, result))
	}
}

func TestBadCommandFails(t *testing.T) {
	cmd := fmt.Sprintf("asdf")
	result, err := Exec(cmd)

	if err == nil {
		t.Error("expecting an error to occur")
	}

	if len(result) > 0 {
		t.Error("the result should be empty")
	}
}
