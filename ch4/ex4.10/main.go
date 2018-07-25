// ex4.10 prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masonelmore/gopl/ch4/ex4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	const (
		day   = time.Hour * 24
		month = day * 30
		year  = day * 365
	)
	recentIssues := []*github.Issue{}
	oldIssues := []*github.Issue{}
	ancientIssues := []*github.Issue{}
	for _, issue := range result.Items {
		age := time.Since(issue.CreatedAt)
		switch {
		case age < month:
			recentIssues = append(recentIssues, issue)
		case age < year:
			oldIssues = append(oldIssues, issue)
		default:
			ancientIssues = append(ancientIssues, issue)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month old")
	for _, issue := range recentIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
	fmt.Println("\nLess than a year old")
	for _, issue := range oldIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
	fmt.Println("\nMore than a year old")
	for _, issue := range ancientIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}
