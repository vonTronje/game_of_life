package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Board struct {
	fields [][]int
	size   int
}

func initializeBoard(size int) Board {
	fields := make([][]int, size)
	for index, _ := range fields {
		row := initializeRow(size)
		fields[index] = row
	}
	return Board{fields: fields, size: size}
}

func initializeRow(size int) []int {
	row := make([]int, size, size)
	for index, _ := range row {
		row[index] = rand.Intn(2)
	}
	return row
}

func (b *Board) advanceField(row int, column int, waitGroup *sync.WaitGroup) {
	state := b.nextState(row, column)
	waitGroup.Done()
	waitGroup.Wait()
	b.fields[row][column] = state
}

func (b Board) nextState(row int, column int) int {
	rowStart := b.startPosition(row)
	rowEnd := b.endPosition(row)
	columnStart := b.startPosition(column)
	columnEnd := b.endPosition(column)

	sum := 0
	for rowIndex := rowStart; rowIndex <= rowEnd; rowIndex++ {
		for columnIndex := columnStart; columnIndex <= columnEnd; columnEnd++ {
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

func (b Board) startPosition(position int) int {
	start := position - 1
	if start < 0 {
		start = 0
	}
	return start
}

func (b Board) endPosition(position int) int {
	upperLimit := b.size - 1
	end := position + 1
	if end > upperLimit {
		end = upperLimit
	}
	return end
}

func (b Board) print() {
	for _, row := range b.fields {
		fmt.Println(row)
	}
}
