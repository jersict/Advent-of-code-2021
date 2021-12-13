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

var heightMap [102][102]int

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

func getMap() {
	fileTextLines := readInput()
	for i := 0; i < len(fileTextLines); i++ {
		numbers := strings.Split(fileTextLines[i], "")
		for j := 0; j < len(numbers); j++ {
			num, err := strconv.Atoi(numbers[j])
			if err != nil {
				os.Exit(1)
			}
			heightMap[i+1][j+1] = num
		}
	}
	for i := 0; i < 102; i++ {
		heightMap[0][i] = 9
		heightMap[i][0] = 9
		heightMap[101][i] = 9
		heightMap[i][101] = 9
	}
}

func checkSurround(i, j int) (check int) {
	check = heightMap[i][j] + 1
	if heightMap[i][j] >= heightMap[i][j-1] || heightMap[i][j] >= heightMap[i][j+1] || heightMap[i][j] >= heightMap[i-1][j] || heightMap[i][j] >= heightMap[i+1][j] {
		check = 0
	}
	return
}

func checkAround(startx, starty int) (size int) {
	size = 1
	heightMap[starty][startx] = 9
	if heightMap[starty+1][startx] != 9 {
		size += checkAround(startx, starty+1)
	}
	if heightMap[starty-1][startx] != 9 {
		size += checkAround(startx, starty-1)
	}
	if heightMap[starty][startx+1] != 9 {
		size += checkAround(startx+1, starty)
	}
	if heightMap[starty][startx-1] != 9 {
		size += checkAround(startx-1, starty)
	}
	return
}

func getSize(starty, startx int) (size int) {
	size = 0
	size += checkAround(starty, startx)
	return
}

func checkForLowPoint(i, j int) (size int) {
	if heightMap[i][j] >= heightMap[i][j-1] || heightMap[i][j] >= heightMap[i][j+1] || heightMap[i][j] >= heightMap[i-1][j] || heightMap[i][j] >= heightMap[i+1][j] {
		size = 0
	} else {
		size = getSize(j, i)
	}
	return
}

func part_one() {
	count := 0
	getMap()
	heightMap[0][0] = 10
	for i := 1; i < len(heightMap)-1; i++ {
		for j := 1; j < len(heightMap[i])-1; j++ {
			count += checkSurround(i, j)
		}
	}
	fmt.Printf("There is %d low points\n", count)
}

func part_two() {
	sizes := []int{}
	size := 0
	getMap()
	for i := 1; i < len(heightMap)-1; i++ {
		for j := 1; j < len(heightMap[i])-1; j++ {
			size = checkForLowPoint(i, j)
			if size != 0 {
				sizes = append(sizes, size)
			}
		}
	}
	sort.Ints(sizes)

	fmt.Printf("Three largest sizes multiplied together equal to %d", sizes[len(sizes)-1]*sizes[len(sizes)-2]*sizes[len(sizes)-3])
}
func main() {
	part_one()
	part_two()
}
