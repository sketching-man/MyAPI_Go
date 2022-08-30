package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sketching-man/MyAPI_Go/module/hackernews"
	"github.com/sketching-man/MyAPI_Go/module/myboard"
	"github.com/sketching-man/MyAPI_Go/module/webcrawler"
)

func hello() {
	fmt.Println("hello, world")
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	hello()
	// 	w.Write([]byte("hello, world"))
	// })

	// // Route Mapping
	// router := mapRoutes()

	// // Start Server
	// http.ListenAndServe(":8080", router)

	// defer myboard.Close()
	webcrawler.Start()
	webcrawler.Close()
}

func mapRoutes() *mux.Router {
	router := mux.NewRouter()

	// region: Hackernews related
	router.HandleFunc("/hackernews/maxitem", hackernews.MaxItem).Methods("GET")
	router.HandleFunc("/hackernews/topstories", hackernews.TopStories).Methods("GET")
	router.HandleFunc("/hackernews/newstories", hackernews.NewStories).Methods("GET")
	router.HandleFunc("/hackernews/beststories", hackernews.BestStories).Methods("GET")
	router.HandleFunc("/hackernews/askstories", hackernews.AskStories).Methods("GET")
	router.HandleFunc("/hackernews/showstories", hackernews.ShowStories).Methods("GET")
	router.HandleFunc("/hackernews/jobstories", hackernews.JobStories).Methods("GET")
	router.HandleFunc("/hackernews/updates", hackernews.Updates).Methods("GET")

	router.HandleFunc("/hackernews/item", hackernews.GetItem).Methods("GET")
	router.HandleFunc("/hackernews/user", hackernews.GetUser).Methods("GET")
	// endregion

	// region: myBoard related
	router.HandleFunc("/myboard/api/member", myboard.CreateMember).Methods("POST")
	router.HandleFunc("/myboard/api/member", myboard.ReadMember).Methods("GET")
	router.HandleFunc("/myboard/api/member", myboard.UpdateMember).Methods("PUT")
	router.HandleFunc("/myboard/api/member", myboard.DeleteMember).Methods("DELETE")
	router.HandleFunc("/myboard/api/article", myboard.CreateArticle).Methods("POST")
	router.HandleFunc("/myboard/api/article", myboard.ReadArticle).Methods("GET")
	router.HandleFunc("/myboard/api/article", myboard.UpdateArticle).Methods("PUT")
	router.HandleFunc("/myboard/api/article", myboard.DeleteArticle).Methods("DELETE")
	// endregion

	return router
}
