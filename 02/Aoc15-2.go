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

func part_one() {
	fileTextLines := readInput()
	depth := 0
	hpos := 0
	for _, eachline := range fileTextLines {
		line := strings.Split(eachline, " ")
		value, err := strconv.Atoi(line[1])
		if err != nil {
			os.Exit(1)
		}
		switch line[0] {
		case "down":
			depth += value
		case "forward":
			hpos += value
		case "up":
			depth -= value
		}
	}
	fmt.Printf("###PART 1###\nDepth is %d, and horizontal position is %d. Multiplied together the result is %d.\n\n", depth, hpos, depth*hpos)
}

func part_two() {
	fileTextLines := readInput()
	depth := 0
	hpos := 0
	aim := 0
	for _, eachline := range fileTextLines {
		line := strings.Split(eachline, " ")
		value, err := strconv.Atoi(line[1])
		if err != nil {
			os.Exit(1)
		}
		switch line[0] {
		case "down":
			aim += value
		case "forward":
			hpos += value
			depth += value * aim
		case "up":
			aim -= value
		}
	}
	fmt.Printf("###PART 2###\nDepth is %d, and horizontal position is %d. Multiplied together the result is %d.\n\n", depth, hpos, depth*hpos)
}
func main() {
	part_one()
	part_two()
}
