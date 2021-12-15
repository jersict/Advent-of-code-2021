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

func vertOrHor(line string) (check bool, startX, startY, endX, endY int) {
	check = false
	start := strings.Split(line, " -> ")[0]
	end := strings.Split(line, " -> ")[1]
	startX, err := strconv.Atoi(strings.Split(start, ",")[0])
	if err != nil {
		os.Exit(1)
	}
	startY, err = strconv.Atoi(strings.Split(start, ",")[1])
	if err != nil {
		os.Exit(1)
	}
	endX, err = strconv.Atoi(strings.Split(end, ",")[0])
	if err != nil {
		os.Exit(1)
	}
	endY, err = strconv.Atoi(strings.Split(end, ",")[1])
	if err != nil {
		os.Exit(1)
	}
	if startX == endX || startY == endY {
		check = true
	}
	return
}

func fillInGrid(grid [1000][1000]int, startX, startY, endX, endY int) [1000][1000]int {
	if startX == endX {
		if startY > endY {
			for i := endY; i <= startY; i++ {
				grid[startX][i]++
			}
		} else {
			for i := startY; i <= endY; i++ {
				grid[startX][i]++
			}
		}
	} else if startY == endY {
		if startX > endX {
			for i := endX; i <= startX; i++ {
				grid[i][startY]++
			}
		} else {
			for i := startX; i <= endX; i++ {
				grid[i][startY]++
			}
		}
	} else {
		if startX > endX && startY > endY {
			j := startY
			for i := startX; i >= endX; i-- {

				grid[i][j]++
				j--
			}
		} else if startX > endX && startY < endY {
			j := startY
			for i := startX; i >= endX; i-- {

				grid[i][j]++
				j++
			}
		} else if startX < endX && startY > endY {
			j := startY
			for i := startX; i <= endX; i++ {
				grid[i][j]++
				j--
			}
		} else if startX < endX && startY < endY {

			j := startY
			for i := startX; i <= endX; i++ {
				grid[i][j]++
				j++
			}
		}
	}
	return grid
}

func checkForCrossings(grid [1000][1000]int) (sum int) {
	sum = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] >= 2 {
				sum++
			}
		}
	}
	return
}

func part_one() {
	grid := [1000][1000]int{}
	fileTextLines := readInput()
	for _, eachline := range fileTextLines {
		use, startX, startY, endX, endY := vertOrHor(eachline)
		if use {
			grid = fillInGrid(grid, startX, startY, endX, endY)
		}
	}
	crossings := checkForCrossings(grid)
	fmt.Printf("If we ignore diagonals, there are %d crossings.\n", crossings)
}

func part_two() {
	grid := [1000][1000]int{}
	fileTextLines := readInput()
	for _, eachline := range fileTextLines {
		use, startX, startY, endX, endY := vertOrHor(eachline)
		if use {
			use = false
		}
		grid = fillInGrid(grid, startX, startY, endX, endY)
	}
	crossings := checkForCrossings(grid)
	fmt.Printf("If we count diagonals, there are %d crossings.", crossings)
}
func main() {
	part_one()
	part_two()
}
