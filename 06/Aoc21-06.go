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

func next_day(school [9]int) [9]int {
	first := school[0]
	for i := 0; i < len(school)-1; i++ {
		school[i] = school[i+1]
	}
	school[6] += first
	school[8] = first
	return school
}
func sum(array [9]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func fishAfterXDays(numOfDays int) {
	school := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	numbers := readInput()
	for _, eachnum := range numbers {
		school[eachnum]++
	}
	for i := 0; i < numOfDays; i++ {
		school = next_day(school)
	}
	fmt.Printf("After %d there is %d fish.\n", numOfDays, sum(school))
}

func main() {
	fishAfterXDays(256)
}
