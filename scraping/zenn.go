package scraping

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
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
