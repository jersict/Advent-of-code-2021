package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func window(lines []string, i int) (sum int) {
	val1, err := strconv.Atoi(lines[i])
	if err != nil {
		os.Exit(1)
	}
	val2, err := strconv.Atoi(lines[i+1])
	if err != nil {
		os.Exit(1)
	}
	val3, err := strconv.Atoi(lines[i+2])
	if err != nil {
		os.Exit(1)
	}
	sum = val1 + val2 + val3
	return
}

func main() {
	fileTextLines := readInput()
	count := 0
	for i := 0; i < len(fileTextLines)-3; i++ {
		val_prev := window(fileTextLines, i)
		val_next := window(fileTextLines, i+1)
		if val_next > val_prev {
			count++
		}
	}
	fmt.Printf("%d measurments are langer than the prevoius one", count)
}
