package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

type Towels struct {
	available     []string
	patterns      []string
	possibilities map[string]int
	cache         map[string]bool
}

func parseInput(ls []string) *Towels {
	t := Towels{
		available:     strings.Split(ls[0], ", "),
		patterns:      ls[2:],
		possibilities: make(map[string]int),
		cache:         make(map[string]bool),
	}

	return &t
}

func part1(ls []string) int {
	t := parseInput(ls)

	count := 0
	for _, pattern := range t.patterns {
		if t.checkPattern(pattern) {
			count++
		}
	}
	return count
}

func (t *Towels) checkPattern(towel string) bool {
	if result, exists := t.cache[towel]; exists {
		return result
	}
	if len(towel) == 0 {
		return true
	}
	for _, pattern := range t.available {
		if strings.HasPrefix(towel, pattern) {
			if t.checkPattern(towel[len(pattern):]) {
				t.cache[towel] = true
				return true
			}
		}
	}

	t.cache[towel] = false
	return false
	//fmt.Printf("%s[PART] %s NOT FOUND\n", prefix, towel)
}

func part2(ls []string) int {
	t := parseInput(ls)

	count := 0
	for _, pattern := range t.patterns {
		count += t.countPatterns(pattern)
	}
	return count
}

func (t *Towels) countPatterns(towel string) int {
	if result, exists := t.possibilities[towel]; exists {
		return result
	}
	if len(towel) == 0 {
		return 1
	}
	combinations := 0
	for _, pattern := range t.available {
		if strings.HasPrefix(towel, pattern) {
			combinations += t.countPatterns(towel[len(pattern):])
		}
	}

	t.possibilities[towel] = combinations
	return combinations
	//fmt.Printf("%s[PART] %s NOT FOUND\n", prefix, towel)
}
