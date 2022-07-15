package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", middleware.Index)
	router.HandleFunc("/up", middleware.Upload).Methods("POST")

	return router
}
