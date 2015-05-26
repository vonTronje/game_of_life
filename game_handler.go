package main

import (
	"html/template"
	"net/http"
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("assets/game_of_life.html")
	(&board).advance()
	template.Execute(w, board)
}
