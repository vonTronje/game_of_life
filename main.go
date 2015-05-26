package main

import (
	"log"
	"net/http"
)

var board Board

func main() {
	board = initializeBoard(25)
	http.HandleFunc("/game_of_life", gameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
