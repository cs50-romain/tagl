package api

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
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
	router := http.NewServeMux()
	router.HandleFunc("/index", getIndex)
	router.HandleFunc("/submit", handleSubmit)
	router.HandleFunc("/inventory", handleInventory)
	router.HandleFunc("/download", handleDownload)
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
