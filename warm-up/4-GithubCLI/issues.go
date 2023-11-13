package main

import (
	"bufio"
	"fmt"
	"go-practice/warm-up/4-GithubCLI/github"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues: \n", result.TotalCount)

	page := 3
	for count := 0; count < result.TotalCount; count += page {

		// Don't go out of bounds of the slice.
		if remaining := result.TotalCount - count; page > remaining {
			page = remaining
		}
		for _, issue := range result.Items[count : count+page] {
			fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
		}
		if result.TotalCount-count > page {
			fmt.Println("continue? n")
			input, _ := reader.ReadString('\n')
			if input == "n" {
				continue
			}
		}

	}

	//for _, issue := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	//}
}
