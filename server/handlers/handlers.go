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

func newsDelete(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "request.DeleteNews"
		post := models.Post{}
		err := render.DecodeJSON(r.Body, &post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		s.Delete(&post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		fmt.Fprintf(w, "item deleted")
	}
}

func newsUpdate(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "request.UpdateNews"
		post := models.Post{}
		err := render.DecodeJSON(r.Body, &post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		s.Update(&post)
		if err != nil {
			fmt.Println(op + err.Error())
		}
		fmt.Fprintf(w, "item deleted")
	}
}
