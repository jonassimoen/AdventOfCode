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

type Gates struct {
	inputs map[string]int
	gates  map[string][]string // z00: [x00,and,y00]
}

func parseInput(ls []string) *Gates {
	g := Gates{
		inputs: make(map[string]int),
		gates:  make(map[string][]string),
	}
	gates := true
	for _, l := range ls {
		if l == "" {
			gates = false
			continue
		}
		if gates {
			split := strings.Split(l, ": ")
			wire := split[0]
			value, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			g.inputs[wire] = value
		} else {
			split := strings.Split(l, " -> ")
			input := strings.Split(split[0], " ")
			outWire := split[1]
			g.gates[outWire] = input
		}
	}
	return &g
}

func (g *Gates) calculateOutput(wire string) int {
	wireLeft := g.gates[wire][0]
	operand := g.gates[wire][1]
	wireRight := g.gates[wire][2]

	inValueLeft, ok := g.inputs[wireLeft]
	if !ok {
		panic("Input not found")
	}

	invalueRight, ok := g.inputs[wireRight]
	if !ok {
		panic("Input not found")
	}

	switch operand {
	case "AND":
		return inValueLeft & invalueRight
	case "XOR":
		return inValueLeft ^ invalueRight
	case "OR":
		return inValueLeft | invalueRight
	}
	return -1
}

func (g *Gates) calculateResult() int {
	zKeys := []string{}
	for k := range g.gates {
		if k[0] == 'z' {
			zKeys = append(zKeys, k)
		}
	}
	sort.Strings(zKeys)
	res := 0
	for i, k := range zKeys {
		value := g.inputs[k]
		if value == 1 {
			res += 1 << i
		}
	}
	return res
}

func part1(ls []string) int {
	g := parseInput(ls)

	outputs := []string{}
	for k := range g.gates {
		outputs = append(outputs, k)
	}

	for len(outputs) > 0 {
		output := outputs[0]
		outputs = outputs[1:]

		//fmt.Printf("Checking %s: needs inputs %v\n", output, g.gates[output])
		inA := g.gates[output][0]
		inB := g.gates[output][2]
		_, okA := g.inputs[inA]
		//if !okA {
		//	fmt.Printf("\tInput %s not found\n", inA)
		//}
		_, okB := g.inputs[inB]
		//if !okB {
		//	fmt.Printf("\tInput %s not found\n", inB)
		//}
		if okA && okB {
			outValue := g.calculateOutput(output)
			g.inputs[output] = outValue
			fmt.Println(output, g.gates[output], output)
		} else {
			outputs = append(outputs, output)
		}
	}

	for k, v := range g.inputs {
		fmt.Println(k, ": ", v)
	}
	return g.calculateResult()
}

func part2(ls []string) int {
	return -1
}
