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

func (s *Server) getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	http.ServeFile(w, r, "./static/index.html")
}

func (s *Server) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	
	employeeName := r.FormValue("employeeName")
	itemName := r.FormValue("itemName")
	quantityStr := r.FormValue("quantity")
	ticket := r.FormValue("ticketNumber")
	//date := r.FormValue("date")
	
	ticketNum, err := strconv.Atoi(ticket)
	if err != nil {
		fmt.Println(err)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse body to create create a new item/employee and pass to storage
	employeeItem := types.NewEmployeeItems(employeeName, itemName, quantity, ticketNum, time.Now())

	err = s.Store.CreateEmployee(employeeItem)
	if err != nil {
		return
	}

	fmt.Println("Record added!")
	
	http.ServeFile(w, r, "./static/index.html")
}

func (s *Server) handleInventory(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	item := types.NewItem("HDMI2VGA", 2)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	WriteJSON(w, item)
}

// User receives a csv file of his inventory. Need csv, writeto.
func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	
	// DATA WILL BE IMPORTED FROM DB LATER
	data := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	if err := WriteCSV(data); err != nil {
	        w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to open file ")
		return	
	}
	w.Header().Set("Content-Disposition", "attachment; filename=inventory.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	http.ServeFile(w, r, "./inventory.csv")
}
