package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

func convToInt(ls []string, i, j int) int {
	c, _ := strconv.Atoi(ls[i])
	return c
}

func checkVertical(i, j int, ls []string) int {
	vert := 0
	if j >= 3 {
		// Check above
		if (ls[j-1][i] == 'M') && (ls[j-2][i] == 'A') && (ls[j-3][i] == 'S') {
			vert++
		}
	}
	if j < (len(ls) - 3) {
		// Check below
		if (ls[j+1][i] == 'M') && (ls[j+2][i] == 'A') && (ls[j+3][i] == 'S') {
			vert++
		}
	}
	return vert
}

func checkDiagonal(i, j int, ls []string) int {
	diag := 0
	canGoLeft := i >= 3
	canGoRight := i < len(ls[0])-3
	canGoUp := j >= 3
	canGoDown := j < len(ls)-3

	if canGoLeft && canGoUp {
		if (ls[j-1][i-1] == 'M') && (ls[j-2][i-2] == 'A') && (ls[j-3][i-3] == 'S') {
			diag++
		}
	}
	if canGoRight && canGoDown {
		if (ls[j+1][i+1] == 'M') && (ls[j+2][i+2] == 'A') && (ls[j+3][i+3] == 'S') {
			diag++
		}
	}
	if canGoLeft && canGoDown {
		if (ls[j+1][i-1] == 'M') && (ls[j+2][i-2] == 'A') && (ls[j+3][i-3] == 'S') {
			diag++
		}
	}
	if canGoRight && canGoUp {
		if (ls[j-1][i+1] == 'M') && (ls[j-2][i+2] == 'A') && (ls[j-3][i+3] == 'S') {
			diag++
		}
	}
	return diag
}

func part1(ls []string) int {
	sum := 0
	for lj, l := range ls {
		regFw := regexp.MustCompile("XMAS")
		regBw := regexp.MustCompile("SAMX")
		mFw := regFw.FindAllString(l, -1)
		mBw := regBw.FindAllString(l, -1)
		sum += len(mBw) + len(mFw)
		for li, c := range l {
			if c == 'X' {
				sum += checkVertical(li, lj, ls)
				sum += checkDiagonal(li, lj, ls)
			}
		}
	}
	return sum
}
func checkDiagonalMAS(i, j int, ls []string) int {
	diag := 0
	lt := ls[j-1][i-1]
	rt := ls[j-1][i+1]
	ld := ls[j+1][i-1]
	rd := ls[j+1][i+1]
	if ls[j][i] == 'A' {
		if ((lt == 'M' && rd == 'S') || (lt == 'S' && rd == 'M')) &&
			((ld == 'M' && rt == 'S') || (ld == 'S' && rt == 'M')) {
			diag++
		}
	}
	return diag
}
func part2(ls []string) int {
	sum := 0
	for j := 1; j < len(ls)-1; j++ {
		for i := 1; i < len(ls[j])-1; i++ {
			sum += checkDiagonalMAS(j, i, ls)
		}
	}
	return sum
}
