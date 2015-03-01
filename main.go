package main

import "time"

func main() {
	delay := time.Duration(3) * time.Second
	board := initializeBoard(25)
	board.print()
	for {
		(&board).advance()
		board.print()
		time.Sleep(delay)
	}
}
