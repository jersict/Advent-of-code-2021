package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func getInputs() (map[string]int, map[string]string) {
	instr := make(map[string]string)
	pairs := make(map[string]int)
	fileTextLines := readInput()
	template := strings.Split(fileTextLines[0], "")
	for i := 2; i < len(fileTextLines); i++ {
		line := strings.Split(fileTextLines[i], " -> ")
		instr[line[0]] = line[1]
		pairs[line[0]] = 0
	}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i]+template[i+1]]++
	}
	return pairs, instr
}

func makeStep(pairs map[string]int, instructions map[string]string) map[string]int {
	newPairs := make(map[string]int)
	for key, value := range pairs {
		if value > 0 {
			ins := instructions[key]
			split := strings.Split(key, "")
			pairs[key] = 0
			if _, ok := newPairs[split[0]+ins]; ok {
				newPairs[split[0]+ins] += value
			} else {
				newPairs[split[0]+ins] = value
			}
			if _, ok := newPairs[ins+split[1]]; ok {
				newPairs[ins+split[1]] += value
			} else {
				newPairs[ins+split[1]] = value
			}
		}
	}
	return newPairs
}

func countElements(pairs map[string]int) int {
	count := make(map[string]int)
	for key, value := range pairs {
		split := strings.Split(key, "")
		if _, ok := count[split[0]]; ok {
			count[split[0]] += value
		} else {
			count[split[0]] = value
		}
		if _, ok := count[split[1]]; ok {
			count[split[1]] += value
		} else {
			count[split[1]] = value
		}
	}
	v := []int{}
	for _, value := range count {
		v = append(v, int(math.Ceil(float64(value)/float64(2.0))))
	}
	sort.Ints(v)
	return v[len(v)-1] - v[0]
}

func part_one() {
	numOfSteps := 10
	pairs, instructions := getInputs()
	for i := 0; i < numOfSteps; i++ {
		pairs = makeStep(pairs, instructions)
	}
	fmt.Printf("The difference between most and least common element is %d\n", countElements(pairs))
}

func part_two() {
	numOfSteps := 40
	pairs, instructions := getInputs()
	for i := 0; i < numOfSteps; i++ {
		pairs = makeStep(pairs, instructions)
	}
	fmt.Printf("The difference between most and least common element is %d\n", countElements(pairs))
}
func main() {
	part_one()
	part_two()
}
