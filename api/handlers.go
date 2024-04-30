package api

import (
	util "cs50-romain/tagl/utils"
	"fmt"
	"net/http"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	http.ServeFile(w, r, "./static/index.html")
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(util.Fore(util.Yellow, "request:"), r.RequestURI)
	fmt.Println(r.Body)
	http.ServeFile(w, r, "/index")
}
