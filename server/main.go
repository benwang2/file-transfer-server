package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on port 3000, view at http://localhost:3000.")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
