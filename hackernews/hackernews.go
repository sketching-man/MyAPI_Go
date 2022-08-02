package hackernews

import (
	"io/ioutil"
	"net/http"
)

const baseUrl string = "https://hacker-news.firebaseio.com/v0/"

func requestToHackerNewsAPI(path string) []byte {
	resp, err := http.Get(baseUrl + path)
	if (nil != err) {
		// do something
	}
	
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if (nil != err) {
		// do something
	}
	return data
}

func MaxItem(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("maxitem.json")
	w.Write(result)
}

func TopStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("topstories.json")
	w.Write(result)
}

func NewStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("newstories.json")
	w.Write(result)
}

func BestStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("beststories.json")
	w.Write(result)
}

func AskStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("askstories.json")
	w.Write(result)
}

func ShowStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("showstories.json")
	w.Write(result)
}

func JobStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("jobstories.json")
	w.Write(result)
}

func Updates(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("updates.json")
	w.Write(result)
}