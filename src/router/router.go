package router

import (
	"src/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.Index)
	router.HandleFunc("/{key}", middleware.Access).Methods("GET", "POST")
	router.HandleFunc("/api/up", middleware.Upload).Methods("POST")
	router.HandleFunc("/api/dl/{key}", middleware.Download).Methods("GET")
	router.HandleFunc("/api/rand", middleware.RandKey)

	return router
}
