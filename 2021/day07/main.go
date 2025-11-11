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

type Program struct {
	min, max int
	values   map[int]int
}

func parseInput(ls []string) *Program {
	p := &Program{
		min:    math.MaxInt,
		max:    math.MinInt,
		values: make(map[int]int),
	}
	for _, sv := range strings.Split(ls[0], ",") {
		v, _ := strconv.Atoi(sv)
		if _, ok := p.values[v]; !ok {
			p.values[v] = 0
		}
		p.values[v]++
		if v < p.min {
			p.min = v
		}
		if v > p.max {
			p.max = v
		}
	}
	return p
}

func (p *Program) calculateMinimalFuel() (int, int) {
	minimalFuel := math.MaxInt
	minimalFuelPosition := -1
	for i := p.min; i <= p.max; i++ {
		fuel := 0
		for idx, v := range p.values {
			fuelStep := int(math.Abs(float64(idx - i)))
			fuel += fuelStep * v
		}
		if fuel < minimalFuel {
			minimalFuel = fuel
			minimalFuelPosition = i
		}
	}
	return minimalFuelPosition, minimalFuel
}

func (p *Program) calculateMinimalFuelAdvanced() (int, int) {
	minimalFuel := math.MaxInt
	minimalFuelPosition := -1
	for i := p.min; i <= p.max; i++ {
		fuel := 0
		for idx, v := range p.values {
			steps := int(math.Abs(float64(idx - i)))
			fuelStep := steps * (steps + 1) / 2
			fuel += fuelStep * v
		}
		if fuel < minimalFuel {
			minimalFuel = fuel
			minimalFuelPosition = i
		}
	}
	return minimalFuelPosition, minimalFuel
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file(s)")

	flag.Parse()

	inputFiles := strings.Split(*inputFilePtr, ",")
	for _, f := range inputFiles {
		processFile(f)
	}
}

func part1(ls []string) int {
	p := parseInput(ls)
	_, fuel := p.calculateMinimalFuel()
	return fuel
}

func part2(ls []string) int {
	p := parseInput(ls)
	_, fuel := p.calculateMinimalFuelAdvanced()
	return fuel
}
