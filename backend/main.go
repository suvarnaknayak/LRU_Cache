package main

import (
	"backend/middleware"
	"backend/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", middleware.CORSMiddleware(r)))
	log.Println("Server started on :8080")
}
