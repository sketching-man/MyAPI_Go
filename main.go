package main

import (
	"fmt"
	"net/http"
)

func hello() {
	fmt.Println("hello, world")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hello()
	})
	http.HandleFunc("/maxitem", hackernews.maxitem)

	http.ListenAndServe(":8080", nil)
}