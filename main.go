package main

import (
	"fmt"
	"net/http"

	"github.com/sketching-man/MyAPI_Go/module/hackernews"
)

func hello() {
	fmt.Println("hello, world")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hello()
		w.Write([]byte("hello, world"))
	})

	// Route Mapping
	mapRoutes()

	// Start Server
	http.ListenAndServe(":8080", nil)
}

func mapRoutes() {
	// region: Hackernews related
	http.HandleFunc("/hackernews/maxitem", hackernews.MaxItem)
	http.HandleFunc("/hackernews/topstories", hackernews.TopStories)
	http.HandleFunc("/hackernews/newstories", hackernews.NewStories)
	http.HandleFunc("/hackernews/beststories", hackernews.BestStories)
	http.HandleFunc("/hackernews/askstories", hackernews.AskStories)
	http.HandleFunc("/hackernews/showstories", hackernews.ShowStories)
	http.HandleFunc("/hackernews/jobstories", hackernews.JobStories)
	http.HandleFunc("/hackernews/updates", hackernews.Updates)

	http.HandleFunc("/hackernews/item", hackernews.GetItem)
	http.HandleFunc("/hackernews/user", hackernews.GetUser)
	// endregion

	// region: myBoard related
	http.HandleFunc("/myboard/api/article", nil)
	http.HandleFunc("/myboard/api/member", nil)
	// endregion
}
