package router

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cache", handlers.SetCache).Methods("POST")
	r.HandleFunc("/cache", handlers.GetCacheList).Methods("GET")
	r.HandleFunc("/cache/{key}", handlers.DeleteCache).Methods("DELETE")
	return r
}
