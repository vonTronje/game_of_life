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
	var waitGroup sync.WaitGroup

	for rowIndex, row := range board.fields {
		for columnIndex, _ := range row {
			waitGroup.Add(1)
			go board.advanceField(rowIndex, columnIndex, &waitGroup)
		}
	}
	waitGroup.Wait()
}
