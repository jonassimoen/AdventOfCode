package main

import (
	"bufio"
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

type Instruction struct {
	instruction string
	operand     string
}

type Program struct {
	instructions   []Instruction
	snowValues     map[Coordinate]bool
	programCounter int
}

type Coordinate struct {
	x, y, z int
}

func parseInput(ls []string) Program {
	instructions := []Instruction{}
	for _, l := range ls {
		split := strings.Split(l, " ")
		if len(split) > 1 {
			instructions = append(instructions, Instruction{split[0], split[1]})
		} else {
			instructions = append(instructions, Instruction{split[0], ""})
		}
	}
	return Program{instructions, map[Coordinate]bool{}, 0}
}

func parseOperand(operand string, c Coordinate) int {
	switch operand {
	case "x":
		return c.x
	case "y":
		return c.y
	case "z":
		return c.z
	default:
		v, err := strconv.Atoi(operand)
		if err != nil {
			panic(err)
		}
		return v
	}
}

func (p Program) run(c Coordinate) int {
	stack := []int{}
	p.programCounter = 0
	for p.programCounter < len(p.instructions) {
		i := p.instructions[p.programCounter]
		switch i.instruction {
		case "push":
			value := parseOperand(i.operand, c)
			stack = append(stack, value)
		case "add":
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "jmpos":
			value := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jump := parseOperand(i.operand, c)
			if value >= 0 {
				p.programCounter += jump
			}
		case "ret":
			return stack[len(stack)-1]
		}
		p.programCounter++
	}
	return -1
}

func (p *Program) countClouds() int {
	visited := map[Coordinate]bool{}
	clouds := 0

	var dfs func(Coordinate)
	dfs = func(c Coordinate) {
		if visited[c] {
			return
		}

		visited[c] = true

		nbs := []Coordinate{
			{c.x + 1, c.y, c.z},
			{c.x - 1, c.y, c.z},
			{c.x, c.y + 1, c.z},
			{c.x, c.y - 1, c.z},
			{c.x, c.y, c.z + 1},
			{c.x, c.y, c.z - 1},
		}

		for _, nb := range nbs {
			if nb.x < 0 || nb.x >= 30 || nb.y < 0 || nb.y >= 30 || nb.z < 0 || nb.z >= 30 {
				continue
			}
			if p.snowValues[nb] && !visited[nb] {
				fmt.Printf("> Searching %v\n", nb)
				dfs(nb)
			}
		}
	}

	for coord, value := range p.snowValues {
		if value && !visited[coord] {
			fmt.Printf("Searching %v\n", coord)
			clouds++
			dfs(coord)
		}
	}
	return clouds
}

func solve(inputFile string) {
	fmt.Println(inputFile)
	ls, err := readFile(inputFile)
	if err != nil {
		panic(err)
	}

	p := parseInput(ls)

	total := 0
	for x := 0; x < 30; x++ {
		for y := 0; y < 30; y++ {
			for z := 0; z < 30; z++ {
				res := p.run(Coordinate{x, y, z})
				if res > 0 {
					p.snowValues[Coordinate{x, y, z}] = true
				}
				total += res
			}
		}
	}
	clouds := p.countClouds()
	fmt.Printf("Uitkomst part 1: %d\n", total)
	fmt.Printf("Uitkomst part 2: %d\n", clouds)
	fmt.Println()
}

func main() {
	solve("./in_test")
	solve("./in")
}
