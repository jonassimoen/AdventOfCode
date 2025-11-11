package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

func parseInput(s []string) []int {
	var ints []int
	for _, c := range s {
		i, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func part1(ls []string) int {
	lsi := parseInput(ls)
	pi := math.MinInt
	c := 0
	for _, i := range lsi {
		if i > pi {
			//fmt.Printf("Increasing: from %d to %d\n", pi, i)
			c += 1
		}
		pi = i
	}
	return c - 1
}

func part2(ls []string) int {
	lsi := parseInput(ls)
	psi := math.MinInt
	c := 0
	for i := 0; i < len(lsi)-2; i++ {
		si := lsi[i] + lsi[i+1] + lsi[i+2]
		if si > psi {
			//fmt.Printf("Increasing: from %d to %d\n", si, i)
			c += 1
		}
		psi = si
	}
	return c - 1
}
