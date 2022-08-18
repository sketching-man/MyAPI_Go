package myboard

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	DB_HOST = "localhost"
	DB_NAME = "board"
	DB_USER = "postgres"
	DB_PW   = "postgres"
)

var (
	db *sql.DB
)

func init() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PW, DB_NAME)

	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
}

func Close() {
	if nil != db {
		db.Close()
	}
}

func CreateMember(w http.ResponseWriter, r *http.Request) {

}

func ReadMember(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	queryString := fmt.Sprintf("SELECT * FROM member WHERE id=%s", params[0])
	rows, err := db.Query(queryString)
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	members := []member{}
	for rows.Next() {
		var member member

		err := rows.Scan(&member.Id, &member.Grade, &member.Password, &member.Username)
		if err != nil {
			panic(err)
		}

		members = append(members, member)
	}

	result, err := json.Marshal(members)
	if nil != err {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {

}

func DeleteMember(w http.ResponseWriter, r *http.Request) {

}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	queryString := fmt.Sprintf("SELECT * FROM article WHERE id=%s", params[0])

	rows, err := db.Query(queryString)
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	articles := []article{}
	for rows.Next() {
		var article article

		err := rows.Scan(&article.Id, &article.Body, &article.Title, &article.WrittenTime, &article.WriterId)
		if err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}

	result, err := json.Marshal(articles)
	if nil != err {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
