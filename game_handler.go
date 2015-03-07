package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	messageType := websocket.TextMessage
	delay := time.Duration(3) * time.Second
	for {
		message, err := json.Marshal(board)
		if err != nil {
			log.Println(err)
			continue
		}
		if err = connection.WriteMessage(messageType, message); err != nil {
			log.Println(err)
		}
		(&board).advance()
		time.Sleep(delay)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
