package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

type Matrix struct {
	m                [][]int
	lenJ             int
	lenI             int
	zeros            []Coordinate
	trailHeads       map[Coordinate][]Coordinate
	trailHeadsRating map[Coordinate]int
}

type Coordinate struct {
	j, i int
}

func parseIput(ls []string) *Matrix {
	var m [][]int
	var zeros []Coordinate
	for j, l := range ls {
		m = append(m, []int{})
		for i, c := range l {
			if c == '.' {
				m[j] = append(m[j], 99)
				continue
			}
			n, _ := strconv.Atoi(string(c))
			if n == 0 {
				zeros = append(zeros, Coordinate{j, i})
			}
			m[j] = append(m[j], n)
		}
	}
	return &Matrix{
		m:                m,
		lenJ:             len(m),
		lenI:             len(m[0]),
		zeros:            zeros,
		trailHeads:       make(map[Coordinate][]Coordinate),
		trailHeadsRating: make(map[Coordinate]int),
	}
}

func (m *Matrix) checkNines() {
	for _, n := range m.zeros {
		m.runTrailReverse(n, n, 0, 0)
	}
}

func (m *Matrix) calculateScore() int {
	sum := 0
	for _, v := range m.trailHeads {
		sum += len(v)
	}
	return sum
}

func (m *Matrix) isInBounds(c Coordinate) bool {
	return ((c.j >= 0) && (c.j < m.lenJ)) && ((c.i >= 0) && (c.i < m.lenI))
}

func isInArray[T comparable](element T, array []T) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}
	return false
}

func (m *Matrix) runTrailReverse(originC Coordinate, c Coordinate, expectedValue int, encounteredZeros int) {
	//fmt.Printf("On my way to check: %v, expected value = %d, value = %d [%v]\n", c, expectedValue, m.m[c.j][c.i], m.trailHeads)
	if expectedValue != m.m[c.j][c.i] {
		return
	}
	if expectedValue == 9 {
		if m.trailHeads[c] == nil {
			m.trailHeads[c] = []Coordinate{}
		}
		if !isInArray(originC, m.trailHeads[c]) {
			m.trailHeads[c] = append(m.trailHeads[c], originC)
		}
		return
	}
	upper := Coordinate{c.j - 1, c.i}
	lower := Coordinate{c.j + 1, c.i}
	left := Coordinate{c.j, c.i - 1}
	right := Coordinate{c.j, c.i + 1}
	if m.isInBounds(upper) {
		m.runTrailReverse(originC, upper, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(lower) {
		m.runTrailReverse(originC, lower, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(left) {
		m.runTrailReverse(originC, left, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(right) {
		m.runTrailReverse(originC, right, expectedValue+1, encounteredZeros)
	}
	return
}

func part1(ls []string) int {
	m := parseIput(ls)
	m.checkNines()

	//fmt.Println(m.trailHeads)
	return m.calculateScore()
}

func (m *Matrix) runTrailReverseRating(originC Coordinate, c Coordinate, expectedValue int, encounteredZeros int) int {
	//fmt.Printf("On my way to check: %v, expected value = %d, value = %d [%v]\n", c, expectedValue, m.m[c.j][c.i], m.trailHeads)
	if expectedValue != m.m[c.j][c.i] {
		return 0
	}
	if expectedValue == 9 {
		return 1
	}
	upper := Coordinate{c.j - 1, c.i}
	lower := Coordinate{c.j + 1, c.i}
	left := Coordinate{c.j, c.i - 1}
	right := Coordinate{c.j, c.i + 1}
	sum := 0
	if m.isInBounds(upper) {
		sum += m.runTrailReverseRating(originC, upper, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(lower) {
		sum += m.runTrailReverseRating(originC, lower, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(left) {
		sum += m.runTrailReverseRating(originC, left, expectedValue+1, encounteredZeros)
	}
	if m.isInBounds(right) {
		sum += m.runTrailReverseRating(originC, right, expectedValue+1, encounteredZeros)
	}
	return sum
}

func part2(ls []string) int {
	m := parseIput(ls)

	sum := 0
	for _, n := range m.zeros {
		sum += m.runTrailReverseRating(n, n, 0, 0)
	}

	//fmt.Println(m.trailHeads)
	return sum
}
