package main

import (
	"sync"
)

func main() {
	board := initializeBoard(5)
	board.print()
	for {
		advance(&board)
		board.print()
	}
}

func advance(board *Board) {
	var calculateGroup sync.WaitGroup
	calculateGroup.Add(board.size())

	for rowIndex, row := range board.fields {
		for columnIndex, _ := range row {
			go board.advanceField(rowIndex, columnIndex, &calculateGroup)
		}
	}
	calculateGroup.Wait()
}
