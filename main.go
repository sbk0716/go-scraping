package main

import (
	"fmt"
	"os"

	"github.com/sbk0716/go-scraping/scraping"
)

func main() {
	fmt.Printf("Execute main function.\n")
	taskName := "books"

	if envTask := os.Getenv("TASK"); envTask != "" {
		taskName = envTask
	}
	fmt.Printf("taskName: %#v\n", taskName)

	if taskName == "books" {
		fmt.Printf("Execute BooksExplore function.\n")
		scraping.BooksExplore()
	} else if taskName == "articles" {
		fmt.Printf("Execute ArticlesExplore function.\n")
		scraping.ArticlesExplore()
	} else if taskName == "my-articles" {
		fmt.Printf("Execute myArticles function.\n")
		scraping.MyArticles()
	} else {
		fmt.Printf("Execute ExampleScrape function.\n")
		scraping.ExampleScrape()
	}
}
