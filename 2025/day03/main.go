package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readFile(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ls []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ls = append(ls, sc.Text())
	}
	return ls, sc.Err()
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Input file: %s\n", *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	out := part1(ls)
	fmt.Printf("Output P1: %d\n", out)
	out = part2(ls)
	fmt.Printf("Output P2: %d\n", out)
}

func parseInput(ls []string) [][]int {
	var grid [][]int
	for _, l := range ls {
		var row []int
		for _, c := range l {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}

func part1(ls []string) int {
	out := 0
	g := parseInput(ls)
	for _, row := range g {
		maxIdx, secIdx := -1, -1
		for idx := 0; idx < len(row)-1; idx++ {
			if (maxIdx == -1) || (row[maxIdx] < row[idx]) {
				maxIdx = idx
			}
		}
		for idx := maxIdx + 1; idx <= len(row)-1; idx++ {
			if (secIdx == -1) || (row[secIdx] < row[idx]) {
				secIdx = idx
			}
		}
		fmt.Printf("> %d\n", row[maxIdx]*10+row[secIdx])
		out += row[maxIdx]*10 + row[secIdx]
	}
	return out
}

func part2(ls []string) int {
	out := 0
	g := parseInput(ls)
	for _, row := range g {
		rowJoltage := 0
		prevMaxIdx := -1
		for joltageIdx := 11; joltageIdx >= 0; joltageIdx-- {
			maxIdx := -1
			for idx := prevMaxIdx; idx < len(row)-joltageIdx; idx++ {
				if (maxIdx == -1) || (row[maxIdx] < row[idx]) {
					maxIdx = idx
				}
			}
			//fmt.Printf("> %d + [%d] %d ==> %d\n", rowJoltage, maxIdx, row[maxIdx], rowJoltage*10+row[maxIdx])
			rowJoltage = (rowJoltage * 10) + row[maxIdx]
			//fmt.Printf("\t%d\n", rowJoltage)
			prevMaxIdx = maxIdx + 1
		}
		fmt.Printf("> %d\n", rowJoltage)
		out += rowJoltage
	}
	return out
}
