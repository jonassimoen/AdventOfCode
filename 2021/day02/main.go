package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func processFile(fname string) {
	fmt.Printf("Input file: %s\n", fname)
	ls, err := readFile(fname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	out := part1(ls)
	fmt.Printf("Output P1: %d\n", out)
	out = part2(ls)
	fmt.Printf("Output P2: %d\n", out)
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

func parseStep(action string, operand int) (int, int) {
	switch action {
	case "forward":
		return operand, 0
	case "down":
		return 0, operand
	case "up":
		return 0, -operand
	}
	return 0, 0
}

func part1(ls []string) int {
	x, y := 0, 0
	for _, line := range ls {
		lineParsed := strings.Split(line, " ")
		operand, err := strconv.Atoi(lineParsed[1])
		if err != nil {
			panic(err)
		}
		dx, dy := parseStep(lineParsed[0], operand)
		x += dx
		y += dy
	}
	return x * y
}

func parseStepBis(action string, operand, currentAim int) (int, int, int) {
	switch action {
	case "forward":
		return operand, currentAim * operand, 0
	case "down":
		return 0, 0, operand
	case "up":
		return 0, 0, -operand
	}
	return 0, 0, 0
}

func part2(ls []string) int {
	x, y, aim := 0, 0, 0
	for _, line := range ls {
		lineParsed := strings.Split(line, " ")
		operand, err := strconv.Atoi(lineParsed[1])
		if err != nil {
			panic(err)
		}
		dx, dy, dAim := parseStepBis(lineParsed[0], operand, aim)
		x += dx
		y += dy
		aim += dAim
	}
	return x * y
}
