package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortString(s string) (sorted string) {
	sorted = ""
	toSort := strings.Split(s, "")
	sort.Strings(toSort)
	for i := 0; i < len(toSort); i++ {
		sorted += toSort[i]
	}
	return
}

func difference(s1, s2 string) []string {
	a := strings.Split(s1, "")
	b := strings.Split(s2, "")
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func readInput() (all, output [][]string) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileTextLines := []string{}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	for _, eachline := range fileTextLines {
		both := strings.Split(eachline, " | ")
		allLine := strings.Split(both[0], " ")
		outputLine := strings.Split(both[1], " ")
		for i := 0; i < len(allLine); i++ {
			allLine[i] = sortString(allLine[i])
		}
		for i := 0; i < len(outputLine); i++ {
			outputLine[i] = sortString(outputLine[i])
		}
		all = append(all, allLine)
		output = append(output, outputLine)
	}
	return
}

func findOneFourSevenEight(all []string) (one, four, seven, eight string) {
	for j := 0; j < len(all); j++ {
		if len(all[j]) == 2 {
			one = all[j]
		} else if len(all[j]) == 4 {
			four = all[j]
		} else if len(all[j]) == 3 {
			seven = all[j]
		} else if len(all[j]) == 7 {
			eight = all[j]
		}
	}
	return
}
func findZeroSixNine(all []string, eight, one, four string) (zero, six, nine string) {
	partsOfOne := strings.Split(one, "")
	for j := 0; j < len(all); j++ {
		if len(all[j]) == 6 {
			if difference(eight, all[j])[0] == partsOfOne[0] || difference(eight, all[j])[0] == partsOfOne[1] {
				six = all[j]
			} else if len(difference(all[j], four)) == 2 {
				nine = all[j]
			} else {
				zero = all[j]
			}
		}
	}
	return
}

func findTwoThreeFive(all []string, six, one string) (two, three, five string) {
	for j := 0; j < len(all); j++ {
		if len(all[j]) == 5 {
			if len(difference(all[j], six)) == 0 {
				five = all[j]
			} else if len(difference(all[j], one)) == 3 {
				three = all[j]
			} else {
				two = all[j]
			}
		}
	}
	return
}

func decode(output []string, zero, one, two, three, four, five, six, seven, eight, nine string) (code string) {
	code = ""
	for i := 0; i < len(output); i++ {
		switch output[i] {
		case zero:
			code += "0"
		case one:
			code += "1"
		case two:
			code += "2"
		case three:
			code += "3"
		case four:
			code += "4"
		case five:
			code += "5"
		case six:
			code += "6"
		case seven:
			code += "7"
		case eight:
			code += "8"
		case nine:
			code += "9"
		}
	}
	return
}

func part_one() {
	count := 0
	all, output := readInput()
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			if len(output[i][j]) == 2 || len(output[i][j]) == 3 || len(output[i][j]) == 4 || len(output[i][j]) == 7 {
				count++
			}
		}
	}
	all[0] = all[1]
	fmt.Printf("There is %d digits with unique segment values\n", count)
}

func part_two() {
	sum := 0
	all, output := readInput()
	for i := 0; i < len(all); i++ {
		one, four, seven, eight := findOneFourSevenEight(all[i])
		zero, six, nine := findZeroSixNine(all[i], eight, one, four)
		two, three, five := findTwoThreeFive(all[i], six, one)
		result := decode(output[i], zero, one, two, three, four, five, six, seven, eight, nine)
		num, err := strconv.Atoi(result)
		if err != nil {
			os.Exit(1)
		} else {
			sum += num
		}
	}

	fmt.Printf("The sum of decoded values is %d", sum)
}
func main() {
	part_one()
	part_two()
}
