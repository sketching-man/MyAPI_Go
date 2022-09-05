package webcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

func FilterAndSaveContent(url string, doc *goquery.Document) {
	var title string
	doc.Find("head title").Each(
		func(_ int, selection *goquery.Selection) {
			title = selection.Text()
		},
	)

	queryResult, err := db.Exec(
		"INSERT INTO documents(url, title, \"timestamp\") VALUES ($1, $2, $3)",
		url, title, time.Now())
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}

	queryResult, err = db.Exec(
		"INSERT INTO docRecord(hash) VALUES ($1)",
		getContentHash(doc))
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func FilterAndGetLinks(doc *goquery.Document) *map[string]string {
	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	result := make(map[string]string)
	doc.Find("body a").FilterFunction(f).Each(
		func(_ int, tag *goquery.Selection) {
			link, _ := tag.Attr("href")
			linkText := tag.Text()
			result[link] = linkText
		})

	return &result
}
