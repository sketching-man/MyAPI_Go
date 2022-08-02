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

	// TODO: 지금은 return하면 텍스트 형태로 받음. json 형태로 받게끔 해보자.
	// TODO: 받아서 dict 형태로 parse하고, 그걸 다시 json으로 serialize?
	// jsonData, err := json.Marshal(data)
	// if (nil != err) {
	// 	// do something
	// }

	// return jsonData
}

// region : HackerNews Live Data related
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
// endregion

// region: HackerNews Item related
func GetItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result := requestToHackerNewsAPI("item/" + id + ".json")
	w.Write(result)
}
// endregion

// region: HackerNews User related
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result := requestToHackerNewsAPI("user/" + id + ".json")
	w.Write(result)
}
// endregion
