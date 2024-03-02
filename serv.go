package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type News struct {
	ID          int
	Title       string
	Description string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./news.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
