package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

func ExampleScrape() {
	//values := url.Values{}
	////req, err := http.NewRequest("GET", "http://metalsucks.net", strings.NewReader(values.Encode()))
	////req, err := http.NewRequest("GET", "https://www.baidu.com/", strings.NewReader(values.Encode()))
	//req, err := http.NewRequest("GET", "http://herzorf.info/#/about", strings.NewReader(values.Encode()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	//client := &http.Client{}
	//res, err := client.Do(req)
	//defer func() {
	//	err2 := res.Body.Close()
	//	if err2 != nil {
	//		fmt.Println("close err", err)
	//	}
	//}()
	//if res.StatusCode != 200 {
	//	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	//}
	////html := "<section class=\"personalData appear\" data-v-7c60eede=\"\"><div class=\"left\" data-v-7c60eede=\"\"><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">姓名</div><div class=\"info\">何中峰(Herzorf)</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">人生目标</div><div class=\"info\">身体自由，思想自由，财务自由</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">目前具备的技能</div><div class=\"info\">web开发</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">以后将要具备的技能</div><div class=\"info\">全栈开发</div></div></div><div class=\"center\" data-v-7c60eede=\"\"><div class=\"border\" data-v-7c60eede=\"\"><img src=\"/assets/me.142f5938.jpg\" height=\"500\" alt=\"自拍\" data-v-7c60eede=\"\"></div></div><div class=\"right\" data-v-7c60eede=\"\"><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">工作经验</div><div class=\"info\">2022.8.1---至今</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">年龄</div><div class=\"info\">23</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">邮箱</div><div class=\"info\">1446450047@qq.com</div></div><div class=\"itemWrapper\" data-v-7c60eede=\"\"><div class=\"title\">目前居住地</div><div class=\"info\">上海</div></div></div></section>"
	//body, err := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	////doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//doc.Find(".about").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the title
	//	//a := s.Find("a")
	//	//p := a.Find("p").Text()
	//	val, _ := s.Attr("p")
	//	fmt.Printf("Review %d: %s\n", i, val)
	//})
	//doc.Find("#app .personalData .right .itemWrapper").EachWithBreak(func(i int, s *goquery.Selection) bool {
	//	title := s.Find(".title").First().Text()
	//	info := s.Find(".info").First().Text()
	//	fmt.Printf("%s: %s\n", title, info)
	//	return true
	//})

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://herzorf.info/#/`),
		chromedp.OuterHTML(`document.querySelector("#app")`, &res, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}

func main() {
	ExampleScrape()
}
