package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Board struct {
	fields                                [][]int
	sideLength                            int
	calculateGroup, applyGroup, stepGroup sync.WaitGroup
}

func initializeBoard(sideLength int) Board {
	fields := make([][]int, sideLength)
	for index, _ := range fields {
		row := initializeRow(sideLength)
		fields[index] = row
	}
	return Board{fields: fields, sideLength: sideLength}
}

func initializeRow(sideLength int) []int {
	row := make([]int, sideLength, sideLength)
	for index, _ := range row {
		row[index] = rand.Intn(2)
	}
	return row
}

func (board *Board) advance() {
	board.calculateGroup.Add(board.size())
	board.stepGroup.Add(board.size())
	board.applyGroup.Add(1)

	for rowIndex, row := range board.fields {
		for columnIndex, _ := range row {
			go board.advanceField(rowIndex, columnIndex)
		}
	}
	board.calculateGroup.Wait()
	board.applyGroup.Done()
	board.stepGroup.Wait()
}

func (board *Board) advanceField(row int, column int) {
	state := board.nextState(row, column)
	board.calculateGroup.Done()
	board.applyGroup.Wait()
	(*board).fields[row][column] = state
	board.stepGroup.Done()
}

func (b *Board) nextState(row int, column int) int {
	rowStart := b.startPosition(row)
	rowEnd := b.endPosition(row)
	columnStart := b.startPosition(column)
	columnEnd := b.endPosition(column)

	sum := 0
	for rowIndex := rowStart; rowIndex <= rowEnd; rowIndex++ {
		for columnIndex := columnStart; columnIndex <= columnEnd; columnIndex++ {
			if rowIndex == row && columnIndex == column {
				continue
			}
			sum = sum + b.fields[rowIndex][columnIndex]
		}
	}
	previousState := b.fields[row][column]
	var result int
	switch {
	case sum <= 2:
		result = 0
	case sum == 2 || sum == 3:
		result = previousState
	case sum >= 3:
		result = 0
	case sum == 3:
		result = 1
	}
	return result
}

func (b *Board) startPosition(position int) int {
	start := position - 1
	if start < 0 {
		start = 0
	}
	return start
}

func (b *Board) endPosition(position int) int {
	upperLimit := b.sideLength - 1
	end := position + 1
	if end > upperLimit {
		end = upperLimit
	}
	return end
}

func (b Board) size() int {
	return b.sideLength * b.sideLength
}

func (b Board) print() {
	for _, row := range b.fields {
		fmt.Println(row)
	}
}
