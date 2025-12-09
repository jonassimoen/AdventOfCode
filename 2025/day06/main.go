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

type Program struct {
	values     [][]int
	operations []rune
}

func (p *Program) calculate() int {
	total := 0
	for idx, op := range p.operations {
		result := -1
		for i := 0; i < len(p.values); i++ {
			value := p.values[i][idx]
			if result == -1 {
				result = value
			} else {
				switch op {
				case '*':
					result *= value
				case '+':
					result += value
				}
			}
		}
		total += result
	}
	return total
}

func (p *Program) calculateBis() int {
	total := 0
	for idx, op := range p.operations {
		result := -1
		for _, value := range p.values[idx] {
			if result == -1 {
				result = value
			} else {
				switch op {
				case '*':
					result *= value
				case '+':
					result += value
				}
			}
		}
		total += result
	}
	return total
}

func parseInput(ls []string) *Program {
	p := &Program{
		values:     [][]int{},
		operations: []rune{},
	}
	for idx, l := range ls {
		if idx == len(ls)-1 {
			break
		}
		lValues := []int{}
		split := strings.Split(l, " ")
		for _, v := range split {
			if v == "" {
				continue
			}
			val, _ := strconv.Atoi(v)
			lValues = append(lValues, val)
		}
		p.values = append(p.values, lValues)
	}
	for _, r := range ls[len(ls)-1] {
		if r == ' ' {
			continue
		}
		p.operations = append(p.operations, r)
	}
	return p
}

func part1(ls []string) int {
	p := parseInput(ls)
	return p.calculate()
}

func parseInputBis(ls []string) *Program {
	lineLen := len(ls[0])
	nLines := len(ls)
	p := &Program{
		values:     [][]int{},
		operations: []rune{},
	}
	tempValues := []int{}
	for i := lineLen - 1; i >= 0; i-- {
		value := 0
		for n := 0; n < nLines-1; n++ {
			r := ls[n][i]
			if r == ' ' {
				continue
			}
			value = value*10 + int(r-'0')
		}
		tempValues = append(tempValues, value)
		if ls[nLines-1][i] != ' ' {
			p.values = append(p.values, tempValues)
			p.operations = append(p.operations, rune(ls[nLines-1][i]))
			tempValues = []int{}
			i--
		}
	}
	if len(tempValues) > 0 {
		p.values = append(p.values, tempValues)
	}
	return p
}

func part2(ls []string) int {
	p := parseInputBis(ls)
	return p.calculateBis()
}
