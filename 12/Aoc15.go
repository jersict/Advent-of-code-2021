package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var connections = make(map[string][]string)
var paths [][]string

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

func isLower(val string) bool {
	upper := strings.ToUpper(val)
	if val == upper {
		return false
	} else {
		return true
	}
}

func inSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func findPath(prev []string, next string, bonusVisitAvailable bool) {
	if next == "end" {
		prev = append(prev, next)
		paths = append(paths, prev)
	} else if isLower(next) {
		if !inSlice(prev, next) {
			prev = append(prev, next)
			for i := 0; i < len(connections[next]); i++ {
				findPath(prev, connections[next][i], bonusVisitAvailable)
			}
		} else if bonusVisitAvailable {
			bonusVisitAvailable = false
			prev = append(prev, next)
			for i := 0; i < len(connections[next]); i++ {
				findPath(prev, connections[next][i], bonusVisitAvailable)
			}
		}
	} else {
		prev = append(prev, next)
		for i := 0; i < len(connections[next]); i++ {
			findPath(prev, connections[next][i], bonusVisitAvailable)
		}
	}
}

func part_one() {
	fileTextLines := readInput()
	for _, eachline := range fileTextLines {
		line := strings.Split(eachline, "-")
		if line[1] != "start" {
			connections[line[0]] = append(connections[line[0]], line[1])
		}
		if line[0] != "start" {
			connections[line[1]] = append(connections[line[1]], line[0])
		}
	}
	for i := 0; i < len(connections["start"]); i++ {
		prev := []string{"start"}
		findPath(prev, connections["start"][i], false)
	}
	fmt.Printf("There are %d different paths out of the caves.\n\n", len(paths))
}

func part_two() {
	fileTextLines := readInput()
	for _, eachline := range fileTextLines {
		line := strings.Split(eachline, "-")
		if line[1] != "start" {
			connections[line[0]] = append(connections[line[0]], line[1])
		}
		if line[0] != "start" {
			connections[line[1]] = append(connections[line[1]], line[0])
		}
	}
	for i := 0; i < len(connections["start"]); i++ {
		prev := []string{"start"}
		findPath(prev, connections["start"][i], true)
	}
	fmt.Printf("There are %d different paths out of the caves.\n\n", len(paths))
}
func main() {
	//	part_one()
	part_two()
}
