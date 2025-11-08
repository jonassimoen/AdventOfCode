package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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
	x, y int
}

type Instruction struct {
	instruction string
	operand     int
}

type KPS struct {
	currentPosition  Coordinate
	currentDirection Coordinate

	instructions      []Instruction
	passedCoordinates map[Coordinate]bool
}

func parseInput(ls []string) *KPS {
	kps := &KPS{
		currentPosition:   Coordinate{0, 0},
		currentDirection:  Coordinate{0, 1},
		passedCoordinates: map[Coordinate]bool{},
	}
	instructions := []Instruction{}
	for _, l := range ls {
		split := strings.Split(l, " ")
		operand, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{split[0], operand})
	}
	kps.instructions = instructions
	kps.passedCoordinates[kps.currentPosition] = true
	return kps
}

func (p *KPS) normalizeDir() {
	if p.currentDirection.x != 0 {
		p.currentDirection.x /= int(math.Abs(float64(p.currentDirection.x)))
	}
	if p.currentDirection.y != 0 {
		p.currentDirection.y /= int(math.Abs(float64(p.currentDirection.y)))
	}
}

func (p *KPS) turn(deg int) {
	x, y := p.currentDirection.x, p.currentDirection.y

	switch deg {
	case 90:
		// Rotatie 90° rechtsom: (x,y) -> (y,-x)
		p.currentDirection = Coordinate{y, -x}
	case 180, -180:
		// Rotatie 180°: (x,y) -> (-x,-y)
		p.currentDirection = Coordinate{-x, -y}
	case -90:
		// Rotatie 90° linksom: (x,y) -> (-y,x)
		p.currentDirection = Coordinate{-y, x}
	case 45, -315:
		// (0,1) N => (1,1) NO
		// (1,1) NO => (1,0) O
		// (1,0) O => (1,-1) ZO
		// (1,-1) ZO => (0,-1) Z
		// (0,-1) Z => (-1,-1) ZW
		// (-1,-1) ZW => (-1,0) W
		// (-1,0) W => (-1,1) NW
		// (-1,1) NW => (0,1) N
		p.currentDirection = Coordinate{(x + y), (y - x)}
	case -45, 315:
		// Rotatie 45° linksom: (x,y) -> (x-y, x+y) genormaliseerd
		p.currentDirection = Coordinate{(x - y), (x + y)}
	case 135, -225:
		// (0,1) N => (1,-1) ZO
		// (1,1) NO => (0,-1) Z
		// (1,0) O => (-1,-1) ZW
		// (-1,1) ZO => (-1,0) W
		// (0,-1) Z => (-1,1) NW
		// (-1,-1) ZW => (0,1) N
		// (-1,0) W => (1,1) NO
		// (-1,1) NW => (1,0) O

		// eerst 90° rechtsom (y,-x), dan 45° rechtsom (x + y), (y - x)
		p.currentDirection = Coordinate{y + (-x), (-x) - y}
	case -135, 225:
		// Rotatie 135° linksom: (x,y) -> (x-y, x+y) maar dan nog verder
		// Of: (x,y) -> (-x-y, x-y)
		p.currentDirection = Coordinate{-y - x, -y + x}
	}
	p.normalizeDir()
	//fmt.Printf("\t{%d,%d}\n", p.currentDirection.x, p.currentDirection.y)
}

func (p *KPS) move(steps int, jump bool) {
	if jump {
		p.currentPosition.x += p.currentDirection.x * steps
		p.currentPosition.y += p.currentDirection.y * steps
		p.passedCoordinates[p.currentPosition] = true
	} else {
		for i := 0; i < steps; i++ {
			p.currentPosition.x += p.currentDirection.x
			p.currentPosition.y += p.currentDirection.y
			p.passedCoordinates[p.currentPosition] = true
		}
	}
}

func (p *KPS) run() {
	for _, i := range p.instructions {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		p.draw()
		time.Sleep(50 * time.Millisecond)
		//fmt.Printf("%d. %s [%d]\n", idx+1, i.instruction, i.operand)
		switch i.instruction {
		case "draai":
			p.turn(i.operand)
		case "spring":
			p.move(i.operand, true)
		case "loop":
			p.move(i.operand, false)
		}
	}
}

func part1(kps *KPS) {
	kps.run()
	//fmt.Println(math.Abs(float64(kps.currentPosition.x)) + math.Abs(float64(kps.currentPosition.y)))
}

func (p *KPS) draw() {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32
	for c := range p.passedCoordinates {
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	for y := maxY; y >= 0; y-- {
		builder := strings.Builder{}
		for x := minX; x <= maxX; x++ {
			coord := Coordinate{x, y}
			if p.passedCoordinates[coord] {
				builder.WriteString("#")
			} else {
				builder.WriteString(" ")
			}
		}
		fmt.Println(builder.String())
	}
}

func part2(kps *KPS) {
	kps.draw()
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
	fmt.Println(runtime.GOOS)
	//solve("./in_test")
	solve("./in")
}
