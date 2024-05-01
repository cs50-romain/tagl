package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cs50-romain/tagl/types"
	util "cs50-romain/tagl/utils"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	http.ServeFile(w, r, "./static/index.html")
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	fmt.Println(r.Body)
	http.ServeFile(w, r, "./static/index.html")
}

func handleInventory(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	item := types.NewItem("HDMI2VGA", 2)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	WriteJSON(w, item)
}

// User receives a csv file of his inventory. Need csv, writeto.
func handleDownload(w http.ResponseWriter, r *http.Request) {
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
