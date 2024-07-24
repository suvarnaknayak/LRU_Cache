package router

import (
    "github.com/gorilla/mux"
    "backend/handlers"
)

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/cache/get", handlers.GetCache).Methods("GET")
    r.HandleFunc("/cache/set", handlers.SetCache).Methods("POST")
    r.HandleFunc("/cache/delete", handlers.DeleteCache).Methods("DELETE")
    return r
}
