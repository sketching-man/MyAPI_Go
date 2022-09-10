package webcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hash/fnv"
)

func CheckDuplication(url string, doc *goquery.Document) bool {
	//var exists bool
	exists := false
	// region Content check with hash
	//hash := getContentHash(doc)
	//
	//err := db.QueryRow(
	//	"SELECT COUNT(*) FROM \"docRecord\" WHERE hash=$1",
	//	hash).Scan(&exists)
	//if err != nil {
	//	panic(err)
	//}
	// endregion

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
