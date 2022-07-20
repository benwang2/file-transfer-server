package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.Index)
	router.HandleFunc("/{key}", middleware.Access).Methods("GET", "POST")
	router.HandleFunc("/api/up", middleware.Upload).Methods("POST")
	router.HandleFunc("/api/rand", middleware.RandKey)

	return router
}
