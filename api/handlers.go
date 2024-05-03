package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"cs50-romain/tagl/types"
	util "cs50-romain/tagl/utils"
)

func (s *Server) getIndex(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	http.ServeFile(w, r, "./static/index.html")
	return nil
}

func (s *Server) HandleSubmit(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	
	employeeName := r.FormValue("employeeName")
	itemName := r.FormValue("itemName")
	quantityStr := r.FormValue("quantity")
	ticket := r.FormValue("ticketNumber")
	//date := r.FormValue("date")
	
	ticketNum, err := strconv.Atoi(ticket)
	if err != nil {
		//w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		//w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	// Parse body to create create a new item/employee and pass to storage
	employeeItem := types.NewEmployeeItems(employeeName, itemName, quantity, ticketNum, time.Now())

	err = s.Store.CreateEmployee(employeeItem)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	fmt.Println("Record added!")
	http.Redirect(w, r, "/index", http.StatusSeeOther)
	return nil
}

func (s *Server) handleInventory(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	employeeName := r.FormValue("employeeName")
	fmt.Println("Name:", employeeName)
	employeeItems, err := s.Store.GetEmployeeByName(employeeName)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	WriteJSON(w, employeeItems)
	return nil
}

// User receives a csv file of his inventory. Need csv, writeto.
func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)

	employeeName := r.FormValue("employeeName")
	fmt.Println("Name:", employeeName)
	employeeItems, err := s.Store.GetEmployeeByName(employeeName)
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		return err
	}

	data := DataToCSV(employeeItems)

	if err := WriteCSV(data); err != nil {
	        //w.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(w).Encode("Unable to download file.")
		return err
	}
	w.Header().Set("Content-Disposition", "attachment; filename=inventory.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, "./inventory.csv")
	return nil
}
