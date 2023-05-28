package scraping

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func BooksExplore() {
	// 読み込むページ
	res, err := http.Get("https://zenn.dev/books/explore")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Bodyを読み込む
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// articleタグの中から著者とタイトルのclassを取得する
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		author := s.Find(".BookLargeLink_userName__jNbk5").Text()
		book := s.Find(".BookLargeLink_title__RqL6r").Text()
		if "" == author {
			author = s.Find(".BookLink_userName__avtjq").Text()
		}
		if "" == book {
			book = s.Find(".BookLink_title__b8hGg").Text()
		}
		fmt.Printf("%d 著者: %v / タイトル: %v\n", i, author, book)
	})
}

func ArticlesExplore() {
	// 読み込むページ
	res, err := http.Get("https://zenn.dev/articles/explore")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	time.Sleep(10 * time.Second)

	// Bodyを読み込む
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// articleタグの中から著者とタイトルのclassを取得する
	doc.Find("article").Each(func(i int, s *goquery.Selection) {

		author := s.Find(".ArticleList_userName__GWXDx").Text()
		article := s.Find(".ArticleList_title__P6X2G").Text()
		fmt.Printf("%d 著者: %v / タイトル: %v\n", i, author, article)
	})
}

func MyArticles() {
	userName := "sbk0716"
	if envUserName := os.Getenv("USERNAME"); envUserName != "" {
		userName = envUserName
	}
	fmt.Printf("userName: %#v\n", userName)
	targetUrl := fmt.Sprintf("https://zenn.dev/%s", userName)
	fmt.Printf("targetUrl: %#v\n", targetUrl)

	// 読み込むページ
	res, err := http.Get(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Bodyを読み込む
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// articleタグの中からタイトルと更新日時のclassを取得する
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".ArticleCard_title__UnBHE").Text()
		updatedAt := s.Find(".ArticleCard_dateAndLikes___O23P").Text()
		fmt.Printf("%d タイトル: %v / 更新日時: %v\n", i, title, updatedAt)
	})
}
