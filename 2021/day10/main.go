package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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

func isOpen(r rune) bool {
	return r == '(' || r == '[' || r == '{' || r == '<'
}

func isClose(r rune) bool {
	return r == ')' || r == ']' || r == '}' || r == '>'
}

func endRune(r rune) rune {
	switch r {
	case '{':
		return '}'
	case '(':
		return ')'
	case '<':
		return '>'
	case '[':
		return ']'
	}
	return 'X'
}

func invalidArgumentScore(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '>':
		return 25137
	case '}':
		return 1197
	}
	return 0
}

func runesSliceToString(rs []rune) string {
	var sb strings.Builder
	for _, r := range rs {
		sb.WriteRune(r)
	}
	return sb.String()
}

func processString(s string, endRunes []rune) (int, string) {
	if len(s) == 0 {
		return -1, string(endRunes)
	}
	r := rune(s[0])
	//fmt.Printf("Processing c=%c, s=%s, e=%s\n", r, s, runesSliceToString(endRunes))

	if isOpen(r) {
		rE := endRune(r)
		//fmt.Printf("> [OPEN] Adding %c to endRunes\n", rE)
		return processString(s[1:], append([]rune{rE}, endRunes...))
	} else if isClose(r) {
		expectingRune := endRunes[0]
		//fmt.Printf("> [CLOSE] Expecting %c, got %c\n", expectingRune, s[0])
		if rune(s[0]) != expectingRune {
			//fmt.Printf("Invalid character: %c instead of %c\n", r, expectingRune)
			return invalidArgumentScore(rune(s[0])), ""
		} else {
			//fmt.Printf("Valid: proceeding\n")
			return processString(s[1:], endRunes[1:])
		}
	}
	return 0, ""
}

func part1(ls []string) int {
	total := 0
	for _, l := range ls {
		stringScore, _ := processString(l, []rune{})
		if stringScore != -1 {
			total += stringScore
		}

	}
	return total
}

func autoCompletePoints(r rune) int {
	switch r {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func calculateScore(rs string) int {
	score := 0
	for _, r := range rs {
		score *= 5
		score += autoCompletePoints(r)
	}
	return score
}

func part2(ls []string) int {
	totals := []int{}
	for _, l := range ls {
		stringScore, toComplete := processString(l, []rune{})
		if stringScore == -1 {
			//fmt.Printf("Incomplete: %s\n", toComplete)
			totals = append(totals, calculateScore(toComplete))
		}

	}
	slices.Sort(totals)
	return totals[len(totals)/2]
}
