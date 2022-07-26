package webcrawler

import (
	"io"
	"net/http"
)

// func CrawlTest() {
// 	resp, err := http.Get("https://news.ycombinator.com/news")
// 	if nil != err {
// 		// do something
// 	}
// 	defer resp.Body.Close()
//
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if nil != err {
// 		// do something
// 	}
//
// 	f := func(i int, s *goquery.Selection) bool {
// 		link, _ := s.Attr("href")
// 		return strings.HasPrefix(link, "https")
// 	}
//
// 	doc.Find("body a").FilterFunction(f).Each(
// 		func(_ int, tag *goquery.Selection) {
// 			link, _ := tag.Attr("href")
// 			linkText := tag.Text()
// 			fmt.Printf("%s %s\n", linkText, link)
// 		})
// }

func Fetch(url string) {
	// Get HTML from a website
	resp, err := http.Get(url)
	if nil != err {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// Parse a website
	doc := ParseDocument(&resp.Body)

	// Check if a website is duplicated
	if CheckDuplication(url, doc) {
		return
	}

	// Get body from the parsed document and save to DB
	FilterAndSaveContent(url, doc)

	// Get links from the parsed document and renew frontier queue
	links := FilterAndGetLinks(doc)
	DropDuplicatedUrls(links)
	EnqueueUrls(links)
}
