package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for _, eachline := range fileTextLines {
		//Do something with eachline
	}
	fmt.Printf("Print answer")
}

func part_two() {
	fileTextLines := readInput()
	for _, eachline := range fileTextLines {
		//Do something with eachline
	}
	fmt.Printf("Print answer")
}
func main() {
	part_one()
	part_two()
}
