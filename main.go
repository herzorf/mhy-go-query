package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	var num = 1
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	doc.Find(".mhy-article-card .mhy-article-card__preview div").Each(func(i int, selection *goquery.Selection) {
		large, existsLarge := selection.Attr("large")
		dataSrc, existsDataSrc := selection.Attr("data-src")
		if existsLarge {
			err := downloadFile(large, strconv.Itoa(num)+".jpg")
			if err != nil {
				fmt.Println("download err ", err)
			} else {
				fmt.Println("large下载成功", num)
			}
			num = num + 1
		} else if existsDataSrc {
			err := downloadFile(dataSrc, strconv.Itoa(num)+".jpg")
			if err != nil {
				fmt.Println("download err ", err)
			} else {
				fmt.Println("datasrc下载成功", num)
			}
			num = num + 1
		}
	})
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	file, err := os.Create("./images/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ExampleScrape()
	//downloadFile("https://upload-bbs.miyoushe.com/upload/2022/12/20/75276539/c2aff9d62128e681b6a85b95775ad5c1_8462257324862820376.jpg", "name.jpg")
}
