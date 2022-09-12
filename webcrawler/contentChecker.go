package webcrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hash/fnv"
)

func CheckDuplication(url string, doc *goquery.Document) bool {
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

	// TODO: 문서 단위로 중복 판별하는 방법을 구현해야 함
	return false
}

// 문서의 Hash 값은 같은 문서라도 시간이 다르면 달라지는 것을 확인. 믿을 수 없는 값.
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
