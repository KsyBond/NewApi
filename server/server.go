package server

import (
	"awesomeProject/server/handlers"
	"fmt"
	"net/http"
)

type Server struct {
	Host string
	Port string
}

func New() *Server {
	var s Server
	return &s
}

func (s *Server) Start() {
	http.HandleFunc("/news", handlers.NewsOpen)
	err := http.ListenAndServe(s.Host+":"+s.Port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
