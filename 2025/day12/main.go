package main

import (
	"bufio"
	"flag"
	"fmt"
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

type Simplified struct {
	sizes      []int
	rectangles []Rectangle
}

type Rectangle struct {
	w, h     int
	presents []int
}

func parseSimplified(ls []string) *Simplified {
	sizes := make([]int, 6)

	for l := 0; l < 30; l += 5 {
		idx := ls[l][0] - '0'
		size := 0
		for _, r := range ls[l+1 : l+5] {
			for _, c := range r {
				if c == '#' {
					size++
				}
			}
		}
		sizes[idx] = size
	}

	var rectangles []Rectangle
	for _, l := range ls[30:] {
		split := strings.Split(l, ": ")
		size := strings.Split(split[0], "x")
		w, _ := strconv.Atoi(size[0])
		h, _ := strconv.Atoi(size[1])

		var presents []int
		presentSplits := strings.Split(split[1], " ")
		for _, p := range presentSplits {
			present, _ := strconv.Atoi(p)
			presents = append(presents, present)
		}
		rectangles = append(rectangles, Rectangle{w, h, presents})
	}

	return &Simplified{sizes, rectangles}
}

type Shape struct {
	cells        [][]bool
	size         int
	orientations [][][]bool // Pre-compute all unique orientations
}

type Program struct {
	shapes     []Shape
	rectangles []Rectangle
}

func parseShapes(ls []string) []Shape {
	var shapes []Shape
	for i := 0; i < len(ls); i += 5 {
		var cells [][]bool
		size := 0
		for j := 1; j <= 3; j++ {
			l := ls[i+j]
			row := make([]bool, len(l))
			for k, c := range l {
				if c == '#' {
					row[k] = true
					size++
				}
			}
			cells = append(cells, row)
		}
		orientations := computeOrientations(cells)
		shapes = append(shapes, Shape{cells, size, orientations})
	}
	return shapes
}

func computeOrientations(cells [][]bool) [][][]bool {
	seen := make(map[string]bool)
	var unique [][][]bool

	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			transformed := transform(cells, rot, flip == 1)
			key := serializeShape(transformed)
			if !seen[key] {
				seen[key] = true
				unique = append(unique, transformed)
			}
		}
	}
	return unique
}

func serializeShape(cells [][]bool) string {
	var sb strings.Builder
	for _, row := range cells {
		for _, cell := range row {
			if cell {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		sb.WriteByte('|')
	}
	return sb.String()
}

func parseInput(ls []string) *Program {
	shapes := parseShapes(ls[:30])
	var rectangles []Rectangle
	for _, l := range ls[30:] {
		split := strings.Split(l, ": ")
		size := strings.Split(split[0], "x")
		w, _ := strconv.Atoi(size[0])
		h, _ := strconv.Atoi(size[1])

		var presents []int
		presentSplits := strings.Split(split[1], " ")
		for _, p := range presentSplits {
			present, _ := strconv.Atoi(p)
			presents = append(presents, present)
		}
		rectangles = append(rectangles, Rectangle{w, h, presents})
	}
	return &Program{shapes, rectangles}
}

func (p *Program) canFit(w, h int, required []int) bool {
	// Early exit: check total area
	totalArea := 0
	for shapeIdx, count := range required {
		totalArea += p.shapes[shapeIdx].size * count
	}
	if totalArea > w*h {
		return false
	}

	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}

	// Build list of shapes to place, sorted by size (largest first)
	type ShapeToPlace struct {
		idx  int
		size int
	}
	var toPlace []ShapeToPlace
	for shapeIdx, count := range required {
		for i := 0; i < count; i++ {
			toPlace = append(toPlace, ShapeToPlace{shapeIdx, p.shapes[shapeIdx].size})
		}
	}

	// Sort by size descending - place larger shapes first
	sort.Slice(toPlace, func(i, j int) bool {
		return toPlace[i].size > toPlace[j].size
	})

	shapeIndices := make([]int, len(toPlace))
	for i, s := range toPlace {
		shapeIndices[i] = s.idx
	}

	return p.backtrack(grid, shapeIndices, 0)
}

func (p *Program) backtrack(grid [][]bool, toPlace []int, idx int) bool {
	if idx == len(toPlace) {
		return true
	}

	shapeIdx := toPlace[idx]
	shape := p.shapes[shapeIdx]

	// Try each pre-computed orientation
	for _, transformed := range shape.orientations {
		h, w := len(transformed), len(transformed[0])

		// Try all positions with early termination
		for row := 0; row <= len(grid)-h; row++ {
			for col := 0; col <= len(grid[0])-w; col++ {
				if canPlaceFast(grid, transformed, row, col) {
					place(grid, transformed, row, col, true)

					if p.backtrack(grid, toPlace, idx+1) {
						return true
					}

					place(grid, transformed, row, col, false)
				}
			}
		}
	}

	return false
}

func canPlaceFast(grid [][]bool, shape [][]bool, startRow, startCol int) bool {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] && grid[startRow+i][startCol+j] {
				return false
			}
		}
	}
	return true
}

func transform(cells [][]bool, rotation int, flip bool) [][]bool {
	result := cells

	// Apply flip first
	if flip {
		result = flipHorizontal(result)
	}

	// Apply rotation
	for i := 0; i < rotation; i++ {
		result = rotate90(result)
	}

	return result
}

func flipHorizontal(cells [][]bool) [][]bool {
	result := make([][]bool, len(cells))
	for i := range cells {
		result[i] = make([]bool, len(cells[i]))
		for j := range cells[i] {
			result[i][len(cells[i])-1-j] = cells[i][j]
		}
	}
	return result
}

func rotate90(cells [][]bool) [][]bool {
	if len(cells) == 0 {
		return cells
	}

	rows := len(cells)
	cols := len(cells[0])
	result := make([][]bool, cols)

	for i := 0; i < cols; i++ {
		result[i] = make([]bool, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = cells[rows-1-j][i]
		}
	}

	return result
}

func canPlace(grid [][]bool, shape [][]bool, startRow, startCol int) bool {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] {
				r, c := startRow+i, startCol+j
				if r >= len(grid) || c >= len(grid[0]) || grid[r][c] {
					return false
				}
			}
		}
	}
	return true
}

func place(grid [][]bool, shape [][]bool, startRow, startCol int, set bool) {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] {
				grid[startRow+i][startCol+j] = set
			}
		}
	}
}

func part1(ls []string) int {
	s := parseSimplified(ls)
	ok := 0
	for _, rect := range s.rectangles {
		maxSize := rect.w * rect.h
		presentsAllSize := 0
		for i, p := range rect.presents {
			presentsAllSize += p * s.sizes[i]
		}
		if presentsAllSize <= maxSize {
			ok++
		}
	}

	fmt.Printf("Simplified solution: %d\n", ok)
	ok = 0

	sol := parseInput(ls)
	for _, rect := range sol.rectangles {
		if sol.canFit(rect.w, rect.h, rect.presents) {
			ok++
		}
	}
	fmt.Printf("Final solution: %d\n", ok)
	return ok
}

func part2(ls []string) int {
	return -1
}
