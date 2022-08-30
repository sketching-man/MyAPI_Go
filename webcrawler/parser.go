package webcrawler

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(htmlBody *io.ReadCloser) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(*htmlBody)
	if nil != err {
		panic(err)
	}

	return doc
}
