package webcrawler

import (
	"sync"

	"github.com/sketching-man/MyAPI_Go/module/datastructure"
)

var seedUrlList [2]string = [2]string{
	"https://news.ycombinator.com/news",
	"https://www.bbc.com/news",
}
var urlQueue datastructure.Queue

func StartCrawling() {
	urlQueue = *datastructure.NewQueue()

	for _, url := range seedUrlList {
		urlQueue.Push(url)
	}

	var wait sync.WaitGroup
	wait.Add(1)

	go run(&wait, &urlQueue)
	wait.Wait()
}

func run(wait *sync.WaitGroup, urlQueue *datastructure.Queue) {
	defer wait.Done()

	for {
		url := urlQueue.Pop()
		if nil == url {
			break
		}

		Fetch(url.(string))
	}
}

func EnqueueUrls(urls *map[string]string) {
	for link, _ := range *urls {
		urlQueue.Push(link)
	}
}
