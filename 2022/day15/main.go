package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
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

type Coordinate struct {
	sensorX, sensorY int
	beaconX, beaconY int
	manhattanDist    int
}

type Bounds struct {
	minX, minY, maxX, maxY int
}

func parseCoordinates(ls []string) ([]Coordinate, int, Bounds) {
	var coordinates []Coordinate
	maxExtraEdgeDistance := -1
	bounds := Bounds{999, 999, 0, 0}
	for _, l := range ls {
		r := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
		matches := r.FindStringSubmatch(l)[1:]

		Sx, _ := strconv.Atoi(matches[0])
		Sy, _ := strconv.Atoi(matches[1])
		Bx, _ := strconv.Atoi(matches[2])
		By, _ := strconv.Atoi(matches[3])
		coord := Coordinate{
			sensorX:       Sx,
			sensorY:       Sy,
			beaconX:       Bx,
			beaconY:       By,
			manhattanDist: manhattanDistance(Sx, Sy, Bx, By),
		}
		coordinates = append(coordinates, coord)
		maxExtraEdgeDistance = max(maxExtraEdgeDistance, coord.manhattanDist)
		bounds.minX = min(bounds.minX, min(Sx, Bx))
		bounds.maxX = max(bounds.maxX, max(Sx, Bx))
		bounds.minY = min(bounds.minY, min(Sy, By))
		bounds.maxY = max(bounds.maxY, max(Sy, By))
	}
	return coordinates, maxExtraEdgeDistance, bounds
}

func manhattanDistance(xA, yA, xB, yB int) int {
	return int(math.Abs(float64(yA-yB))) + int(math.Abs(float64(xA-xB)))
}

func processLine(y int, coordinates []Coordinate, bounds Bounds, maxDistance int) int {
	totalCloser := 0
	for x := bounds.minX - maxDistance; x < bounds.maxX+maxDistance; x++ {
		if isBeaconCloserAtLocation(coordinates, x, y) == 1 {
			totalCloser++
		}
	}
	return totalCloser
}

// Returns 0 if already beacon, 1 if it would be a closer beacon, 2 if it is a valid place
func isBeaconCloserAtLocation(coordinates []Coordinate, x, y int) int {
	for _, coord := range coordinates {
		if coord.beaconX == x && coord.beaconY == y {
			return 0
		} else if coord.manhattanDist >= manhattanDistance(x, y, coord.sensorX, coord.sensorY) {
			return 1
		}
	}
	return 2
}

func checkBorders(coordinates []Coordinate) {
	all := [][]int{}
	for _, coord := range coordinates {
		x, y := coord.sensorX, coord.sensorY
		lX := max(PART_2_LOWER_BOUND, x-coord.manhattanDist-1)
		hX := min(PART_2_UPPER_BOUND, x+coord.manhattanDist+1)
		lY := max(PART_2_LOWER_BOUND, y-coord.manhattanDist-1)
		hY := min(PART_2_UPPER_BOUND, y+coord.manhattanDist+1)
		pts := getPointsOutsideBorder(lX, hX, lY, hY, coord.manhattanDist+1, x, y)
		fmt.Printf("%v ==> %v", []int{x, y}, pts)
		//checkPoints(pts, coordinates)
	}
	fmt.Println(len(all))
}

func checkPoints(pts [][]int, coordinates []Coordinate) {
	for _, pt := range pts {
		ptWithinRange := false
		for _, coord := range coordinates {
			if coord.manhattanDist <= manhattanDistance(coord.sensorX, coord.sensorY, pt[0], pt[1]) {
				ptWithinRange = true
			}
		}
		if !ptWithinRange {
			fmt.Printf("%v is outside ranges of all sensors", pt)
		}
	}
}

func getPointsOutsideBorder(lX, hX, lY, hY int, dist int, sx, sy int) [][]int {
	pts := [][]int{}
	for x := lX; x <= hX; x++ {
		for y := lY; y <= hY; y++ {
			if manhattanDistance(x, y, sx, sy) == dist {
				pts = append(pts, []int{x, y})
			}
		}
	}
	return pts
}

// const PART_1_LINE = 10
// const PART_2_LOWER_BOUND = 0
// const PART_2_UPPER_BOUND = 20
const PART_1_LINE = 2000000
const PART_2_LOWER_BOUND = 0
const PART_2_UPPER_BOUND = 4000000

func part1(ls []string) int {
	coordinates, maxDistance, bounds := parseCoordinates(ls)
	x := processLine(PART_1_LINE, coordinates, bounds, maxDistance)
	return x
}

func part2(ls []string) int {
	coordinates, _, _ := parseCoordinates(ls)
	freq := -1
	checkBorders(coordinates)
	return freq
}
