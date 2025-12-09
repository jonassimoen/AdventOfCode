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

func parseInput(ls []string) [][]int {
	var grid [][]int
	for _, l := range ls {
		var row []int
		for _, r := range l {
			if r == '@' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]int) {
	for _, r := range grid {
		for _, c := range r {
			switch c {
			case 0:
				fmt.Printf(".")
			case 1:
				fmt.Printf("@")
			case 2:
				fmt.Printf("x")
			}
		}
		fmt.Println()
	}
}

func part1(ls []string) int {
	g := parseInput(ls)
	cnt := 0
	for ri, r := range g {
		for ci, _ := range r {
			if accessible(g, ri, ci) {
				g[ri][ci] = 2
				cnt++
			}
		}
	}
	return cnt
}

func isValid(g [][]int, r, c int) bool {
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
}

func accessible(g [][]int, r, c int) bool {
	if g[r][c] == 0 {
		return false
	}
	nbs := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
	accessibleCount := 0
	for _, nb := range nbs {
		nr, nc := r+nb[0], c+nb[1]
		if isValid(g, nr, nc) && g[nr][nc] >= 1 {
			accessibleCount++
		}
	}
	return accessibleCount < 4
}

func part2(ls []string) int {
	g := parseInput(ls)
	removedRolls := 1
	totalRemoved := 0
	for removedRolls > 0 {
		removedRolls = 0
		for ri, r := range g {
			for ci, _ := range r {
				if accessible(g, ri, ci) {
					g[ri][ci] = 2
					removedRolls++
				}
			}
		}
		//println()
		//printGrid(g)
		for ri, r := range g {
			for ci, _ := range r {
				if g[ri][ci] == 2 {
					g[ri][ci] = 0
				}
			}
		}
		totalRemoved += removedRolls
	}
	return totalRemoved
}
