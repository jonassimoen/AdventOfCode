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

func main() {
	part1ptr := flag.Int("p", 1, "Part")
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Part %d - Input file: %s\n", *part1ptr, *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	var out int
	if *part1ptr == 1 {
		out = part1(ls)
	} else {
		out = part2(ls)
	}
	fmt.Printf("Output: %d\n", out)
}

func parseRange(rangeStr string) (int, int) {
	rangeSlice := strings.Split(rangeStr, "-")
	a, err := strconv.Atoi(rangeSlice[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(rangeSlice[1])
	if err != nil {
		panic(err)
	}
	return a, b
}

func part1(ls []string) int {
	count := 0
	for _, l := range ls {
		elves := strings.Split(l, ",")
		sA, eA := parseRange(elves[0])
		sB, eB := parseRange(elves[1])

		//s := max(min(sA, sB)-1, 0)
		//e := max(eA, eB) + 1

		if ((sA <= sB) && (eA >= eB)) || ((sA >= sB) && (eA <= eB)) {
			count++
			//fmt.Println(">> Overlap")
		}
	}
	return count
}

func part2(ls []string) int {
	count := 0
	for _, l := range ls {
		elves := strings.Split(l, ",")
		sA, eA := parseRange(elves[0])
		sB, eB := parseRange(elves[1])
		//fmt.Printf("Elf A: %d %d\n", sA, eA)
		//fmt.Printf("Elf B: %d %d\n", sB, eB)
		if ((sA <= sB) && (eA >= sB)) ||
			((sA <= eB) && (eA >= eB)) ||
			((sB <= sA) && (eB >= sA)) ||
			((sB <= eA) && (eB >= eA)) {
			//fmt.Printf("Overlap 4\n")
			count++
		}
	}
	return count
}
