package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"src/database"
	"src/env"
	"src/router"
	"syscall"
)

func cleanup() {
	database.Stop()
}

func main() {
	env.LoadVars()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
	database.Start()
	r := router.Router()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)
	fmt.Println("Starting server on port 3000, view at http://localhost:3000.")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
