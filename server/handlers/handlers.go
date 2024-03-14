package handlers

import (
	"fmt"
	"net/http"
	- "github.com/mattn/go-sqlite3"
)

func NewsOpen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func newsCreate(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	stmt, err := db.Prepare("INSERT INTO News (Title, Description) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := stmt.Exec(title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "News item created with ID: %d", lastID)
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
