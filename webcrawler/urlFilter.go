package webcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func FilterAndSaveContent(doc *goquery.Document) {
	doc.Find("h1").Each(
		func(_ int, selection *goquery.Selection) {
			header := selection.Text()
			fmt.Println(header)
		},
	)
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
