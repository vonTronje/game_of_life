package main

import "net/http"

var board Board

func main() {
	board = initializeBoard(25)
	http.HandleFunc("/gol", gameHandler)
}
