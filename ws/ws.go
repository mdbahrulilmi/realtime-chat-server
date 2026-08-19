package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connections []*websocket.Conn

func removeConnection(conn *websocket.Conn) {
	for i, client := range connections {
		if client == conn {
			connections = append(
				connections[:i],
				connections[i+1:]...,
			)

			break
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP request masuk")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade gagal:", err)
		return
	}

	defer conn.Close()

	connections = append(connections, conn)

	log.Println("Jumlah client:", len(connections))

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read gagal:", err)
            removeConnection(conn)
            log.Println("Jumlah client:", len(connections))
			return
		}

		for _, client := range connections {
            if client == conn {
                continue
            }

            err := client.WriteMessage(
                websocket.TextMessage,
                message,
            )

            if err != nil {
                log.Println("Broadcast error:", err)
            }
        }
	}
}