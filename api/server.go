package api

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"cs50-romain/tagl/storage"
)

type Server struct {
	addr string
	Store storage.Storer
}

func NewServer(addr string, store storage.Storer) *Server {
	return &Server{
		addr: addr,
		Store: store,
	}
}

func (s *Server) Start() error {
	router := http.NewServeMux()
	router.HandleFunc("/index", s.getIndex)
	router.HandleFunc("/submit", s.HandleSubmit)
	router.HandleFunc("/inventory", s.handleInventory)
	router.HandleFunc("/download", s.handleDownload)
	return http.ListenAndServe(s.addr, router)
}

func WriteJSON(w io.Writer, d any) error {
	return json.NewEncoder(w).Encode(d)
}

func WriteCSV(data [][]string) error {
	file, err := os.OpenFile("./inventory.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	cw := csv.NewWriter(file)
	for _, d := range data {
		if err := cw.Write(d); err != nil {
			return err
		}
	}

	cw.Flush()
	return nil
}
