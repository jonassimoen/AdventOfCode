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

func main() {
	part1ptr := flag.Int("p", 1, "Part")
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Part %d - Input file: %s\n", *part1ptr, *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	var out int
	if *part1ptr == 1 {
		out = part1(ls)
	} else {
		out = part2(ls)
	}
	fmt.Printf("Output: %d\n", out)
}

type BaseCoordinate struct {
	x, y, z int
}

type Coordinate struct {
	BaseCoordinate
	sides int
}
type CoordinatePart2 struct {
	BaseCoordinate
	sides int
}

func (c *Coordinate) isAdjacent(cc *Coordinate) bool {
	return int(math.Abs(float64(c.x-cc.x))+math.Abs(float64(c.y-cc.y))+math.Abs(float64(c.z-cc.z))) == 1
}

func parseInput3D(ls []string) []Coordinate {
	s := make([]Coordinate, len(ls))
	for i, l := range ls {
		split := strings.Split(l, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		c := Coordinate{BaseCoordinate{x, y, z}, 6}
		s[i] = c
	}
	return s
}

func calculateSides(s []Coordinate) []Coordinate {
	for i := 0; i < len(s); i++ {
		ci := s[i]
		for j := i + 1; j < len(s); j++ {
			cj := s[j]
			if cj.isAdjacent(&ci) {
				s[i].sides--
				s[j].sides--
			}
		}
	}
	return s
}

func calculateSumSides(s []Coordinate) int {
	sum := 0
	for _, c := range s {
		sum += c.sides
	}
	return sum
}

func part1(ls []string) int {
	coordinateSlice := parseInput3D(ls)
	fmt.Printf("%v\n", coordinateSlice)
	coordinateSliceBis := calculateSides(coordinateSlice)
	fmt.Printf("%v\n", coordinateSlice)
	fmt.Printf("%v\n", coordinateSliceBis)
	return calculateSumSides(coordinateSliceBis)
}

func getMinMax(c []Coordinate) ([]int, []int, []int) {
	lx, ly, lz := 999999, 999999, 999999
	hx, hy, hz := -999999, -999999, -999999
	for _, ci := range c {
		lx = min(lx, ci.x)
		ly = min(ly, ci.y)
		lz = min(lz, ci.z)
		hx = max(hx, ci.x)
		hy = max(hy, ci.y)
		hz = max(hz, ci.z)
	}
	return []int{lx, hx}, []int{ly, hy}, []int{lz, hz}
}

func parseInput3Dint(c []Coordinate, x, y, z []int) [][][]int {
	s := make([][][]int, z[1]-z[0]+1)
	for zi := 0; zi <= z[1]-z[0]; zi++ {
		s[zi] = make([][]int, y[1]-y[0]+1)
		for yi := 0; yi <= y[1]-y[0]; yi++ {
			s[zi][yi] = make([]int, x[1]-x[0]+1)
			//for xi := 0; xi <= x[1]-x[0]; xi++ {
			//	s[zi][yi][xi] = 0
			//}
		}
	}
	for _, ci := range c {
		s[ci.z-z[0]][ci.y-y[0]][ci.x-x[0]] = 1
	}
	return s
}

func getNeighbors(c Coordinate, rX, rY, rZ []int) []BaseCoordinate {
	nb := make([]BaseCoordinate, 0)
	if c.x > rX[0] {
		nb = append(nb, BaseCoordinate{c.x - 1, c.y, c.z})
	}
	if c.y > rY[0] {
		nb = append(nb, BaseCoordinate{c.x, c.y - 1, c.z})
	}
	if c.z > rZ[0] {
		nb = append(nb, BaseCoordinate{c.x, c.y, c.z - 1})
	}
	if c.x < rX[1] {
		nb = append(nb, BaseCoordinate{c.x + 1, c.y, c.z})
	}
	if c.y < rY[1] {
		nb = append(nb, BaseCoordinate{c.x, c.y + 1, c.z})
	}
	if c.z < rZ[1] {
		nb = append(nb, BaseCoordinate{c.x, c.y, c.z + 1})
	}
	return nb
}

func calculateClosedIn(cs []Coordinate, ranges [][]int) int {
	queue := make([]BaseCoordinate, 0)
	queue = append(queue, cs[0].BaseCoordinate)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Println(top.x, top.y, top.z)
	}
	return 0
}

func part2(ls []string) int {
	coordinateSlice := parseInput3D(ls)
	coordinateSliceSides := calculateSides(coordinateSlice)
	x, y, z := getMinMax(coordinateSliceSides)
	//fmt.Println(coordinateSlice)
	//fmt.Println(x, y, z)
	calculateClosedIn(coordinateSliceSides, [][]int{x, y, z})
	return -1
}
