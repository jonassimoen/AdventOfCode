package main

import (
	"bufio"
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

type Coordinate struct {
	x, y float64
}

type Package struct {
	coords []Coordinate
}

type ListPackages struct {
	packages []Package
}

func parseInput(ls []string) *ListPackages {
	packages := []Package{}
	for _, l := range ls {
		coordinatesPackage := []Coordinate{}
		coordinates := strings.Split(l, "), ")
		for _, coordinateStr := range coordinates {
			removeBrackets := strings.ReplaceAll(strings.ReplaceAll(coordinateStr, "(", ""), ")", "")
			coords := strings.Split(removeBrackets, ", ")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err)
			}
			coordinatesPackage = append(coordinatesPackage, Coordinate{float64(x), float64(y)})
		}
		packages = append(packages, Package{coordinatesPackage})
	}
	return &ListPackages{packages}
}

func calculateRadiusToPoint(c *Coordinate) float64 {
	a := math.Pow(c.x, 2)
	b := math.Pow(c.y, 2)
	return math.Sqrt(a + b)
}

func calculateBalloonRadius(p *Package) float64 {
	maxRadius := float64(0)
	for _, c := range p.coords {
		radius := calculateRadiusToPoint(&c)
		if radius > maxRadius {
			maxRadius = radius
		}
	}
	return maxRadius
}

func part1(l *ListPackages) {
	sumRadii := float64(0)
	for _, p := range l.packages {
		r := calculateBalloonRadius(&p)
		sumRadii += r
	}

	fmt.Printf("Part 1: %d\n", int(math.Floor(sumRadii)))
}

func calculateRadiusPointToPoint(c1 *Coordinate, c2 *Coordinate) float64 {
	a := math.Pow(c1.x-c2.x, 2)
	b := math.Pow(c1.y-c2.y, 2)
	return math.Sqrt(a + b)
}

func calculateCircleCenter(c1 *Coordinate, c2 *Coordinate, c3 *Coordinate) (*Coordinate, bool) {
	x1, y1 := c1.x, c1.y
	x2, y2 := c2.x, c2.y
	x3, y3 := c3.x, c3.y

	// D = 2(x₁(y₂-y₃) + x₂(y₃-y₁) + x₃(y₁-y₂))
	d := 2 * (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))

	if math.Abs(d) < 1e-10 {
		return nil, false
	}

	//cx = [(x₁²+y₁²)(y₂-y₃) + (x₂²+y₂²)(y₃-y₁) + (x₃²+y₃²)(y₁-y₂)] / D
	//cy = [(x₁²+y₁²)(x₃-x₂) + (x₂²+y₂²)(x₁-x₃) + (x₃²+y₃²)(x₂-x₁)] / D
	x1y1 := x1*x1 + y1*y1
	x2y2 := x2*x2 + y2*y2
	x3y3 := x3*x3 + y3*y3

	cx := (((x1y1) * (y2 - y3)) + (((x2y2) * (y3 - y1)) + ((x3y3) * (y1 - y2)))) / d
	cy := (((x1y1) * (x3 - x2)) + (((x2y2) * (x1 - x3)) + ((x3y3) * (x2 - x1)))) / d

	return &Coordinate{x: cx, y: cy}, true
}

func pointsWithinCircle(p *Package, center *Coordinate, radius float64) bool {
	for _, c := range p.coords {
		if calculateRadiusPointToPoint(center, &c) > radius {
			return false
		}
	}
	return true
}

func calculateBalloonRadiusCircumcircle(p *Package) float64 {
	minRadius := math.MaxFloat64

	for i := 0; i < len(p.coords); i++ {
		for j := i + 1; j < len(p.coords); j++ {
			c1, c2 := &p.coords[i], &p.coords[j]

			center := &Coordinate{x: (c1.x + c2.x) / 2, y: (c1.y + c2.y) / 2}
			radius := calculateRadiusPointToPoint(center, c1)

			if pointsWithinCircle(p, center, radius) && radius < minRadius {
				minRadius = radius
			}

			for k := j + 1; k < len(p.coords); k++ {
				ci, cj, ck := &p.coords[i], &p.coords[j], &p.coords[k]
				center, ok := calculateCircleCenter(ci, cj, ck)
				if !ok {
					continue
				}
				radius := calculateRadiusPointToPoint(center, &p.coords[i])
				if pointsWithinCircle(p, center, radius) && radius < minRadius {
					minRadius = radius
				}
			}
		}
	}
	return minRadius
}

func part2(l *ListPackages) {
	sumRadii := float64(0)
	for _, p := range l.packages {
		r := calculateBalloonRadiusCircumcircle(&p)
		sumRadii += r
	}

	fmt.Printf("Part 2: %d\n", int(math.Floor(sumRadii)))
}

func solve(inputFile string) {
	fmt.Println(inputFile)
	ls, err := readFile(inputFile)
	if err != nil {
		panic(err)
	}

	l := parseInput(ls)

	part1(l)
	part2(l)
}

func main() {
	solve("./in_test")
	solve("./in_test2")
	solve("./in")
}
