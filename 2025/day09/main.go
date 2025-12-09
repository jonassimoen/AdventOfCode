package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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

type Interval struct {
	start, end int
}

type Point struct {
	x, y int
}

type Grid struct {
	points     []Point
	pointSet   map[Point]int
	boundary   map[Point]bool
	minX, maxX int
	minY, maxY int
	intervals  map[int][]Interval
}

func parseInput(ls []string) *Grid {
	g := &Grid{
		points:    make([]Point, 0, len(ls)),
		pointSet:  make(map[Point]int, len(ls)),
		boundary:  make(map[Point]bool),
		intervals: make(map[int][]Interval),
	}
	for i, l := range ls {
		split := strings.Split(l, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		pt := Point{x, y}
		g.points = append(g.points, pt)
		g.pointSet[pt] = i

		if i == 0 || x < g.minX {
			g.minX = x
		}
		if i == 0 || x > g.maxX {
			g.maxX = x
		}
		if i == 0 || y < g.minY {
			g.minY = y
		}
		if i == 0 || y > g.maxY {
			g.maxY = y
		}
	}
	return g
}

func part1(ls []string) int {
	g := parseInput(ls)
	maxArea := 0
	for i := 0; i < len(g.points); i++ {
		for j := i + 1; j < len(g.points); j++ {
			dx := abs(g.points[i].x-g.points[j].x) + 1
			dy := abs(g.points[i].y-g.points[j].y) + 1
			if dx*dy > maxArea {
				maxArea = dx * dy
			}
		}
	}
	return maxArea
}

func (g *Grid) print() {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if _, exists := g.pointSet[Point{x, y}]; exists {
				fmt.Print("#")
			} else if g.boundary[Point{x, y}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) printWithPoints(xA, yA, xB, yB int) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if (x == xA && y == yA) || (x == xB && y == yB) {
				fmt.Print("A")
				continue
			}
			if _, exists := g.pointSet[Point{x, y}]; exists {
				fmt.Print("#")
			} else if g.boundary[Point{x, y}] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insideCorners(xMin, yMin, xMax, yMax int) [][]float64 {
	return [][]float64{
		{float64(xMin) + 0.25, float64(yMin) + 0.25},
		{float64(xMax) - 0.25, float64(yMin) + 0.25},
		{float64(xMin) + 0.25, float64(yMax) - 0.25},
		{float64(xMax) - 0.25, float64(yMax) - 0.25},
	}
}

func edges(xMin, yMin, xMax, yMax int) [][]Point {
	return [][]Point{
		{Point{xMin, yMin}, Point{xMax, yMin}},
		{Point{xMax, yMin}, Point{xMax, yMax}},
		{Point{xMax, yMax}, Point{xMin, yMax}},
		{Point{xMin, yMax}, Point{xMin, yMin}},
	}
}

func (g *Grid) isPointInsidePolygon(p []float64) bool {
	x, y := p[0], p[1]
	inside := false
	for i := 0; i < len(g.points); i++ {
		var prev, cur *Point
		if i == 0 {
			prev = &g.points[len(g.points)-1]
		} else {
			prev = &g.points[i-1]
		}
		cur = &g.points[i]
		if cur.y == prev.y {
			continue
		}

		prevY := float64(prev.y)
		curY := float64(cur.y)
		prevX := float64(prev.x)
		curX := float64(cur.x)

		if (prevY < y && y < curY) || (curY < y && y < prevY) {
			x_ints := prevX + (y-prevY)*(curX-prevX)/(curY-prevY)
			if x_ints > x {
				inside = !inside
			}
		}
	}
	return inside
}

func (g *Grid) doesLinesIntersect(a1, a2, b1, b2 *Point) bool {
	ax1, ay1 := a1.x, a1.y
	ax2, ay2 := a2.x, a2.y
	bx1, by1 := b1.x, b1.y
	bx2, by2 := b2.x, b2.y

	if by1 == by2 && ax1 == ax2 {
		ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 = bx1, by1, bx2, by2, ax1, ay1, ax2, ay2
	} else if ay1 != ay2 || bx1 != bx2 {
		return false
	}

	// Sort to get min and max
	xMin, xMax := ax1, ax2
	if ax1 > ax2 {
		xMin, xMax = ax2, ax1
	}
	yMin, yMax := by1, by2
	if by1 > by2 {
		yMin, yMax = by2, by1
	}

	return (xMin < bx1 && bx1 < xMax) && (yMin < ay1 && ay1 < yMax)
}

func (g *Grid) rectangleInsidePolygon(bl, tr *Point) bool {
	for _, p := range insideCorners(bl.x, bl.y, tr.x, tr.y) {
		if !g.isPointInsidePolygon(p) {
			return false
		}
	}

	for _, e := range edges(bl.x, bl.y, tr.x, tr.y) {
		e1, e2 := e[0], e[1]
		for i := 0; i < len(g.points); i++ {
			var prev, cur *Point
			if i == 0 {
				prev = &g.points[len(g.points)-1]
			} else {
				prev = &g.points[i-1]
			}
			cur = &g.points[i]
			if g.doesLinesIntersect(&e1, &e2, prev, cur) {
				return false
			}
		}
	}

	return true
}

func part2(ls []string) int {
	g := parseInput(ls)
	//g.buildBoundary()
	//g.buildIntervals()

	maxArea := 0

	for i := 0; i < len(g.points); i++ {
		for j := i + 1; j < len(g.points); j++ {
			xA, yA := g.points[i].x, g.points[i].y
			xB, yB := g.points[j].x, g.points[j].y

			dx := abs(xA-xB) + 1
			dy := abs(yA-yB) + 1
			area := dx * dy

			if area > maxArea {
				bole := &Point{min(xA, xB), min(yA, yB)}
				tori := &Point{max(xA, xB), max(yA, yB)}
				if g.rectangleInsidePolygon(bole, tori) {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
