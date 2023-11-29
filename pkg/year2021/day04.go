// Code generated by aocgen; DO NOT EDIT.
package year2021

import (
	"regexp"
	"strconv"
	"strings"
)

type Day04 struct{}

type BingoBoard [][]BingoNumber

type BingoBoards []BingoBoard

type BingoNumber struct {
	value string
	marked bool
}

func (bs BingoBoards) getWinningBoard() *BingoBoard{
	for _, board := range bs {
		if board.HasBoardBingo() {
			return &board
		}
	}

	return nil
}

func (bb BingoBoard) SumOfAllUnmarkedNumbers() int {
	sum := 0
	for _, row := range bb {
		for _, number := range row {
			if !number.marked {
				value, err := strconv.Atoi(number.value)
				if err != nil {
					panic(err)
				}
				sum += value
			}
		}
	}
	return sum
}

func (bb BingoBoard) HasBoardBingo() bool {
	colMarkedAmounts := map[int]int{}
	rowMarkedAmounts := map[int]int{}

	for rowIndex, row := range bb {
		for colIndex, number := range row {
			if number.marked {
				rowMarkedAmounts[rowIndex]++
				colMarkedAmounts[colIndex]++
			}
		}
	}

	for _, colAmount := range colMarkedAmounts {
		if colAmount == len(bb) {
			return true
		}
	}

	for _, rowAmount := range rowMarkedAmounts {
		if rowAmount == len(bb) {
			return true
		}
	}

	return false
}

func (bb BingoBoard) MarkNumberWhenPresent(drawnNumber string) {
	for row, bingoRow := range bb {
		for col,number := range bingoRow {
			if drawnNumber == number.value {
				number.marked = true
				bb[row][col] = number
			}
		}
	}
}

func (p Day04) PartA(lines []string) any {
	bingoNumbers := strings.Split(lines[0],",")
	bingoBoardsRaw := lines[2:]
	var bingoBoards BingoBoards
	currentBoardIndex := 0

	for _, boardRow := range bingoBoardsRaw {
		if boardRow == "" {
			currentBoardIndex++
			continue
		}


		rowOfBingoNumbers := parseRowOfBingoNumbers(boardRow)

		if len(bingoBoards) != currentBoardIndex + 1 {
			bingoBoards = append(bingoBoards,[][]BingoNumber{})
		}
		currentBingoBoard := bingoBoards[currentBoardIndex]
		currentBingoBoard = append(currentBingoBoard, rowOfBingoNumbers)
		bingoBoards[currentBoardIndex] = currentBingoBoard
	}

	for _, drawnNumber := range bingoNumbers {
		for _, board := range bingoBoards {
			board.MarkNumberWhenPresent(drawnNumber)
		}

		if winningBoard := bingoBoards.getWinningBoard(); winningBoard != nil {
			sum := winningBoard.SumOfAllUnmarkedNumbers()
			drawnNumberInt,err := strconv.Atoi(drawnNumber)
			if err != nil {
				panic(err)
			}

			return sum * drawnNumberInt

			break
		}
	}

	return 0
}



func parseRowOfBingoNumbers(boardRow string) []BingoNumber {
	bingoRowNumbersRaw := regexp.MustCompile("\\s+").Split(boardRow, -1)
	var bingoRowNumbers []BingoNumber
	for _, number := range bingoRowNumbersRaw {
		bingoRowNumbers = append(bingoRowNumbers, BingoNumber{value: number})
	}
	return bingoRowNumbers
}

func (p Day04) PartB(lines []string) any {
	return "implement_me"
}
