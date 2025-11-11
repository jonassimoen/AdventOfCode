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

var SegmentMap map[int]string = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

type Entry struct {
	in  []string
	out []string
}

func parseInput(ls []string) []Entry {
	var entries []Entry
	for _, l := range ls {
		e := Entry{}
		lStripped := strings.Split(l, " | ")
		e.in = strings.Split(lStripped[0], " ")
		e.out = strings.Split(lStripped[1], " ")
		entries = append(entries, e)
	}
	return entries
}

func part1(ls []string) int {
	e := parseInput(ls)
	uniqueSegments := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}

	for _, e := range e {
		for _, d := range e.out {
			for uq := range uniqueSegments {
				if len(d) == len(SegmentMap[uq]) {
					uniqueSegments[uq] += 1
					break
				}
			}
		}
	}

	return uniqueSegments[1] + uniqueSegments[4] + uniqueSegments[7] + uniqueSegments[8]
}

func isSubset(a []rune, b string) bool {
	for _, r := range a {
		if !strings.ContainsRune(b, r) {
			return false
		}
	}
	return true
}

func sortedString(a string) string {
	c := []rune(a)
	slices.Sort(c)
	return string(c)
}

func (e *Entry) process() int {
	uniqueSegments := map[int][]rune{
		1: []rune{},
		4: []rune{},
		7: []rune{},
		8: []rune{},
	}
	fiveSeg, sixSeg := map[string]bool{}, map[string]bool{}
	for _, d := range e.in {
		for uq, rs := range uniqueSegments {
			if len(d) == len(SegmentMap[uq]) && len(rs) == 0 {
				uniqueSegments[uq] = []rune(sortedString(d))
				break
			}
		}
		if len(d) == 5 {
			if _, ok := fiveSeg[sortedString(d)]; !ok {
				fiveSeg[sortedString(d)] = true
			}
		}
		if len(d) == 6 {
			if _, ok := sixSeg[sortedString(d)]; !ok {
				sixSeg[sortedString(d)] = true
			}
		}
	}

	// 3 = five segments, fully contains 1
	for five, ok := range fiveSeg {
		if !ok {
			continue
		}
		if isSubset(uniqueSegments[1], five) {
			//fmt.Printf("Found 3: %s\n", five)
			uniqueSegments[3] = []rune(five)
			fiveSeg[five] = false
			break
		}
	}

	// 9 = six segments, fully contains 4
	for six, ok := range sixSeg {
		if !ok {
			continue
		}
		if isSubset(uniqueSegments[4], six) {
			//fmt.Printf("Found 9: %s\n", six)
			uniqueSegments[9] = []rune(six)
			sixSeg[six] = false
			break
		}
	}

	// 0 = six segments, fully contains 1, not 9 (already found)
	for six, ok := range sixSeg {
		if !ok {
			continue
		}
		if isSubset(uniqueSegments[1], six) {
			//fmt.Printf("Found 0: %s\n", six)
			uniqueSegments[0] = []rune(six)
			sixSeg[six] = false
			break
		}
	}

	// 6 = last of the six segments
	for six, ok := range sixSeg {
		if ok {
			//fmt.Printf("Found 6: %s\n", six)
			uniqueSegments[6] = []rune(six)
			sixSeg[six] = false
		}
	}

	// 5 = five segments, fully contained in 6
	for five, ok := range fiveSeg {
		if !ok {
			continue
		}
		if isSubset([]rune(five), string(uniqueSegments[6])) {
			//fmt.Printf("Found 5: %s\n", five)
			uniqueSegments[5] = []rune(five)
			fiveSeg[five] = false
		}
	}

	// 2 = last one of the five segments
	for five, ok := range fiveSeg {
		if ok {
			//fmt.Printf("Found 2: %s\n", five)
			uniqueSegments[2] = []rune(five)
		}
	}

	mapSegments := map[string]int{}
	for i, s := range uniqueSegments {
		mapSegments[string(s)] = i
	}

	sum := 0
	for _, d := range e.out {
		dSorted := sortedString(d)
		v := mapSegments[dSorted]
		sum = sum*10 + v
	}
	fmt.Printf("%v -> %d\n", e.out, sum)
	return sum
}

func part2(ls []string) int {
	es := parseInput(ls)
	totalSum := 0
	for _, e := range es {
		totalSum += e.process()
	}
	return totalSum
}
