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

func Fetch(url string) {
	// Get HTML from a website
	resp, err := http.Get(url)
	if nil != err {
		// do something
	}
	defer resp.Body.Close()

	// Parse a website
	doc := ParseDocument(&resp.Body)

	// Check if a website is duplicated
	if CheckDuplication(doc) {
		return
	}

	// Get body and <a>s from the parsed document
	if FilterAndSaveContent(doc) {
		return
	}

	links := FilterAndGetLinks(doc)
}
