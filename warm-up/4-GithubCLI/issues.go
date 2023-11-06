package main

import (
	"fmt"
	"go-practice/warm-up/4-GithubCLI/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \n", result.TotalCount)
	for _, issue := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}

}
