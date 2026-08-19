package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mdbahrulilmi/realtime-chat-server/ws"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func main() {
	port := ":2500"
	http.HandleFunc("/ws", ws.Handler)
	http.HandleFunc("/", Handler)
	fmt.Printf("server runnning at http://localhost%s ", port)
    log.Fatal(http.ListenAndServe(port, nil))
}