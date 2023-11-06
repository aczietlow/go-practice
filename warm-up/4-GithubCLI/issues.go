package main

import (
	"github.com/aczietlow/go-practice/warm-up/4-GithubCLI/github/github"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
}
