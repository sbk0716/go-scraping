package scraping

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
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

	// Chromedpのオプションを設定
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	// Chromedpのセッションを開始
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Chromedpを使用してブラウザを制御
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// タイムアウトを設定
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// ページを開く
	err := chromedp.Run(ctx,
		chromedp.Navigate(targetUrl),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ページが完全に読み込まれるまで待機
	time.Sleep(2 * time.Second)

	// ページのHTMLのBodyを取得
	var htmlBody string
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools(`document.body.innerHTML`, &htmlBody),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 取得したHTMLのBodyを表示
	// fmt.Println("HTML Body:")
	// fmt.Println(htmlBody)

	// Bodyを読み込む
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("########################################\n")
	// articleタグの中からタイトルと更新日時といいねのclassを取得する
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".ArticleCard_title__UnBHE").Text()
		updatedAt := s.Find(".ArticleCard_dateAndLikes___O23P > time").Text()
		likes := s.Find(".ArticleCard_dateAndLikes___O23P > .ArticleCard_likes__YCOFM").Text()
		fmt.Printf("%d タイトル: %v / 更新日時: %v いいね: %v\n", i, title, updatedAt, likes)
	})
	fmt.Printf("########################################\n\n")
}
