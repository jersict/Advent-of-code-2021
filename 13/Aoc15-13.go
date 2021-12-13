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

func getDotsAndFolds() ([][]int, []string) {
	fileTextLines := readInput()
	paper := [][]int{}
	instr := []string{}
	dots := true
	for _, eachline := range fileTextLines {
		if dots {
			if eachline != "" {
				values := strings.Split(eachline, ",")
				x, err := strconv.Atoi(values[0])
				if err != nil {
					os.Exit(1)
				}
				y, err := strconv.Atoi(values[1])
				if err != nil {
					os.Exit(1)
				}
				dot := []int{x, y}
				paper = append(paper, dot)
			} else {
				dots = false
			}
		} else {
			instr = append(instr, eachline)
		}
	}
	return paper, instr
}

func inSlice(slice [][]int, val []int) bool {
	for _, item := range slice {
		if item[0] == val[0] && item[1] == val[1] {
			return true
		}
	}
	return false
}

func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func fold(dots [][]int, instr string) [][]int {
	axis := 0
	if strings.Split(instr, "")[11] == "y" {
		axis = 1
	}
	value := strings.Split(instr, "=")[1]
	val, err := strconv.Atoi(value)
	if err != nil {
		os.Exit(2)
	}
	for i := 0; i < len(dots); i++ {
		if dots[i][axis] > val {
			newdot := []int{0, 0}
			newdot[axis] = (dots[i][axis] - 2*val) * -1
			newdot[(axis+1)%2] = dots[i][(axis+1)%2]
			if !inSlice(dots, newdot) {
				dots[i][axis] = newdot[axis]
			} else {
				dots = remove(dots, i)
				i--
			}
		}

	}
	return dots
}

func part_one() {
	dots, folds := getDotsAndFolds()
	for _, eachFold := range folds {
		dots = fold(dots, eachFold)
		break
	}
	fmt.Printf("After the first fold there are %d visible dots\n", len(dots))
}

func printDots(dots [][]int) {
	lines := [6][40]string{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			lines[i][j] = " "
		}
	}
	for _, dot := range dots {
		lines[dot[1]][dot[0]] = "#"
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			fmt.Print(lines[i][j])
		}
		fmt.Println("")
	}
}

func part_two() {
	dots, folds := getDotsAndFolds()
	for _, eachFold := range folds {
		dots = fold(dots, eachFold)
	}
	printDots(dots)
}
func main() {
	part_one()
	part_two()
}
