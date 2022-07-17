package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/database"
	"server/env"
	"server/router"
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

	// "ben1:harbor-enquirer-written@tcp(107.174.63.205:3306)/mydb"

	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on port 3000, view at http://localhost:3000.")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
