package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (start []int) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileTextLines := []string{}

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	vals := strings.Split(fileTextLines[0], ",")
	for i := 0; i < len(vals); i++ {
		num, err := strconv.Atoi(vals[i])
		if err != nil {
			log.Fatalf("failed to convert: %s", err)
		} else {
			start = append(start, num)
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(array [1000]int) (min int) {
	min = 9999999999
	for i := 0; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return
}

func minMax(a int, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func part_one() {
	start := readInput()
	moves := [1000]int{}
	for i := 0; i < len(moves); i++ {
		fuel := 0
		for j := 0; j < len(start); j++ {
			fuel += abs(start[j] - i)
		}
		moves[i] = fuel
	}
	fmt.Printf("***PART 1***\nLowest amount of fuel needed is %d\n\n", min(moves))
}

func part_two() {
	start := readInput()
	moves := [1000]int{}
	for i := 0; i < len(moves); i++ {
		fuel := 0
		for j := 0; j < len(start); j++ {
			s, e := minMax(start[j], i)
			for k := s; k < e; k++ {
				fuel += e - k
			}
		}
		moves[i] = fuel
	}
	fmt.Printf("***PART 2***\nLowest amount of fuel needed is %d ", min(moves))
}
func main() {
	part_one()
	part_two()
}
