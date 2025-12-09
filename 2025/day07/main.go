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

type Set []int

func (s *Set) add(x int) {
	exists := false
	for _, v := range *s {
		if v == x {
			exists = true
			break
		}
	}
	if !exists {
		*s = append(*s, x)
	}
}

func part1(ls []string) int {
	w := len(ls[0])
	lineBeams := Set{}
	for idx, char := range ls[0] {
		if char == 'S' {
			lineBeams.add(idx)
		}
	}

	splits := 0
	for _, l := range ls[1:] {
		newBeams := Set{}
		for _, beam := range lineBeams {
			if l[beam] == '.' {
				newBeams.add(beam)
			} else if l[beam] == '^' {
				if beam-1 >= 0 {
					newBeams.add(beam - 1)
				}
				if beam+1 < w {
					newBeams.add(beam + 1)
				}
				splits++
			}
		}
		lineBeams = newBeams
		//fmt.Printf("Beams on line %d: %v\n", idx+2, lineBeams)
	}
	return splits
}

func part2(ls []string) int {
	beam := 0
	for idx, char := range ls[0] {
		if char == 'S' {
			beam = idx
		}
	}

	p := Program{ls: ls, cache: map[[2]int]int{}}

	return p.calculatePaths(1, beam)
}

type Program struct {
	ls    []string
	cache map[[2]int]int
}

func (p *Program) calculatePaths(level int, beam int) int {
	if level >= len(p.ls) {
		return 1
	}

	if beam < 0 || beam >= len(p.ls[level]) {
		return 0
	}

	key := [2]int{level, beam}
	if c, ok := p.cache[key]; ok {
		return c
	}

	for level < len(p.ls) && p.ls[level][beam] == '.' {
		level++
	}

	if level >= len(p.ls) {
		p.cache[key] = 1
		return 1
	}

	for p.ls[level][beam] == '.' {
		level++
		if level >= len(p.ls) {
			p.cache[key] = 1 // Use original key (origLevel, beam)
			return 1
		}
	}

	total := p.calculatePaths(level+1, beam-1) + p.calculatePaths(level+1, beam+1)

	p.cache[key] = total // Use original key (origLevel, beam)
	return total
}
