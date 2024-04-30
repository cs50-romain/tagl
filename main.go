package main

import (
	"flag"
	"fmt"
	"log"

	"cs50-romain/tagl/api"
	"cs50-romain/tagl/utils"
)

func main() {
	portPtr := flag.String("port", ":6969", "port for http server to listen")
	flag.Parse()

	s := api.NewServer(*portPtr)
	fmt.Println(util.Fore(util.Cyan, "info: server running on port"), *portPtr)
	log.Fatal(s.Start())
}
