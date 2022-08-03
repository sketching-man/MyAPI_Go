package hackernews

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseUrl string = "https://hacker-news.firebaseio.com/v0/"

func requestToHackerNewsAPI(path string) []byte {
	resp, err := http.Get(baseUrl + path)
	if nil != err {
		// do something
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		// do something
	}

	return data
}

func byteArrayToMap(rawData []byte) map[string]interface{} {
	result := make(map[string]interface{})
	err := json.Unmarshal(rawData, &result)
	if nil != err {
		// do something
	}

	return result
}

func mapToJson(mapData map[string]interface{}) []byte {
	result, err := json.Marshal(mapData)
	if nil != err {
		// do something
	}

	return result
}

// region : HackerNews Live Data related
func MaxItem(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("maxitem.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func TopStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("topstories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func NewStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("newstories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func BestStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("beststories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func AskStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("askstories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func ShowStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("showstories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func JobStories(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("jobstories.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func Updates(w http.ResponseWriter, r *http.Request) {
	result := requestToHackerNewsAPI("updates.json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// endregion

// region: HackerNews Item related
func GetItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result := requestToHackerNewsAPI("item/" + id + ".json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// endregion

// region: HackerNews User related
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result := requestToHackerNewsAPI("user/" + id + ".json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

// endregion
