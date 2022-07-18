package router

import (
	"net/http"
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	router.HandleFunc("/", middleware.Index)
	router.HandleFunc("/{key}", middleware.Access).Methods("GET", "POST")
	router.HandleFunc("/api/up", middleware.Upload).Methods("POST")
	router.HandleFunc("/api/rand", middleware.RandKey)

	return router
}
