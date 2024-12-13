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

type Plot struct {
	Coordinate
	EdgeExplainer
	inArea bool
	value  rune
}

type Garden struct {
	m [][]Plot
}

func (p Plot) String() string {
	return fmt.Sprintf("[%c] (%d,%d) %t", p.value, p.x+1, p.y+1, p.inArea)
}

func parseInput(ls []string) Garden {
	var m [][]Plot
	for i, l := range ls {
		m = append(m, []Plot{})
		for j, r := range l {
			m[i] = append(m[i], Plot{Coordinate{j, i}, EdgeExplainer{false, false, false, false}, false, r})
		}
	}
	return Garden{
		m: m,
	}
}

func (g *Garden) findArea(p Plot) []Plot {
	var s []Plot
	//fmt.Printf("\t%v:\n", p)
	if p.x < len(g.m[p.y])-1 {
		right := &g.m[p.y][p.x+1]
		if (right.value == p.value) && !right.inArea {
			//fmt.Printf("%v RIGHT", p)
			g.m[p.y][p.x+1].inArea = true
			s = append(s, g.findArea(*right)...)
		}
	}
	if p.y < len(g.m)-1 {
		bottom := &g.m[p.y+1][p.x]
		if bottom.value == p.value && !bottom.inArea {
			//fmt.Printf("%v BOTTOM", p)
			g.m[p.y+1][p.x].inArea = true
			s = append(s, g.findArea(*bottom)...)
		}
	}
	if p.y > 0 {
		left := &g.m[p.y-1][p.x]
		if left.value == p.value && !left.inArea {
			//fmt.Printf("%v LEFT", p)
			g.m[p.y-1][p.x].inArea = true
			s = append(s, g.findArea(*left)...)
		}
	}
	if p.x > 0 {
		top := &g.m[p.y][p.x-1]
		if top.value == p.value && !top.inArea {
			//fmt.Printf("%v LEFT", p)
			g.m[p.y][p.x-1].inArea = true
			s = append(s, g.findArea(*top)...)
		}
	}
	g.m[p.y][p.x].inArea = true
	s = append(s, p)
	return s
}

func isInArray(element Plot, array []Plot) bool {
	for _, e := range array {
		if (e.x == element.x) && (e.y == element.y) {
			return true
		}
	}
	return false
}

func isInArrayCoords(element Coordinate, array []Coordinate) bool {
	for _, e := range array {
		if (e.x == element.x) && (e.y == element.y) {
			return true
		}
	}
	return false
}

func (g *Garden) calculateNeighbors(p Plot) int {
	n := 0
	if (p.x > 0) && (g.m[p.y][p.x-1].value == p.value) {
		n++
	}
	if (p.x < len(g.m[p.y])-1) && (g.m[p.y][p.x+1].value == p.value) {
		n++
	}
	if (p.y > 0) && (g.m[p.y-1][p.x].value == p.value) {
		n++
	}
	if (p.y < len(g.m)-1) && (g.m[p.y+1][p.x].value == p.value) {
		n++
	}
	return n
}

func part1(ls []string) int {
	g := parseInput(ls)
	var s [][]Plot
	for _, pr := range g.m {
		for _, p := range pr {
			if !p.inArea {
				//fmt.Printf("Start @ %v\n", p)
				s = append(s, g.findArea(p))
			}
		}
	}
	sum := 0
	for _, si := range s {
		var unique []Plot
		for _, p := range si {
			if !isInArray(p, unique) {
				unique = append(unique, p)
			}
		}
		area := len(unique)
		perim := 0
		for _, uq := range unique {
			perim += (4 - g.calculateNeighbors(uq))
		}
		sum += perim * area
		//fmt.Println(si[0].value, "==>", area, perim, area*perim)
	}
	//fmt.Println(g)
	return sum
}

type Coordinate struct {
	x, y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

// Map regions by starting coordinate
func (g *Garden) parseToMap() map[Plot][]Coordinate {
	dict := make(map[Plot][]Coordinate)
	for y, r := range g.m {
		for x, c := range r {
			fmt.Println(x, y, c)
			//coord := Coordinate{x, y}
			//if dict[coord] == nil {
			//	dict[coord] = []Coordinate{}
			//}
			//dict[coord] = append(dict[coord], coord)
		}
	}
	return dict
}

type EdgeExplainer struct {
	t, b, l, r bool
}

func (g *Garden) addEdges() {
	for y, r := range g.m {
		for x, c := range r {
			if ((y > 0) && (g.m[y-1][x].value != c.value)) || y == 0 {
				g.m[y][x].t = true
			}
			if ((y < len(g.m)-1) && (g.m[y+1][x].value != c.value)) || y == len(g.m)-1 {
				g.m[y][x].b = true
			}
			if ((x > 0) && (g.m[y][x-1].value != c.value)) || x == 0 {
				g.m[y][x].l = true
			}
			if ((x < len(g.m[y])-1) && (g.m[y][x+1].value != c.value)) || x == len(g.m[y])-1 {
				g.m[y][x].r = true
			}
		}
	}
}

func (g *Garden) countEdgesForRegion(regionCoords []Coordinate) int {
	edges := 0
	for _, r := range g.m {
		for _, c := range r {
			if !isInArrayCoords(c.Coordinate, regionCoords) {
				continue
			}
			leftHasTop := (c.x > 0) && !g.m[c.y][c.x-1].t
			leftHasBottom := (c.x > 0) && !g.m[c.y][c.x-1].b
			topHasLeft := (c.y > 0) && !g.m[c.y-1][c.x].l
			topHasRight := (c.y > 0) && !g.m[c.y-1][c.x].r
			//fmt.Printf("%v: ", c.Coordinate)

			if c.t && (leftHasTop || c.l) {
				//fmt.Printf("TOP, ")
				edges++
			}
			if c.b && (leftHasBottom || c.l) {
				//fmt.Printf("BOTTOM, ")
				edges++
			}
			if c.l && (topHasLeft || c.t) {
				//fmt.Printf("LEFT, ")
				edges++
			}
			if c.r && (topHasRight || c.t) {
				//fmt.Printf("RIGHT, ")
				edges++
			}
		}
	}
	return edges
}

func part2(ls []string) int {
	g := parseInput(ls)
	g.addEdges()
	var s [][]Plot
	for _, pr := range g.m {
		for _, p := range pr {
			if !p.inArea {
				s = append(s, g.findArea(p))
			}
		}
	}
	sum := 0
	for _, si := range s {
		var unique []Coordinate
		for _, p := range si {
			if !isInArrayCoords(p.Coordinate, unique) {
				unique = append(unique, p.Coordinate)
			}
		}
		area := len(unique)
		edges := g.countEdgesForRegion(unique)
		sum += edges * area
		//fmt.Println(si[0].value, "==>", area, perim, area*perim)
	}
	return sum
}
