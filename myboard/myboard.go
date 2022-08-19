package myboard

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	dbinfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PW, DB_NAME)

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
	var member member
	err := json.NewDecoder(r.Body).Decode(&member)
	if nil != err {
		panic(err)
	}

	queryResult, err := db.Exec(
		"INSERT INTO member(grade, password, user_name) VALUES ($1, $2, $3);",
		member.Grade, member.Password, member.Username)
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func ReadMember(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	rows, err := db.Query("SELECT * FROM member WHERE id=$1;", params[0])
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
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	var member member
	err := json.NewDecoder(r.Body).Decode(&member)
	if nil != err {
		panic(err)
	}

	queryResult, err := db.Exec(
		"UPDATE member SET grade=$1, password=$2, user_name=$3 WHERE id=$4;",
		member.Grade, member.Password, member.Username, params[0])
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	queryResult, err := db.Exec("DELETE FROM member WHERE id=$1;", params[0])
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article article
	err := json.NewDecoder(r.Body).Decode(&article)
	if nil != err {
		panic(err)
	}

	queryResult, err := db.Exec(
		"INSERT INTO article(body, title, written_time, writer_id) VALUES ($1, $2, $3, $4)",
		article.Body, article.Title, time.Now(), article.WriterId)
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	rows, err := db.Query("SELECT * FROM article WHERE id=$1", params[0])
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
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	var article article
	err := json.NewDecoder(r.Body).Decode(&article)
	if nil != err {
		panic(err)
	}

	queryResult, err := db.Exec(
		"UPDATE article SET body=$1, title=$2 WHERE id=$3;",
		article.Body, article.Title, params[0])
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["id"]
	if !ok || 1 > len(params[0]) {
		return
	}

	queryResult, err := db.Exec("DELETE FROM article WHERE id=$1;", params[0])
	if nil != err {
		fmt.Println(queryResult)
		panic(err)
	}
}
