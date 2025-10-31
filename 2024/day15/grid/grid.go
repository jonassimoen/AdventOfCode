package grid

import (
	"fmt"
	"time"
)

type Cell struct {
	typ int // 0 - nothing, 1 - wall, 2 - box, 3 - robot
}

type Direction [2]int

var TOP = Direction{0, -1}
var BOTTOM = Direction{0, 1}
var RIGHT = Direction{1, 0}
var LEFT = Direction{-1, 0}

type Grid struct {
	cells      [][]Cell
	robotMoves []Direction
	robCurPos  Direction
}

func ParseInput(ls []string) *Grid {
	idx := 0
	s := Direction{}
	var g [][]Cell
	for ls[idx] != "" {
		g = append(g, []Cell{})
		for idx2, c := range ls[idx] {
			if c == '@' {
				s = Direction{idx2, idx}
			}
			g[idx] = append(g[idx], Cell{parsePosition(c)})
		}
		idx++
	}
	idx++
	var m []Direction
	for idx < len(ls) {
		for _, d := range ls[idx] {
			m = append(m, ParseDirection(d))
		}
		idx++
	}
	return &Grid{g, m, s}
}

func parsePosition(c rune) int {
	switch c {
	case '#':
		return 1
	case '@':
		return 3
	case 'O':
		return 2
	case '.':
		return 0
	default:
		return -1
	}
}

func unparsePosition(c int) rune {
	switch c {
	case 1:
		return '#'
	case 3:
		return '@'
	case 2:
		return 'O'
	case 21:
		return '['
	case 22:
		return ']'
	case 0:
		return '.'
	default:
		return 'X'
	}
}

func ParseDirection(c rune) Direction {
	switch c {
	case '<':
		return LEFT
	case '>':
		return RIGHT
	case '^':
		return TOP
	case 'v':
		return BOTTOM
	default:
		return Direction{0, 0}
	}
}

func UnparseDirection(c Direction) rune {
	switch c {
	case LEFT:
		return '<'
	case RIGHT:
		return '>'
	case TOP:
		return '^'
	case BOTTOM:
		return 'v'
	default:
		return '?'
	}
}

func (g *Grid) Print() {
	for _, r := range g.cells {
		for _, d := range r {
			fmt.Printf("%c", unparsePosition(d.typ))
		}
		fmt.Println()
	}
}

func (g *Grid) StartMoving() {
	for _, d := range g.robotMoves {
		newPos := Direction{g.robCurPos[0] + d[0], g.robCurPos[1] + d[1]}
		//fmt.Printf("%c\n", UnparseDirection(d))
		if g.cells[newPos[1]][newPos[0]].typ == 1 {
			//Wall, can't do anything

		} else if g.cells[newPos[1]][newPos[0]].typ == 2 {
			// Box at newPosition, should push
			if g.push(newPos, d) {
				g.cells[newPos[1]][newPos[0]].typ = 3
				g.cells[g.robCurPos[1]][g.robCurPos[0]].typ = 0
				g.robCurPos = newPos
			}
		} else if g.cells[newPos[1]][newPos[0]].typ == 0 {
			// Nothing at newPos, just move
			g.cells[newPos[1]][newPos[0]].typ = 3
			g.cells[g.robCurPos[1]][g.robCurPos[0]].typ = 0
			g.robCurPos = newPos
		}
		//g.Print()
	}
}

func (g *Grid) push(p Direction, d Direction) bool {
	nextPos := Direction{p[0] + d[0], p[1] + d[1]}
	if g.cells[nextPos[1]][nextPos[0]].typ == 0 {
		g.cells[p[1]][p[0]].typ = 0
		g.cells[nextPos[1]][nextPos[0]].typ = 2
		return true
	}
	if g.cells[nextPos[1]][nextPos[0]].typ == 2 {
		pp := g.push(nextPos, d)
		if pp {
			g.cells[p[1]][p[0]].typ = 0
			g.cells[nextPos[1]][nextPos[0]].typ = 2
			return true
		}
		return false
	}
	return false
}

func (g *Grid) CalculateSumCoordinates() int {
	sum := 0
	for i := 0; i < len(g.cells); i++ {
		for j := 0; j < len(g.cells[i]); j++ {
			if g.cells[i][j].typ == 2 {
				x := i*100 + j
				sum += x
			}
		}
	}
	return sum
}

// Part 2

func (g *Grid) DoubleGrid() {
	var grid [][]Cell
	for y := 0; y < len(g.cells); y++ {
		x := 0
		grid = append(grid, make([]Cell, len(g.cells[y])*2))
		for x < len(g.cells[y]) {
			if g.cells[y][x].typ == 0 {
				grid[y][2*x].typ = 0
				grid[y][2*x+1].typ = 0
			} else if g.cells[y][x].typ == 1 {
				grid[y][2*x].typ = 1
				grid[y][2*x+1].typ = 1
			} else if g.cells[y][x].typ == 2 {
				grid[y][2*x].typ = 21
				grid[y][2*x+1].typ = 22
			} else if g.cells[y][x].typ == 3 {
				grid[y][2*x].typ = 3
				grid[y][2*x+1].typ = 0
			}
			x++
		}
	}
	g.cells = grid
	g.robCurPos = Direction{g.robCurPos[0] * 2, g.robCurPos[1]}
}

func (g *Grid) printStep(move Direction, step int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("\033[H\033[2J")
	fmt.Printf("Move: %c, step: %d\n", UnparseDirection(move), step)
	g.Print()
}

func (g *Grid) StartMovingBis() {
	for _, d := range g.robotMoves {
		newPos := Direction{g.robCurPos[0] + d[0], g.robCurPos[1] + d[1]}
		if g.cells[newPos[1]][newPos[0]].typ == 1 {
			//Wall, can't do anything

		} else if g.cells[newPos[1]][newPos[0]].typ == 21 || g.cells[newPos[1]][newPos[0]].typ == 22 {
			// Box at newPosition, should push
			if g.isAbleToMove(newPos, d) {
				if d[1] != 0 {
					g.moveAllBoxesVertical(newPos, d)
				} else {
					g.moveAllBoxesHorizontal(newPos, d)
				}
				g.cells[newPos[1]][newPos[0]].typ = 3
				g.cells[g.robCurPos[1]][g.robCurPos[0]].typ = 0
				g.robCurPos = newPos
			}

		} else if g.cells[newPos[1]][newPos[0]].typ == 0 {
			// Nothing at newPos, just move
			g.cells[newPos[1]][newPos[0]].typ = 3
			g.cells[g.robCurPos[1]][g.robCurPos[0]].typ = 0
			g.robCurPos = newPos
		}
		//g.printStep(d, s)
	}
}

func (g *Grid) isAbleToMove(p Direction, d Direction) bool {
	if g.cells[p[1]][p[0]].typ == 0 {
		return true
	}

	if d[1] != 0 {
		// VERTICAL
		isLeftPart := g.cells[p[1]][p[0]].typ == 21
		var leftPos, rightPos Direction
		if isLeftPart {
			leftPos = Direction{p[0], p[1]}
			rightPos = Direction{p[0] + 1, p[1]}
		} else {
			leftPos = Direction{p[0] - 1, p[1]}
			rightPos = Direction{p[0], p[1]}
		}

		if (g.cells[leftPos[1]+d[1]][leftPos[0]+d[0]].typ == 0) && (g.cells[rightPos[1]+d[1]][rightPos[0]+d[0]].typ == 0) {
			// Box can just move
			return true
		}

		if (g.cells[leftPos[1]+d[1]][leftPos[0]+d[0]].typ == 1) || (g.cells[rightPos[1]+d[1]][rightPos[0]+d[0]].typ == 1) {
			// Box will never be able ==> wall is blocking at least the mid
			return false
		}

		// Check the parents
		return g.isAbleToMove(Direction{leftPos[0] + d[0], leftPos[1] + d[1]}, d) && g.isAbleToMove(Direction{rightPos[0] + d[0], rightPos[1] + d[1]}, d)
	} else {
		// HORIZONTAL
		if (p[0]+2*d[0] < 0) || (p[0]+2*d[0] >= len(g.cells[p[1]])) {
			return false
		}
		if g.cells[p[1]+d[1]][p[0]+2*d[0]].typ == 0 {
			return true
		} else {
			return g.isAbleToMove(Direction{p[0] + 2*d[0], p[1] + d[1]}, d)
		}
	}
	return false
}

func (g *Grid) moveFullBox(p, d Direction) {
	isLeftPart := g.cells[p[1]][p[0]].typ == 21
	var leftPos, rightPos Direction
	if isLeftPart {
		leftPos = Direction{p[0], p[1]}
		rightPos = Direction{p[0] + 1, p[1]}
	} else {
		leftPos = Direction{p[0] - 1, p[1]}
		rightPos = Direction{p[0], p[1]}
	}
	g.cells[leftPos[1]][leftPos[0]].typ = 0
	g.cells[rightPos[1]][rightPos[0]].typ = 0
	g.cells[leftPos[1]+d[1]][leftPos[0]+d[0]].typ = 21
	g.cells[rightPos[1]+d[1]][rightPos[0]+d[0]].typ = 22
}

func (g *Grid) moveAllBoxesVertical(p, d Direction) {
	isLeftPart := g.cells[p[1]][p[0]].typ == 21
	var leftPos, rightPos Direction
	if isLeftPart {
		leftPos = Direction{p[0], p[1]}
		rightPos = Direction{p[0] + 1, p[1]}
	} else {
		leftPos = Direction{p[0] - 1, p[1]}
		rightPos = Direction{p[0], p[1]}
	}
	if g.cells[leftPos[1]+d[1]][leftPos[0]+d[0]].typ == 21 || g.cells[leftPos[1]+d[1]][leftPos[0]+d[0]].typ == 22 {
		g.moveAllBoxesVertical(Direction{leftPos[0] + d[0], leftPos[1] + d[1]}, d)
	}
	if g.cells[leftPos[1]+d[1]][rightPos[0]+d[0]].typ == 21 || g.cells[rightPos[1]+d[1]][leftPos[0]+d[0]].typ == 22 {
		g.moveAllBoxesVertical(Direction{rightPos[0] + d[0], rightPos[1] + d[1]}, d)
	}
	g.moveFullBox(leftPos, d)
	//fmt.Printf("======\nMoving %v\n", leftPos)
	//g.print()
}

func (g *Grid) moveAllBoxesHorizontal(p, d Direction) {
	if g.cells[p[1]+d[1]][p[0]+2*d[0]].typ == 21 || g.cells[p[1]+d[1]][p[0]+2*d[0]].typ == 22 {
		g.moveAllBoxesHorizontal(Direction{p[0] + 2*d[0], p[1] + d[1]}, d)
	}
	g.moveFullBox(p, d)
}

func (g *Grid) CalculateSumCoordinatesBis() int {
	sum := 0
	for i := 0; i < len(g.cells); i++ {
		for j := 0; j < len(g.cells[i]); j++ {
			if g.cells[i][j].typ == 21 {
				x := i*100 + j
				sum += x
			}
		}
	}
	return sum
}
