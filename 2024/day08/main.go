package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

type Coordinate struct {
	x, y int
}

func parseAntennas(ls []string) (map[rune][]Coordinate, int, int) {
	antennas := make(map[rune][]Coordinate)
	maxY, maxX := len(ls), len(ls[0])
	for yi, l := range ls {
		for xi, lv := range l {
			if lv != '.' {
				if _, ok := antennas[lv]; !ok {
					antennas[lv] = make([]Coordinate, 0)
				}
				antennas[lv] = append(antennas[lv], Coordinate{xi, yi})
			}
		}
	}
	return antennas, maxX, maxY
}

func distance(a Coordinate, b Coordinate) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
}

func findLineLocations(a Coordinate, b Coordinate, maxX, maxY int) []Coordinate {
	coord := make([]Coordinate, 0)
	m := float64(b.y-a.y) / float64(b.x-a.x)
	for x := 0; x < maxX; x++ {
		yf := m*float64(x-a.x) + float64(a.y)
		if math.Floor(yf) != math.Floor(yf*100)/float64(100) {
			continue
		}
		y := int(yf)
		if ((y < maxY) && (y >= 0)) && ((x != a.x) || (y != a.y)) && ((x != b.x) || (y != b.y)) {
			coord = append(coord, Coordinate{x, y})
		}
	}
	return coord
}

func isAntinode(c Coordinate, a Coordinate, b Coordinate) bool {
	distA := distance(c, a)
	distB := distance(c, b)
	return ((distB == 2*distA) || (distA == 2*distB))
}

func printArray(antennas map[rune][]Coordinate, antinodes []Coordinate, maxX, maxY int) {
	prArr := make([][]rune, 0)
	for yi := 0; yi < maxY; yi++ {
		prArr = append(prArr, make([]rune, 0))
		for xi := 0; xi < maxX; xi++ {
			prArr[yi] = append(prArr[yi], '.')
		}
	}

	for a, csA := range antennas {
		for _, cA := range csA {
			prArr[cA.y][cA.x] = a
		}
	}

	for _, an := range antinodes {
		if prArr[an.y][an.x] == '.' {
			prArr[an.y][an.x] = '#'
		}
	}

	for yi := 0; yi < maxY; yi++ {
		for xi := 0; xi < maxX; xi++ {
			fmt.Printf("%c", prArr[yi][xi])
		}
		fmt.Println()
	}
}

func isElementOfArray(c Coordinate, cs []Coordinate) bool {
	for _, ci := range cs {
		if ci.x == c.x && ci.y == c.y {
			return true
		}
	}
	return false
}

func part1(ls []string) int {
	ants, x, y := parseAntennas(ls)
	//printArray(ants, []Coordinate{{x: 0, y: 0}}, x, y)
	an := make([]Coordinate, 0)
	for _, v := range ants {
		for ai, antA := range v {
			for aii := ai + 1; aii < len(v); aii++ {
				antB := v[aii]
				cs := findLineLocations(antA, antB, x, y)
				//fmt.Println(cs)
				for _, c := range cs {
					if isAntinode(c, antA, antB) && !isElementOfArray(c, an) {
						fmt.Println(c.x, c.y)
						an = append(an, c)
					}
				}
			}
		}
	}
	printArray(ants, an, x, y)
	return len(an)
}

func part2(ls []string) int {
	return -1
}
