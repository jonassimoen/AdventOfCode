package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

type Inventory struct {
	ranges      [][]int
	ingredients []int
}

func (i *Inventory) reduce() {
	newRanges := [][]int{}
	ranges := i.ranges
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	tempStart, tempEnd := ranges[0][0], ranges[0][1]
	for _, r := range ranges[1:] {
		start, end := r[0], r[1]
		if start <= tempEnd {
			tempEnd = max(end, tempEnd)
		} else {
			newRanges = append(newRanges, []int{tempStart, tempEnd})
			tempStart, tempEnd = start, end
		}
	}
	newRanges = append(newRanges, []int{tempStart, tempEnd})
	for _, r := range newRanges {
		fmt.Printf("%d %d\n", r[0], r[1])
	}
	i.ranges = newRanges
}

func parseInput(ls []string) Inventory {
	i := Inventory{
		ranges:      [][]int{},
		ingredients: []int{},
	}
	fresh := true
	for _, l := range ls {
		if l == "" {
			fresh = false
		}
		if fresh {
			split := strings.Split(l, "-")
			a, _ := strconv.Atoi(split[0])
			b, _ := strconv.Atoi(split[1])
			i.ranges = append(i.ranges, []int{a, b})
		} else {
			x, _ := strconv.Atoi(l)
			i.ingredients = append(i.ingredients, x)
		}
	}
	//i.print()
	//println("--")
	i.reduce()
	//i.print()
	return i
}

func (i *Inventory) print() {
	for _, r := range i.ranges {
		fmt.Printf("%d-%d\n", r[0], r[1])
	}
	fmt.Println()
}

func part1(ls []string) int {
	inv := parseInput(ls)
	total := 0
	for _, i := range inv.ingredients {
		for _, r := range inv.ranges {
			if i >= r[0] && i <= r[1] {
				total += 1
				break
			}
		}
	}
	return total
}

func part2(ls []string) int {
	inv := parseInput(ls)
	total := 0
	for _, r := range inv.ranges {
		total += r[1] - r[0] + 1
	}
	return total
}
