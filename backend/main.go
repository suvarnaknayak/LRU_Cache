package main

import (
	"log"
	"net/http"
	"backend/router"
	"backend/websocket"
)

func main() {
	r := router.NewRouter()
	go websocket.HandleConnections()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
