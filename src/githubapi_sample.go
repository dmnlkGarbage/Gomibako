package main

import (
	"github.com/google/go-github/github"
	"github.com/k0kubun/pp"
)

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "owner"}
	repos, _, err := client.Repositories.ListByOrg("github", opt)
	if err != nil {
		return
	}
	pp.Print(repos)
}
