package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

type Coordinate struct {
	x, y int
}

type Grid struct {
	counts map[Coordinate]int
	xL, xH int
	yL, yH int
}

func (g *Grid) processCoordinate(c Coordinate) {
	if c.x < g.xL {
		g.xL = c.x
	}
	if c.x > g.xH {
		g.xH = c.x
	}
	if c.y < g.yL {
		g.yL = c.y
	}
	if c.y > g.yH {
		g.yH = c.y
	}
}

func (g *Grid) print() {
	for y := 0; y <= g.yH; y++ {
		for x := 0; x <= g.xH; x++ {
			if c, ok := g.counts[Coordinate{x, y}]; ok {
				fmt.Print(c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func stringToCoordinate(s string) *Coordinate {
	split := strings.Split(s, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return &Coordinate{x, y}
}

func lineToCoordinates(a, b *Coordinate, allowDiagonal bool) []Coordinate {
	var c []Coordinate

	dx := b.x - a.x
	if dx != 0 {
		dx /= int(math.Abs(float64(b.x - a.x)))
	}
	dy := b.y - a.y
	if dy != 0 {
		dy /= int(math.Abs(float64(b.y - a.y)))
	}
	if dx == 0 {
		for y := int(math.Min(float64(b.y), float64(a.y))); y <= int(math.Max(float64(b.y), float64(a.y))); y++ {
			c = append(c, Coordinate{a.x, y})
		}
	} else if dy == 0 {
		for x := int(math.Min(float64(b.x), float64(a.x))); x <= int(math.Max(float64(b.x), float64(a.x))); x++ {
			c = append(c, Coordinate{x, a.y})
		}
	} else if math.Abs(float64(dx)) != math.Abs(float64(dy)) {
		panic(fmt.Sprintf("dx != dy: %d != %d", dx, dy))
	} else if allowDiagonal {
		for i := 0; i <= int(math.Abs(float64(b.x-a.x))); i++ {
			c = append(c, Coordinate{a.x + i*dx, a.y + i*dy})
		}
	}
	return c
}

func parseInput(ls []string, allowDiagonal bool) *Grid {
	g := &Grid{
		xL: math.MaxInt, yL: math.MaxInt, xH: math.MinInt, yH: math.MinInt,
		counts: make(map[Coordinate]int),
	}
	for _, l := range ls {
		coords := strings.Split(l, " -> ")
		a, b := coords[0], coords[1]
		aCoord := stringToCoordinate(a)
		bCoord := stringToCoordinate(b)
		abCoords := lineToCoordinates(aCoord, bCoord, allowDiagonal)
		for _, c := range abCoords {
			g.processCoordinate(c)
			g.counts[c]++
		}
	}
	return g
}

func (g *Grid) countOverlaps() int {
	c := 0
	for _, v := range g.counts {
		if v > 1 {
			c++
		}
	}
	return c
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

func part1(ls []string) int {
	return parseInput(ls, false).countOverlaps()
}

func part2(ls []string) int {
	return parseInput(ls, true).countOverlaps()
}
