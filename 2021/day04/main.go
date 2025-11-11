package main

import (
	"bufio"
	"flag"
	"fmt"
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

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

type Board [][]int

type Game struct {
	boards          []Board
	numbers         []int
	currentDraw     int
	completedBoards map[int]bool
}

func parseInput(ls []string) Game {
	g := Game{
		numbers:         []int{},
		boards:          []Board{},
		currentDraw:     -1,
		completedBoards: make(map[int]bool),
	}
	nLine := strings.Split(ls[0], ",")
	for _, n := range nLine {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		g.numbers = append(g.numbers, i)
	}

	board := Board{}
	for _, l := range ls[2:] {
		if len(l) == 0 {
			g.boards = append(g.boards, board)
			board = Board{}
			continue
		}
		var lineNumbers []int
		for _, n := range strings.Split(l, " ") {
			if n == "" {
				continue
			}
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			lineNumbers = append(lineNumbers, i)
		}
		board = append(board, lineNumbers)
	}
	g.boards = append(g.boards, board)

	return g
}

func (b Board) markNumber(n int) (int, int) {
	for i, row := range b {
		for j, v := range row {
			if v == n {
				b[i][j] = -1
				return j, i
			}
		}
	}
	return -1, -1
}

func (b Board) checkWin(x int, y int) bool {
	rowWon := true
	// Check row
	for i := 0; i < len(b[0]); i++ {
		if b[y][i] != -1 {
			rowWon = false
			break
		}
	}
	colWon := true
	// Check col
	for i := 0; i < len(b); i++ {
		if b[i][x] != -1 {
			colWon = false
			break
		}
	}
	return rowWon || colWon
}

func (g *Game) draw() int {
	g.currentDraw++
	drawnNumber := g.numbers[g.currentDraw]
	firstWon := -1
	for idx, b := range g.boards {
		if g.completedBoards[idx] {
			continue
		}
		drawnX, drawnY := b.markNumber(drawnNumber)
		if drawnX != -1 && drawnY != -1 {
			if b.checkWin(drawnX, drawnY) {
				if firstWon == -1 {
					firstWon = idx
				}
				g.completedBoards[idx] = true
			}
		}
	}
	return firstWon
}

func (b Board) print() {
	for _, row := range b {
		for _, v := range row {
			fmt.Printf("%3d ", v)
		}
		fmt.Printf("\n")
	}
}

func (g *Game) print() {
	for y := range g.boards[0] {
		for _, b := range g.boards {
			for _, x := range b[y] {
				fmt.Printf("%3d ", x)
			}
			fmt.Printf("\t")
		}
		fmt.Printf("\n")
	}
}

func (b Board) sumUnmarked() int {
	sum := 0
	for _, row := range b {
		for _, v := range row {
			if v != -1 {
				sum += v
			}
		}
	}
	return sum
}

func (g *Game) getFirstBoardWin() int {
	for range g.numbers {
		won := g.draw()
		if won != -1 {
			return g.boards[won].sumUnmarked() * g.numbers[g.currentDraw]
		}
	}
	return 0
}

func part1(ls []string) int {
	g := parseInput(ls)
	return g.getFirstBoardWin()
}

func (g *Game) getLastBoardWin() int {
	for range g.numbers {
		won := g.draw()
		if won != -1 && len(g.completedBoards) == len(g.boards) {
			return g.boards[won].sumUnmarked() * g.numbers[g.currentDraw]
		}

	}
	return 0
}

func part2(ls []string) int {
	g := parseInput(ls)
	return g.getLastBoardWin()
}
