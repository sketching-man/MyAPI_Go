package webcrawler

var seedUrlList [2]string = [2]string{
	"https://news.ycombinator.com/news",
	"https://www.bbc.com/news",
}

func Run() {
	for _, url := range seedUrlList {
		Fetch(url)
	}
}
