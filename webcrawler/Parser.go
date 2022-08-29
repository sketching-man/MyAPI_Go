package webcrawler

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(htmlBody *io.ReadCloser) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(*htmlBody)
	if nil != err {
		fmt.Println(err)
	}

	return doc
}
