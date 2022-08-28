package webcrawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CrawlTest() {
	resp, err := http.Get("https://news.ycombinator.com/news")
	if nil != err {
		// do something
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if nil != err {
		// do something
	}

	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	doc.Find("body a").FilterFunction(f).Each(
		func(_ int, tag *goquery.Selection) {
			link, _ := tag.Attr("href")
			linkText := tag.Text()
			fmt.Printf("%s %s\n", linkText, link)
		})
}
