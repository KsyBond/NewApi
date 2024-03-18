package handlers

import (
	"awesomeProject/server/models"
	"awesomeProject/server/render"
	"awesomeProject/storage"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func NewsOpen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func CreateNews(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "request.CreateNews"
		post := models.Post{}
		err := render.DecodeJSON(r.Body, &post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		s.Create(&post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		fmt.Fprintf(w, "item created")
	}
}

func newsDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/news/delete/")]

	stmt, err := db.Prepare("DELETE FROM News WHERE ID=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "News item with ID %s deleted. Rows affected: %d", id, rowsAffected)
}

func newsUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/news/update/")]
	description := r.FormValue("description")

	stmt, err := db.Prepare("UPDATE News SET Description=? WHERE ID=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(description, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "News item with ID %s updated", id)
}
