package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("WebSocket endpoint hit")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket Upgrade Error:", err)
			return
		}
		defer conn.Close()
		log.Println("WebSocket connection established")

		clients[conn] = true

		for {
			var msg map[string]interface{}
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("WebSocket Read Error:", err)
				delete(clients, conn)
				break
			}
			broadcast <- msg
		}
	})

	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Println("WebSocket Write Error:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()
}
