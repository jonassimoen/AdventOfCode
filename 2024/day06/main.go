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

type Cell struct {
	blocked bool
	current bool
	visited bool
	special bool
}

type Direction [2]int

var UP = Direction{0, -1}
var DOWN = Direction{0, 1}
var LEFT = Direction{-1, 0}
var RIGHT = Direction{1, 0}

func parseDirection(r rune) (Direction, bool) {
	if r == '<' {
		return LEFT, true
	}
	if r == '>' {
		return RIGHT, true
	}
	if r == '^' {
		return UP, true
	}
	if r == 'v' {
		return DOWN, true
	}
	return Direction{0, 0}, false
}

func unParseDirection(d Direction) (rune, bool) {
	if d == LEFT {
		return '<', true
	}
	if d == RIGHT {
		return '>', true
	}
	if d == UP {
		return '^', true
	}
	if d == DOWN {
		return 'v', true
	}
	return '?', false
}

func parseInput(ls []string) ([][]Cell, Direction, Direction) {
	c := make([][]Cell, len(ls))
	var curDir Direction
	position := [2]int{0, 0}
	for j, l := range ls {
		c[j] = make([]Cell, len(l))
		for i, b := range l {
			c[j][i] = Cell{}
			if rune(b) == '#' {
				c[j][i].blocked = true
				continue
			}
			dir, found := parseDirection(rune(b))
			if found {
				c[j][i].current = true
				c[j][i].visited = true
				curDir = dir
				position = [2]int{i, j}
			}
		}
	}
	return c, curDir, position
}

func parseCell(c Cell, curDir Direction, dissplayCurrent bool) rune {
	if c.special {
		return 'O'
	}
	if c.blocked {
		return '#'
	}
	if c.current {
		if dissplayCurrent {
			r, _ := unParseDirection(curDir)
			return r
		} else {
			return 'X'
		}
	}
	if c.visited {
		return 'X'
	}
	return '.'
}

func printCells(c [][]Cell, curDir Direction, displayCurrent bool) {
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			cell := c[i][j]
			fmt.Printf("%c", parseCell(cell, curDir, displayCurrent))
		}
		fmt.Println()
	}
}

func rotate90deg(curDir Direction) Direction {
	if curDir == LEFT {
		return UP
	}
	if curDir == RIGHT {
		return DOWN
	}
	if curDir == UP {
		return RIGHT
	}
	if curDir == DOWN {
		return LEFT
	}
	return Direction{0, 0}
}

func inBounds(pos Direction, maxX, maxY int) bool {
	return pos[0] < maxX && pos[1] < maxY && pos[0] >= 0 && pos[1] >= 0
}

func determinePath(p, d Direction, cells *[][]Cell) {
	maxX, maxY := len((*cells)[0])-1, len(*cells)-1
	for inBounds(p, maxX, maxY) {
		nextStep := &(*cells)[p[1]+d[1]][p[0]+d[0]]
		if nextStep.blocked {
			d = rotate90deg(d)
		} else {
			(*cells)[p[1]][p[0]].current = false
			nextStep.visited = true
			nextStep.current = true
			p = Direction{p[0] + d[0], p[1] + d[1]}
		}
		//printCells(*cells, d, true)
	}
}

func countDistinct(c [][]Cell) int {
	sum := 0
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c); j++ {
			if c[i][j].visited {
				sum += 1
			}
		}
	}
	return sum
}

func part1(ls []string) int {
	cells, dir, pos := parseInput(ls)
	//printCells(cells, dir, true)
	determinePath(pos, dir, &cells)
	//printCells(cells, dir, false)
	return countDistinct(cells)
}

func inArray(dirs []Direction, dir Direction) bool {
	for _, d := range dirs {
		if (d[0] == dir[0]) && (d[1] == dir[1]) {
			return true
		}
	}
	return false
}

func determineOptions(p, d Direction, cells *[][]Cell) []Direction {
	var ds []Direction
	maxX, maxY := len((*cells)[0])-1, len(*cells)-1

	for inBounds(p, maxX, maxY) {
		nextStepXY := Direction{p[0] + d[0], p[1] + d[1]}
		//fmt.Printf("Step %v\n", nextStepXY)
		nextStep := &(*cells)[p[1]+d[1]][p[0]+d[0]]

		if nextStep.blocked {
			d = rotate90deg(d)
			//fmt.Printf("> Blocked, rotating\n")
		} else {
			if !nextStep.visited && checkLoop(p, d, *cells) {
				ds = append(ds, nextStepXY)
				fmt.Println("Found loop: ", len(ds))
			}
			(*cells)[p[1]][p[0]].current = false
			nextStep.visited = true
			nextStep.current = true
			p = Direction{p[0] + d[0], p[1] + d[1]}
		}
	}
	return ds
}

func checkLoop(p, d Direction, cells [][]Cell) bool {
	obst := Direction{p[0] + d[0], p[1] + d[1]}
	//fmt.Println("obs", obst)
	visited := make(map[Direction][]Direction)
	maxX, maxY := len(cells[0]), len(cells)
	for inBounds(p, maxX, maxY) {
		nextStepXY := Direction{p[0] + d[0], p[1] + d[1]}
		//fmt.Printf("\t[L] %v\n", nextStepXY)
		if !inBounds(nextStepXY, maxX, maxY) {
			//fmt.Printf("\t[OOB] %v\n", nextStepXY)
			return false
		}
		nextStep := &cells[nextStepXY[1]][nextStepXY[0]]

		if nextStep.blocked || ((nextStepXY[0] == obst[0]) && (nextStepXY[1] == obst[1])) {
			d = rotate90deg(d)
			//fmt.Printf("ROTATED %v\n", d)
			continue

		}
		if visited[nextStepXY] != nil {
			if inArray(visited[nextStepXY], d) {
				return true
			}
		} else {
			visited[nextStepXY] = []Direction{}
		}

		p = nextStepXY
		visited[p] = append(visited[p], d)
	}
	return false
}

func addSpecials(cells *[][]Cell, ds []Direction) [][]Cell {
	for _, d := range ds {
		(*cells)[d[1]][d[0]].special = true
	}
	return *cells
}

func part2(ls []string) int {
	cells, dir, pos := parseInput(ls)
	//printCells(cells, dir, true)
	ds := determineOptions(pos, dir, &cells)
	return len(ds)
}
