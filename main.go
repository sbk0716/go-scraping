package main

import (
	"fmt"
	"os"

	"github.com/sbk0716/go-scraping/scraping"
)

func main() {
	fmt.Printf("Execute main function.\n")
	taskName := "example"

	if envTask := os.Getenv("TASK"); envTask != "" {
		taskName = envTask
	}
	fmt.Printf("taskName: %#v\n", taskName)

	if taskName == "example" {
		fmt.Printf("Execute ExampleScrape function.\n")
		scraping.ExampleScrape()
	} else if taskName == "articles" {
		fmt.Printf("Execute ArticlesScrape function.\n")
	} else {
		fmt.Printf("Execute main function.\n")
		scraping.ExampleScrape()
	}
}
