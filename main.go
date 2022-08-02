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
	http.HandleFunc("/maxitem", hackernews.MaxItem)
	http.HandleFunc("/topstories", hackernews.TopStories)
	http.HandleFunc("/newstories", hackernews.NewStories)
	http.HandleFunc("/beststories", hackernews.BestStories)
	http.HandleFunc("/askstories", hackernews.AskStories)
	http.HandleFunc("/showstories", hackernews.ShowStories)
	http.HandleFunc("/jobstories", hackernews.JobStories)
	http.HandleFunc("/updates", hackernews.Updates)

	http.HandleFunc("/item", hackernews.GetItem)
	http.HandleFunc("/user", hackernews.GetUser)

	// Start Server
	http.ListenAndServe(":8080", nil)
}