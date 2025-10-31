package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

const MAX_SIZE = 70
const MAX_STEP = 12

type Cell struct {
	x, y     int
	valid    bool
	minScore int
}

type Grid struct {
	cells [][]Cell

	invalidCells []Cell
}

func parseInput(ls []string, part1 bool) *Grid {
	g := &Grid{make([][]Cell, MAX_SIZE+1), []Cell{}}
	for y := 0; y <= MAX_SIZE; y++ {
		g.cells[y] = make([]Cell, MAX_SIZE+1)
		for x := 0; x <= MAX_SIZE; x++ {
			g.cells[y][x] = Cell{
				x:        x,
				y:        y,
				valid:    true,
				minScore: math.MaxInt32,
			}
		}
	}
	steps := 1
	for _, l := range ls {
		coord := strings.Split(l, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		g.invalidCells = append(g.invalidCells, g.cells[y][x])
		if part1 {
			g.cells[y][x].valid = false
		}
		if part1 && (steps == MAX_STEP) {
			break
		}
		steps++
	}
	return g
}

type QueueItem struct {
	x, y  int
	score int
}

func sortQueue(q []QueueItem) {
	sort.Slice(q, func(i, j int) bool {
		return q[i].score < q[j].score
	})
}

func (g *Grid) drawInvalidate(step int) {
	for _, c := range g.invalidCells[:step] {
		g.cells[c.y][c.x].valid = false
	}
}

func (g *Grid) removeInvalidate() {
	for y := 0; y <= MAX_SIZE; y++ {
		for x := 0; x <= MAX_SIZE; x++ {
			g.cells[y][x].valid = true
		}
	}
}

func (g *Grid) print(step int) {
	g.removeInvalidate()
	g.drawInvalidate(step)
	for y := 0; y <= MAX_SIZE; y++ {
		for x := 0; x <= MAX_SIZE; x++ {
			if g.cells[y][x].valid {
				if g.cells[y][x].minScore < math.MaxInt32 {
					fmt.Printf("%2d", g.cells[y][x].minScore)
				} else {
					fmt.Printf(" .")
				}
			} else {
				fmt.Printf(" #")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) hasBeenInvalidated(x, y int, step int) (bool, int) {
	for i, c := range g.invalidCells[:step] {
		if c.x == x && c.y == y {
			return true, i
		}
	}
	return false, 0
}

func part1(ls []string) int {
	g := parseInput(ls, true)
	fmt.Println(len(g.cells))
	g.print(12)

	q := make([]QueueItem, 0)
	q = append(q, QueueItem{x: 0, y: 0, score: 0})

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		//fmt.Println("\033[H\033[2J")
		//g.print(i.score)

		//fmt.Printf("[%d, %d] = %d", i.x, i.y, i.score)

		//if inv, invStep := g.hasBeenInvalidated(i.x, i.y, i.score); inv {
		//	fmt.Printf("= invalidated at step %d (currently at step %d)\n", invStep, i.score)
		//	continue
		//}

		if !g.cells[i.y][i.x].valid {
			//	fmt.Printf("= invalidated")
			continue
		}

		if g.cells[i.y][i.x].minScore <= i.score {
			//fmt.Printf("= other shorter path found (%d steps, currently %d steps)\n", g.cells[i.y][i.x].minScore, i.score)
			continue
		}

		//fmt.Printf("= shortest path found = %d\n", i.score)
		g.cells[i.y][i.x].minScore = i.score

		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			x := i.x + d[0]
			y := i.y + d[1]
			if (x < 0 || x >= MAX_SIZE+1) || (y < 0 || y >= MAX_SIZE+1) || !g.cells[y][x].valid {
				continue
			}
			//fmt.Printf("\tadded [%d, %d], step %d\n", x, y, i.score+1)
			q = append(q, QueueItem{x, y, i.score + 1})
		}
		sortQueue(q)
	}
	//g.print(0)
	return g.cells[MAX_SIZE][MAX_SIZE].minScore
}

func inCluster(c Cell, cl *[]Cell) bool {
	for _, ci := range *cl {
		if ci.x == c.x && ci.y == c.y {
			return true
		}
	}
	return false
}

func (g *Grid) isOneCluster(step int) bool {
	cluster := make([]Cell, 0)
	q := make([]Cell, 0)
	q = append(q, g.cells[0][0])
	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		if inCluster(i, &cluster) {
			continue
		}

		cluster = append(cluster, i)

		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			x := i.x + d[0]
			y := i.y + d[1]
			if (x < 0 || x >= MAX_SIZE+1) || (y < 0 || y >= MAX_SIZE+1) || !g.cells[y][x].valid {
				continue
			}
			q = append(q, g.cells[y][x])
		}

	}
	return inCluster(g.cells[MAX_SIZE][MAX_SIZE], &cluster)
}

func (g *Grid) printCluster(c []Cell) {
	for y := 0; y <= MAX_SIZE; y++ {
		for x := 0; x <= MAX_SIZE; x++ {
			if g.cells[y][x].valid {
				if inCluster(g.cells[y][x], &c) {
					fmt.Printf("O")
				} else {
					fmt.Printf(".")
				}
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func part2(ls []string) int {
	g := parseInput(ls, false)
	for step := 0; step < len(g.invalidCells); step++ {
		g.drawInvalidate(step + 1)
		if !g.isOneCluster(step) {
			fmt.Printf("%d,%d\n", g.invalidCells[step].x, g.invalidCells[step].y)
			return 0
		}
	}
	return 0
}
