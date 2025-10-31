package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

type Coordinate [2]int

var TOP = Coordinate{0, -1}
var BOTTOM = Coordinate{0, 1}
var RIGHT = Coordinate{1, 0}
var LEFT = Coordinate{-1, 0}

func (c Coordinate) add(other Coordinate) Coordinate {
	return Coordinate{c[0] + other[0], c[1] + other[1]}
}

func (c Coordinate) min(other Coordinate) Coordinate {
	return Coordinate{c[0] - other[0], c[1] - other[1]}
}

func (c Coordinate) Stringify() string {
	if c == TOP {
		return "TOP"
	}
	if c == BOTTOM {
		return "BOTTOM"
	}
	if c == LEFT {
		return "LEFT"
	}
	if c == RIGHT {
		return "RIGHT"
	}
	return ""
}

func (c Coordinate) Charify() string {
	if c == TOP {
		return "^"
	}
	if c == BOTTOM {
		return "v"
	}
	if c == LEFT {
		return "<"
	}
	if c == RIGHT {
		return ">"
	}
	return "."
}

type Cell struct {
	typ       int // 0 free - 1 wall
	minScore  int
	origins   []Coordinate
	partOfWay bool
}

type Grid struct {
	cells [][]Cell
	s     Coordinate
	e     Coordinate
	paths []Coordinate
}

type QueueItem struct {
	c     Coordinate
	d     Coordinate
	score int
}

func (q QueueItem) String() string {
	return fmt.Sprintf("(%d,%d) [%s] (%d)", q.c[0], q.c[1], q.d.Stringify(), q.score)
}

func parseInput(ls []string) *Grid {
	cells := [][]Cell{}
	s := Coordinate{}
	e := Coordinate{}
	for j, l := range ls {
		cells = append(cells, []Cell{})
		for i, c := range l {
			typ := 0
			if c == '#' {
				typ = 1
			} else if c == 'S' {
				s = Coordinate{i, j}
			} else if c == 'E' {
				e = Coordinate{i, j}
			}
			cells[j] = append(cells[j], Cell{
				typ,
				math.MaxInt32,
				nil,
				false,
			})
		}
	}
	cells[s[1]][s[0]].minScore = 0
	return &Grid{
		cells: cells,
		s:     s,
		e:     e,
	}
}

func turn(cw bool, d Coordinate) Coordinate {
	if cw {
		switch d {
		case TOP:
			return RIGHT
		case BOTTOM:
			return LEFT
		case LEFT:
			return TOP
		case RIGHT:
			return BOTTOM
		}
	} else {
		switch d {
		case TOP:
			return LEFT
		case BOTTOM:
			return RIGHT
		case LEFT:
			return BOTTOM
		case RIGHT:
			return TOP
		}
	}
	return Coordinate{}
}

func sortQueue(q []QueueItem) {
	sort.Slice(q, func(i, j int) bool {
		return q[i].score < q[j].score
	})
}

func (g *Grid) isValidCell(p Coordinate) bool {
	if (p[0] < 0 || p[1] < 0 || p[1] >= len(g.cells[0]) || p[0] >= len(g.cells)) ||
		g.cells[p[1]][p[0]].typ == 1 {
		return false
	}
	return true
}

func (g *Grid) isInOrigins(p Coordinate, origin Coordinate) bool {
	for _, c := range g.cells[p[1]][p[0]].origins {
		if c == origin {
			return true
		}
	}
	return false
}

func (g *Grid) findWay(log bool) int {
	q := make([]QueueItem, 0)
	//path := make([]Coordinate, 0)

	best_path_score := math.MaxInt32

	q = append(q, QueueItem{
		c:     g.s,
		d:     RIGHT,
		score: 0,
	})

	for (len(q) > 0) && (q[0].score <= best_path_score) {
		item := q[0]
		q = q[1:]

		dir := item.d
		coord := item.c

		if log {
			fmt.Printf("From q: %v %s %d: ", coord, dir.Stringify(), item.score)
		}

		if (coord[0] == g.e[0]) && (coord[1] == g.e[1]) {
			best_path_score = item.score
			if log {
				fmt.Printf("@ end with score=%d\n", best_path_score)
			}
			continue
		}

		if g.cells[coord[1]][coord[0]].minScore < item.score {
			if log {
				fmt.Printf("score %d is higher than minScore of coordinate (%d)\n", item.score, g.cells[coord[1]][coord[0]].minScore)
			}
			continue
		}

		if !g.isValidCell(coord) {
			if log {
				fmt.Printf("%v is invalid\n", coord)
			}
			continue
		}

		g.cells[coord[1]][coord[0]].minScore = item.score
		for _, newDir := range []Coordinate{turn(false, dir), dir, turn(true, dir)} {
			newCoord := coord.add(newDir)
			score := item.score + 1
			if newDir != dir {
				score += 1000
			}
			if log {
				fmt.Printf("\n\t[%s] = %d", newCoord.Stringify(), score)
			}
			if !g.isInOrigins(newCoord, coord) {
				g.cells[newCoord[1]][newCoord[0]].origins = append(g.cells[newCoord[1]][newCoord[0]].origins, coord)
				if log {
					fmt.Printf(" ==> added %v to origins", coord)
				}
			}
			if log {
				fmt.Printf(" = %v", g.cells[newCoord[1]][newCoord[0]].origins)
			}

			q = append(q, QueueItem{
				c:     newCoord,
				d:     newDir,
				score: score,
			})
		}

		if log {
			fmt.Println()
		}
		sortQueue(q)

	}
	return best_path_score
}

const DEBUG_LOG = false

func part1(ls []string) int {
	g := parseInput(ls)
	s := g.findWay(DEBUG_LOG)
	return s
}

func (g *Grid) recreatePath(c Coordinate) {
	if g.cells[c[1]][c[0]].partOfWay {
		return
	}
	g.cells[c[1]][c[0]].partOfWay = true
	for _, o := range g.cells[c[1]][c[0]].origins {
		g.recreatePath(o)
	}
}

func (g *Grid) print() {
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.typ == 1 {
				fmt.Print("#")
			} else if cell.partOfWay {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// TODO
func part2(ls []string) int {
	g := parseInput(ls)
	g.findWay(DEBUG_LOG)
	//g.print()
	g.recreatePath(g.e)
	//g.print()
	sum := 0
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.partOfWay && cell.typ == 0 {
				sum += 1
			}
		}
	}
	// SHOULD BE 502, returns 526
	return sum
}
