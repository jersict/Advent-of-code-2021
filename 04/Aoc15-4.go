package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (fileTextLines []string) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	return
}

func getInputs(lines []string) (inputs []string) {
	inputs = strings.Split(lines[0], ",")
	return
}

func getBoards(lines []string) (boards [][][]string) {
	board := [][]string{}
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			boards = append(boards, board)
			board = [][]string{}
		} else {
			board = append(board, strings.Split(lines[i], " "))
		}
	}

	boards = append(boards, board)
	return
}

func crossInput(boards [][][]string, input string) [][][]string {
	for i := 0; i < len(boards); i++ {
		for j := 0; j < len(boards[i]); j++ {
			for k := 0; k < len(boards[i][j]); k++ {
				if boards[i][j][k] == input {
					boards[i][j][k] = "x"
				}
			}
		}
	}
	return boards
}

func checkBoard(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		if board[i][0] == "x" && board[i][1] == "x" && board[i][2] == "x" && board[i][3] == "x" && board[i][4] == "x" {
			return true
		}
		if board[0][i] == "x" && board[1][i] == "x" && board[2][i] == "x" && board[3][i] == "x" && board[4][i] == "x" {
			return true
		}
	}
	return false
}

func checkForBingo(boards [][][]string) int {
	for i := 0; i < len(boards); i++ {
		bingo := checkBoard(boards[i])
		if bingo {
			return i
		}
	}
	return -1
}

func sumOfRemaining(board [][]string) (sum int) {
	sum = 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != "x" {
				num, err := strconv.Atoi(board[i][j])
				if err != nil {
					os.Exit(1)
				} else {
					sum += num
				}

			}
		}
	}
	return
}

func sum(array [100]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func remove(slice [][][]string, s int) [][][]string {
	return append(slice[:s], slice[s+1:]...)
}

func findZero(slice []int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == 0 {
			return i
		}
	}
	return 0
}

func part_one() {
	fileTextLines := readInput()
	inputs := getInputs(fileTextLines)
	boards := getBoards(fileTextLines)
	winner := -1
	lastInput := ""
	for _, input := range inputs {
		boards = crossInput(boards, input)
		winner = checkForBingo(boards)
		if winner != -1 {
			lastInput = input
			break
		}
	}
	sum := sumOfRemaining(boards[winner])
	last, err := strconv.Atoi(lastInput)
	if err != nil {
		os.Exit(2)
	}
	fmt.Printf("Winning score for first winner is %d", sum*last)

}

func part_two() {

	fileTextLines := readInput()
	inputs := getInputs(fileTextLines)
	boards := getBoards(fileTextLines)
	winners := [100]int{}
	winner := -1
	lastInput := ""
	for _, input := range inputs {
		boards = crossInput(boards, input)
	check:
		winner = checkForBingo(boards)
		if winner != -1 {
			if len(boards) == 1 {
				lastInput = input
				break
			}
			lastInput = input
			winners[winner] = 1
			boards = remove(boards, winner)
			goto check
		}
	}
	sum := sumOfRemaining(boards[0])
	last, err := strconv.Atoi(lastInput)
	if err != nil {
		os.Exit(2)
	}
	fmt.Printf("Winning score for last winner is %d", sum*last)

}
func main() {
	part_one()
	fmt.Print(" and the ")
	part_two()
}
