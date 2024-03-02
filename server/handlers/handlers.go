package handlers

import (
	"fmt"
	"net/http"
)

func NewsOpen(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func newsCreate(w http.ResponseWriter, r *http.Request) {

}

func newsDelete(w http.ResponseWriter, r *http.Request) {

}

func newsUpdate(w http.ResponseWriter, r *http.Request) {

}
