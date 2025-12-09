package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

func part1(ls []string) int {
	zeroCount := 0
	idx := 50
	for _, l := range ls {
		r := l[0]
		i, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}
		rotation := 1
		if r == 'L' {
			rotation = -1
		}
		newIdx := (idx + rotation*i) % 100
		if newIdx < 0 {
			newIdx = 100 + newIdx
		}
		fmt.Printf("Move %c %d ==> %d\n", l[0], i, newIdx)
		idx = newIdx % 100

		if idx == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func part2(ls []string) int {
	zeroPasses := int(0)
	idx := 50
	for _, l := range ls {
		r := l[0]
		i, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}
		if r == 'L' {
			if idx == 0 {
				zeroPasses--
			}
			idx -= i
			zeroPasses -= floorDiv(idx, 100)
			if idx%100 == 0 {
				zeroPasses++
			}
		} else {
			idx += i
			zeroPasses += floorDiv(idx, 100)
		}
		idx = ((idx % 100) + 100) % 100
		fmt.Printf("Move %c %d ==> %d\n", l[0], i, idx)
		fmt.Printf("\tZero passes: %d\n", zeroPasses)
	}
	return zeroPasses
}

func floorDiv(a, b int) int {
	if (a < 0) != (b < 0) && a%b != 0 {
		return a/b - 1
	}
	return a / b
}
