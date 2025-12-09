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

func parseInput(ls []string) ([]int, []int) {
	var start, end []int

	for _, l := range ls {
		splits := strings.Split(l, ",")
		for _, split := range splits {
			splitSplit := strings.Split(split, "-")
			a, _ := strconv.Atoi(splitSplit[0])
			b, _ := strconv.Atoi(splitSplit[1])
			start = append(start, a)
			end = append(end, b)
		}
	}
	return start, end

}

func part1(ls []string) int {
	ids := 0
	s, e := parseInput(ls)
	for idx := 0; idx < len(s); idx++ {
		start, end := s[idx], e[idx]
		for v := start; v <= end; v++ {
			s := strconv.Itoa(v)
			length := len(s)
			if s[:length/2] == s[length/2:] {
				ids += v
			}
		}
	}
	return ids
}

func part2(ls []string) int {
	ids := 0
	s, e := parseInput(ls)
	for idx := 0; idx < len(s); idx++ {
		start, end := s[idx], e[idx]
		for v := start; v <= end; v++ {
			if isInvalidId(v) {
				ids += v
			}
		}
	}
	return ids
}

func isInvalidId(v int) bool {
	s := strconv.Itoa(v)
	length := len(s)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}
		//fmt.Printf("[%d] Checking length %d (%d parts)\n", v, i, length/i)
		invalid := true
		for r := 1; r <= length/i; r++ {
			//fmt.Printf("> [%d] from idx %d -> %d: %s ?= %s\n", v, (r-1)*i, r*i, s[0:i], s[(r-1)*i:r*i])
			invalid = invalid && (s[0:i] == s[(r-1)*i:r*i])
		}
		if invalid {
			//fmt.Printf("[%d] > Repeating %s\n", v, s[0:i])
			return true
		}
	}
	return false
}
