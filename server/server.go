package server

import (
	"NewApi/server/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Host   string
	Port   string
	Router *mux.Router
}

func New() *Server {
	var s Server
	s.Router = mux.NewRouter()
	return &s
}

func (s *Server) Start() {
	http.HandleFunc("/news", handlers.NewsOpen)
	err := http.ListenAndServe(s.Host+":"+s.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) configureRouter() {
}
