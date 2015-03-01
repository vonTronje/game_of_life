package main


func main() {
	board := initializeBoard(5)
	board.print()
	for {
		(&board).advance()
		board.print()
	}
}
