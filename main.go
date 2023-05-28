package main

import (
	"os"

	"github.com/sbk0716/go-scraping/scraping"
)

func main() {
	taskName := "books"

	if envTask := os.Getenv("TASK"); envTask != "" {
		taskName = envTask
	}

	if taskName == "books" {
		scraping.BooksExplore()
	} else if taskName == "articles" {
		scraping.ArticlesExplore()
	} else if taskName == "my-articles" {
		scraping.MyArticles()
	} else {
		scraping.ExampleScrape()
	}
}
