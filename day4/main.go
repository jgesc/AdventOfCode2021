package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_PATH = "input.txt"

type Board struct {
	numbers [5][5]uint8
	marks   [5][5]bool
}

func (board *Board) Mark(markedNumber uint8) {
	for i, row := range board.numbers {
		for j, boardNumber := range row {
			if boardNumber == markedNumber {
				board.marks[i][j] = true
			}
		}
	}
}

func (board *Board) checkBingoRow(row uint8) bool {
	for _, mark := range board.marks[row] {
		if !mark {
			return false
		}
	}
	return true
}

func (board *Board) checkBingoColumn(column uint8) bool {
	for i := 0; i < 5; i++ {
		if !board.marks[i][column] {
			return false
		}
	}
	return true
}

func (board *Board) CheckBingo() bool {
	for i := uint8(0); i < 5; i++ {
		if board.checkBingoRow(i) || board.checkBingoColumn(i) {
			return true
		}
	}
	return false
}

func (board *Board) SumUnmarked() uint16 {
	var sum uint16
	for i, row := range board.marks {
		for j, marked := range row {
			if !marked {
				sum += uint16(board.numbers[i][j])
			}
		}
	}
	return sum
}

func parseInput(inputPath string) ([]uint8, []*Board) {
	input, err := os.Open(inputPath)
	if err != nil {
		panic("Input file not found")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// Parse drawn numbers
	var drawnNumbers []uint8
	scanner.Scan()
	nextLine := scanner.Text()
	for _, numberString := range strings.Split(nextLine, ",") {
		number, err := strconv.ParseUint(numberString, 10, 8)
		if err != nil {
			panic(fmt.Sprintf("Could not parse drawn number '%s'", numberString))
		}
		drawnNumbers = append(drawnNumbers, uint8(number))
	}
	scanner.Scan()

	// Parse boards
	var boards []*Board
	for scanner.Scan() {
		var numbers [5][5]uint8
		for i := 0; i < 5; i++ {
			nextLine := scanner.Text()

			nextLine = strings.TrimSpace(nextLine)
			nextLine = strings.ReplaceAll(nextLine, "  ", " ")
			for j, numberString := range strings.Split(nextLine, " ") {
				number, err := strconv.ParseUint(numberString, 10, 8)

				if err != nil {
					panic(fmt.Sprintf("Could not parse board number '%s'", numberString))
				}
				numbers[i][j] = uint8(number)
			}

			scanner.Scan()
		}
		boards = append(boards, &Board{numbers: numbers})

	}

	return drawnNumbers, boards
}

func runGame(drawnNumbers []uint8, boards []*Board) uint64 {
	for _, bingoNumber := range drawnNumbers {
		for _, board := range boards {
			board.Mark(bingoNumber)
			if board.CheckBingo() {
				unmarkedSum := board.SumUnmarked()
				winningNumber := bingoNumber
				return uint64(unmarkedSum) * uint64(winningNumber)
			}
		}
	}

	return 0
}

func runGameUntilLastBoard(drawnNumbers []uint8, boards []*Board) uint64 {
	var winningBoard *Board
	var winningNumber uint8

	for _, bingoNumber := range drawnNumbers {
		var remainingBoards []*Board
		for _, board := range boards {
			board.Mark(bingoNumber)
			if board.CheckBingo() {
				winningBoard = board
				winningNumber = bingoNumber
			} else {
				remainingBoards = append(remainingBoards, board)
			}
		}
		boards = remainingBoards
		if len(boards) == 0 {
			break
		}
	}

	unmarkedSum := winningBoard.SumUnmarked()
	return uint64(unmarkedSum) * uint64(winningNumber)
}

func main() {
	fmt.Print(
		"--- Advent of Code 2021 ---\n",
		"          Day   4          \n",
		"                           \n",
		"          Part  A          \n")

	drawnNumbers, boards := parseInput(INPUT_PATH)
	finalScore := runGame(drawnNumbers, boards)
	fmt.Printf("Final Score: %d\n", finalScore)

	fmt.Printf("\n          Part  B          \n")

	finalScore = runGameUntilLastBoard(drawnNumbers, boards)
	fmt.Printf("Last Board Final Score: %d\n", finalScore)

	fmt.Printf("\n---------------------------\n")
}
