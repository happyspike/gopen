package main

import (
	"testing"
)

func TestGetRepositoryFromOriginUrl(t *testing.T) {
	url := getRepositoryFromOriginUrl("git@github.com:contikiholidays/contiki.git")
	if url != "contikiholidays/contiki" {
		t.Error("contikiholidays/contiki" + ", but got " + url)
	}

	url = getRepositoryFromOriginUrl("")
	if url != "" {
		t.Error("" + ", but got " + url)
	}

	url = getRepositoryFromOriginUrl("asdasd  asdasd")
	if url != "" {
		t.Error("" + ", but got " + url)
	}
}
