package main

import (
	"github.com/google/go-github/github"
	"net/http"
)

func main() {
	c := github.NewClient(http.Client{})
	if c == nil {
		return;
	}
}

