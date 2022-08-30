package webcrawler

import "github.com/PuerkitoBio/goquery"

func FilterAndSaveContent(doc *goquery.Document) {
	// TODO: title은 일단 필수로 확보, body는 어떻게 확보할까?
}

func FilterAndGetLinks(doc *goquery.Document) []string {
	return nil
}
