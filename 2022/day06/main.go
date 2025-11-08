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

func hasDifferentCharacters(s string, n int) bool {
	if len(s) != n {
		return false
	}
	a := map[rune]bool{}
	for _, r := range s {
		if a[r] {
			return false
		}
		a[r] = true
	}
	return true
}

func validateMarker(s string, occurences map[rune]bool) bool {
	for _, si := range s {
		if _, ok := occurences[si]; !ok {
			fmt.Printf("\t> [%c] first occurence\n", si)
			return false
		}
	}
	return true
}

func printMap(m map[rune]bool) {
	for v, _ := range m {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}

func retrieveMarker(l string, distinctChars int) int {
	occuredChars := map[rune]bool{}
	for i := range distinctChars {
		occuredChars[rune(l[i])] = true
	}
	idx := distinctChars
	for idx < len(l) {
		possibleMarker := l[idx-distinctChars : idx]
		//fmt.Println(possibleMarker)
		//printMap(occuredChars)
		if hasDifferentCharacters(possibleMarker, distinctChars) && validateMarker(possibleMarker, occuredChars) {
			fmt.Println("\tMARKER", idx)
			return idx
		}
		occuredChars[rune(l[idx])] = true
		idx++
	}
	return -1
}

func part1(ls []string) int {
	val := 0
	for _, l := range ls {
		val = retrieveMarker(l, 4)
	}
	return val
}

func part2(ls []string) int {
	val := 0
	for _, l := range ls {
		val = retrieveMarker(l, 14)
	}
	return val
}
