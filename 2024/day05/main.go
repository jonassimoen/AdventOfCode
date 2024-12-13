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

func pagesToCheck(shouldBePrinted []int, updatePages []int) []int {
	ret := []int{}
	for _, pageShouldBe := range shouldBePrinted {
		if inArray(pageShouldBe, updatePages) {
			ret = append(ret, pageShouldBe)
		}
	}
	return ret
}

func inArray(val int, array []int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

func parseUpdates(update string, order *map[int][]int) (bool, []int) {
	pages := []int{}
	printed := []int{}
	splits := strings.Split(update, ",")
	for _, split := range splits {
		val, _ := strconv.Atoi(split)
		pages = append(pages, val)
	}

	for _, page := range pages {
		if (*order)[page] == nil {
			printed = append(printed, page)
		} else {
			depsBeforeOk := true
			for _, p := range pagesToCheck((*order)[page], pages) {
				if !inArray(p, printed) {
					depsBeforeOk = false
				}
			}
			if depsBeforeOk {
				printed = append(printed, page)
			}
		}
	}
	return len(printed) == len(pages), pages
}

func parsePrintedAfterToPrintedBefore(printedAfter *map[int][]int) *map[int][]int {
	var printedBefore map[int][]int
	printedBefore = make(map[int][]int)
	for before, after := range *printedAfter {
		for _, valAfter := range after {

			if printedBefore[valAfter] == nil {
				printedBefore[valAfter] = []int{}
			}
			printedBefore[valAfter] = append(printedBefore[valAfter], before)
		}
	}

	return &printedBefore
}

func part1(ls []string) int {
	sum := 0
	printedAfter := make(map[int][]int)
	var printedBefore *map[int][]int
	split := false
	for _, l := range ls {
		if l == "" {
			split = true
			printedBefore = parsePrintedAfterToPrintedBefore(&printedAfter)
		} else {
			if !split {
				split := strings.Split(l, "|")
				a, _ := strconv.Atoi(split[0])
				b, _ := strconv.Atoi(split[1])
				if _, ok := printedAfter[a]; !ok {
					printedAfter[a] = make([]int, 0)
				}
				printedAfter[a] = append(printedAfter[a], b)
			} else {
				isPrinted, printedPages := parseUpdates(l, printedBefore)
				if isPrinted {
					sum += printedPages[len(printedPages)/2]
				}
			}
		}
	}
	return sum
}

func reorderPages(pages []int, order map[int][]int) []int {
	//fmt.Printf("Checking %v\n", pages)
	idx := 0
	for idx < len(pages) {
		p := pages[idx]
		//fmt.Printf("Page %d:\n", p)
		idxChecked := false
		for _, pageBefore := range order[p] {
			if !idxChecked && inArray(pageBefore, pages) && !inArray(pageBefore, pages[:idx]) {
				//fmt.Printf("\t%d should be before %d: %v ==>", pageBefore, p, pages)
				pages = reorderArray(pages, idx, pageBefore)
				idxChecked = true
				//fmt.Println(pages, "idx=", idx)
				break
			}
		}
		if !idxChecked {
			idx++
		}
		//fmt.Printf("Checked page %d (found wrong: %t): idx=%v\n", p, idxChecked, idx)
	}
	return pages
}

func reorderArray(pages []int, idxA, newElement int) []int {
	var n []int
	n = append(n, pages[:idxA]...)
	n = append(n, newElement)
	for _, v := range pages[idxA:] {
		if v != newElement {
			n = append(n, v)
		}
	}
	return n
}

func part2(ls []string) int {
	sum := 0
	printedBefore := make(map[int][]int)
	split := false
	for _, l := range ls {
		if l == "" {
			split = true
		} else {
			if !split {
				split := strings.Split(l, "|")
				a, _ := strconv.Atoi(split[1])
				b, _ := strconv.Atoi(split[0])
				if _, ok := printedBefore[a]; !ok {
					printedBefore[a] = make([]int, 0)
				}
				printedBefore[a] = append(printedBefore[a], b)
			} else {
				isPrinted, printedPages := parseUpdates(l, &printedBefore)
				if !isPrinted {
					t := reorderPages(printedPages, printedBefore)
					sum += t[len(t)/2]

				}
			}
		}
	}
	return sum
}
