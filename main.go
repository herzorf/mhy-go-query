package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

func ExampleScrape() {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.miyoushe.com/ys/`),
		chromedp.OuterHTML(`document.querySelector(".mhy-article-list__body")`, &res, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	doc.Find(".mhy-article-card .mhy-article-card__preview div").Each(func(i int, selection *goquery.Selection) {
		large, existsLarge := selection.Attr("large")
		dataSrc, existsDataSrc := selection.Attr("data-src")
		if existsLarge {
			fmt.Printf("large: %s\n", large)
		}
		if existsDataSrc {
			fmt.Printf("large: %s\n", dataSrc)
		}
	})
}

func main() {
	ExampleScrape()
}
