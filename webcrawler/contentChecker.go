package webcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hash/fnv"
)

func CheckDuplication(doc *goquery.Document) bool {
	hash := getContentHash(doc)

	var exists bool
	err := db.QueryRow(
		"SELECT COUNT(*) FROM \"docRecord\" WHERE hash=$1",
		hash).Scan(&exists)
	if err != nil {
		panic(err)
	}

	return exists
}

func getContentHash(doc *goquery.Document) uint32 {
	html, err := doc.Html()
	if err != nil {
		panic(err)
	}

	h := fnv.New32a()
	length, err := h.Write([]byte(html))
	if err != nil {
		fmt.Println(length)
		panic(err)
	}

	return h.Sum32()
}
