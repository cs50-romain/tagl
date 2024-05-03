package api

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"cs50-romain/tagl/storage"
	"cs50-romain/tagl/types"
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
	file, err := os.OpenFile("./inventory.csv",os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

func DataToCSV(records []*types.EmployeeItems) [][]string {
	data := [][]string{
		{"Record ID", "Employee Name", "Item Name", "Date Acquired", "Quantity", "Ticker Number"},
	}

	// Could be made into a function - better testing
	for _, record := range records {
		strEmployeeID := strconv.Itoa(record.Id)
		strQuantity := strconv.Itoa(record.Quantity)
		strTicket := strconv.Itoa(record.TicketNumber)

		recordData := []string{strEmployeeID, record.EmployeeName, record.ItemName, record.AcquisitionDate.Format("2006-01-02"), strQuantity, strTicket}

		data = append(data, recordData)
	}
	return data

}
