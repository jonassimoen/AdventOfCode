package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

func parseInput(ls []string) [][]int {
	var grid [][]int
	for _, l := range ls {
		var r []int
		for _, c := range l {
			r = append(r, int(c-'0'))
		}
		grid = append(grid, r)
	}
	return grid
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

type Coordinate [2]int

func isLowestPoint(grid [][]int, c Coordinate) bool {
	h := grid[c[1]][c[0]]
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, d := range dirs {
		nx, ny := c[0]+d[0], c[1]+d[1]
		if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
			continue
		}
		if grid[ny][nx] <= h {
			return false
		}
	}
	return true
}

func part1(ls []string) int {
	g := parseInput(ls)
	riskLevels := 0
	for r, gr := range g {
		for c, v := range gr {
			if isLowestPoint(g, Coordinate{c, r}) {
				riskLevels += v + 1
			}
		}
	}
	return riskLevels
}

func neighbors(g [][]int, c Coordinate, value int) []Coordinate {
	var ns []Coordinate
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, dir := range dirs {
		nx, ny := c[0]+dir[0], c[1]+dir[1]
		if nx < 0 || nx >= len(g[0]) || ny < 0 || ny >= len(g) {
			continue
		}
		if g[ny][nx] <= value {
			continue
		}
		ns = append(ns, Coordinate{nx, ny})
	}
	return ns
}

func basinContains(basin []Coordinate, c Coordinate) bool {
	for _, b := range basin {
		if b == c {
			return true
		}
	}
	return false
}

func findBasin(g [][]int, c Coordinate, basin map[Coordinate]bool) int {
	if basin[c] {
		return 0
	}
	if g[c[1]][c[0]] == 9 {
		return 0
	}

	basin[c] = true
	basinLen := 1

	nbs := neighbors(g, c, g[c[1]][c[0]])
	for _, nb := range nbs {
		_, ok := basin[nb]
		if ok || g[nb[1]][nb[0]] == 9 {
			continue
		}
		basinLen += findBasin(g, nb, basin)
	}
	return basinLen
}

func processLowPoints(g [][]int, lowPoints []Coordinate) int {
	basins := []int{}
	for _, p := range lowPoints {
		basin := findBasin(g, p, map[Coordinate]bool{})
		basins = append(basins, basin)
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})
	v := 1
	for _, b := range basins[:3] {
		fmt.Println(b)
		v *= b
	}
	return v
}

func printBasin(g [][]int, basin []Coordinate) {
	fmt.Println(len(basin))
	for r, gr := range g {
		for c, v := range gr {
			if basinContains(basin, Coordinate{c, r}) {
				fmt.Printf("#")
			} else {
				fmt.Printf("%d", v)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func part2(ls []string) int {
	g := parseInput(ls)
	var lowPoints []Coordinate
	for r, gr := range g {
		for c := range gr {
			coord := Coordinate{c, r}
			if isLowestPoint(g, coord) {
				lowPoints = append(lowPoints, coord)
			}
		}
	}
	return processLowPoints(g, lowPoints)
}
