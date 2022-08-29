package webcrawler

import "github.com/PuerkitoBio/goquery"

func CheckDuplication(doc *goquery.Document) bool {
	hash := getContentHash(doc)
	return hash == ""
}

func getContentHash(doc *goquery.Document) string {
	return ""
}
