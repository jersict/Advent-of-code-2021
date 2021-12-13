package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (fileTextLines [][]int) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), "")
		ints := []int{}
		for i := 0; i < len(line); i++ {
			num, err := strconv.Atoi(line[i])
			if err == nil {
				ints = append(ints, num)
			}
		}

		fileTextLines = append(fileTextLines, ints)
	}
	readFile.Close()
	return
}

func make_step(octopusses [][]int) [][]int {
	for i := 0; i < len(octopusses); i++ {
		for j := 0; j < len(octopusses[i]); j++ {
			octopusses[i][j]++
		}
	}
	return octopusses
}

func flash(octopusses [][]int, i, j int) [][]int {
	if i-1 >= 0 && j-1 >= 0 {
		octopusses[i-1][j-1]++
	}
	if i-1 >= 0 {
		octopusses[i-1][j]++
	}
	if j-1 >= 0 {
		octopusses[i][j-1]++
	}
	if i+1 < len(octopusses) {
		octopusses[i+1][j]++
	}
	if j+1 < len(octopusses) {
		octopusses[i][j+1]++
	}
	if i-1 >= 0 && j+1 < len(octopusses) {
		octopusses[i-1][j+1]++
	}
	if i+1 < len(octopusses) && j-1 >= 0 {
		octopusses[i+1][j-1]++
	}
	if i+1 < len(octopusses) && j+1 < len(octopusses) {
		octopusses[i+1][j+1]++
	}
	octopusses[i][j] = -150
	return octopusses
}

func check_for_flashes(octopusses [][]int) ([][]int, bool) {
	change := false
	for i := 0; i < len(octopusses); i++ {
		for j := 0; j < len(octopusses[i]); j++ {
			if octopusses[i][j] > 9 {
				octopusses = flash(octopusses, i, j)
				change = true
			}
		}
	}
	return octopusses, change
}

func count_and_reset(octopusses [][]int) ([][]int, int) {
	count := 0
	for i := 0; i < len(octopusses); i++ {
		for j := 0; j < len(octopusses[i]); j++ {
			if octopusses[i][j] < 0 {
				octopusses[i][j] = 0
				count++
			}
		}
	}
	return octopusses, count
}

func count_check_and_reset(octopusses [][]int) ([][]int, bool) {
	count := 0
	for i := 0; i < len(octopusses); i++ {
		for j := 0; j < len(octopusses[i]); j++ {
			if octopusses[i][j] < 0 {
				octopusses[i][j] = 0
				count++
			}
		}
	}
	if count == 100 {
		return octopusses, true
	}
	return octopusses, false
}

func part_one() {
	numOfSteps := 100
	octopusses := readInput()
	count := 0
	num := 0
	for i := 0; i < numOfSteps; i++ {
		octopusses = make_step(octopusses)
		change := true
		for {
			octopusses, change = check_for_flashes(octopusses)
			if change == false {
				break
			}
		}
		octopusses, num = count_and_reset(octopusses)
		count += num
	}
	fmt.Printf("There has been %d flashes\n", count)
}

func part_two() {
	octopusses := readInput()
	count := 0
	simultanious := false
	for {
		octopusses = make_step(octopusses)
		change := true
		for {
			octopusses, change = check_for_flashes(octopusses)
			if change == false {
				break
			}
		}
		count += 1
		octopusses, simultanious = count_check_and_reset(octopusses)
		if simultanious {
			fmt.Printf("They will align in %d steps\n", count)
			os.Exit(0)
		}
	}
}
func main() {
	part_one()
	part_two()
}
