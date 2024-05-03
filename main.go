package main

import (
	"flag"
	"fmt"
	"log"

	"cs50-romain/tagl/api"
	"cs50-romain/tagl/utils"
	"cs50-romain/tagl/storage"
)

func main() {
	portPtr := flag.String("port", ":6969", "port for http server to listen")
	flag.Parse()

	db, err := storage.NewMysqlStore()
	if err != nil {
		log.Fatal(err)
	}

	s := api.NewServer(*portPtr, db)
	fmt.Println(util.Fore(util.Cyan, "info: server running on port"), *portPtr)
	log.Fatal(s.Start())
}
