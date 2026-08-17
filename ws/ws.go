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

func Handler(w http.ResponseWriter, r *http.Request) {
    log.Println("1. HTTP request masuk")

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("2. Upgrade gagal:", err)
        return
    }

    log.Println("2. HTTP berhasil di-upgrade menjadi WebSocket")

    defer conn.Close()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("3. Read gagal:", err)
            return
        }

        log.Printf("3. Message diterima: %s", p)

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Printf("4. Message gagal dikirim: %v", err)
            return
        }

        log.Printf("4. Message dikirim kembali: %s", p)
    }
}
