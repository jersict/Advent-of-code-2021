package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (fileTextLines [][]string) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, strings.Split(fileScanner.Text(), ""))
	}

	readFile.Close()
	return
}

func part_one() {
	gamma := ""
	epsilon := ""
	fileTextLines := readInput()
	for i := 0; i < len(fileTextLines[0]); i++ {
		ones := 0
		zeros := 0
		for j := 0; j < len(fileTextLines); j++ {
			if fileTextLines[j][i] == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			epsilon += "0"
			gamma += "1"
		}

	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Power consumption of the submarine is %d.\n", g*e)
}

func determine_most_common(list [][]string, i int) (s string) {
	ones := 0
	zeros := 0
	for j := 0; j < len(list); j++ {
		if list[j][i] == "0" {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return "0"
	} else {
		return "1"
	}
}

func determine_least_common(list [][]string, i int) (s string) {
	ones := 0
	zeros := 0
	for j := 0; j < len(list); j++ {
		if list[j][i] == "0" {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return "1"
	} else {
		return "0"
	}
}

func remove(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}

func keepThis(list [][]string, s string, j int) [][]string {
	for i := 0; i < len(list); i++ {
		if list[i][j] != s {
			list = remove(list, i)
			i--
		}
	}
	return list
}

func find_next_oxy(oxy_left [][]string, i int) [][]string {
	keep := determine_most_common(oxy_left, i)
	oxy_left = keepThis(oxy_left, keep, i)
	if len(oxy_left) == 1 {
		return oxy_left
	} else {
		oxy_left := find_next_oxy(oxy_left, i+1)
		return oxy_left
	}
}

func find_next_co2(co2_left [][]string, i int) [][]string {
	keep := determine_least_common(co2_left, i)
	co2_left = keepThis(co2_left, keep, i)
	if len(co2_left) == 1 {
		return co2_left
	} else {
		co2_left := find_next_co2(co2_left, i+1)
		return co2_left
	}
}

func part_two() {
	oxy_gen_rat := ""
	co2_scr_rat := ""
	oxy_left := readInput()
	co2_left := readInput()
	oxy := find_next_oxy(oxy_left, 0)
	co2 := find_next_co2(co2_left, 0)
	for i := 0; i < len(oxy[0]); i++ {
		oxy_gen_rat += oxy[0][i]
		co2_scr_rat += co2[0][i]
	}
	g, err := strconv.ParseInt(oxy_gen_rat, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	e, err := strconv.ParseInt(co2_scr_rat, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Life support rating of the submarine is %d.\n", g*e)
}
func main() {
	part_one()
	part_two()
}
