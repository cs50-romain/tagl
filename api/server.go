package api

import (
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/index", getIndex)
	http.HandleFunc("/submit", handleSubmit)
	return http.ListenAndServe(s.addr, nil)
}
