package myboard

import (
	"database/sql"
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
	queryString := "SELECT * FROM public.member"

	rows, err := db.Query(queryString)
	if nil != err {
		panic(err)
	}
	defer rows.Close()

	var id int64
	var grade int
	var password string
	var username string
	for rows.Next() {
		err := rows.Scan(&id, &grade, &password, &username)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, grade, password, username)
	}
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {

}

func DeleteMember(w http.ResponseWriter, r *http.Request) {

}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

}

func ReadArticle(w http.ResponseWriter, r *http.Request) {

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

}
