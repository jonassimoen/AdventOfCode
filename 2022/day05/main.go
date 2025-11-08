package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	part1ptr := flag.Int("p", 1, "Part")
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Part %d - Input file: %s\n", *part1ptr, *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	var out string

	stack, breakLine := initializeStacks(ls)
	var instr [][]int
	fmt.Println(stack)
	for i, l := range ls {
		if i == breakLine {
			reverseStack(stack)
		} else if i < breakLine {
			decodeCratesLine(l, stack)
		} else if i > breakLine+1 {
			instr = append(instr, decodeProcedure(l))
		}
	}

	if *part1ptr == 1 {
		part1(stack, instr)
	} else {
		part2(stack, instr)
	}
	out = retrieveTopCrates(stack)
	fmt.Printf("Output: %s\n", out)
}

func reverseStack(stack [][]string) {
	for i, _ := range stack {
		slices.Reverse(stack[i])
	}
}

func decodeCratesLine(l string, stack [][]string) {
	for j := 0; j < len(l); j += 4 {
		crate := l[j+1 : j+2]
		if crate != " " {
			stack[j/4] = append(stack[j/4], crate)
		}
	}
}

func initializeStacks(lines []string) ([][]string, int) {
	for idx, l := range lines {
		match, _ := regexp.MatchString(`^ [0-9]`, l)
		if match {
			fmt.Println(l[len(l)-2 : len(l)-1])
			nrStacks, _ := strconv.Atoi(l[len(l)-2 : len(l)-1])
			return make([][]string, nrStacks), idx
		}
	}
	return nil, -1
}

func decodeProcedure(l string) []int {
	r := regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`)
	matches := r.FindStringSubmatch(l)[1:]

	ncount, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])

	fmt.Printf("%d [%d, %d]\n", ncount, from, to)
	return []int{from - 1, to - 1, ncount}
}

func retrieveTopCrates(stack [][]string) string {
	tops := []string{}
	for _, s := range stack {
		tops = append(tops, s[len(s)-1])
	}
	return strings.Join(tops, "")
}

func part1(stack [][]string, instructions [][]int) {
	for _, instr := range instructions {
		//fmt.Printf("From %d to %d, %d crates\n", instr[0], instr[1], instr[2])
		for _ = range instr[2] {
			el := stack[instr[0]][len(stack[instr[0]])-1]
			stack[instr[0]] = stack[instr[0]][:len(stack[instr[0]])-1]
			stack[instr[1]] = append(stack[instr[1]], el)
			//fmt.Println(cnt, stack)
		}
	}
}

func part2(stack [][]string, instructions [][]int) {
	for _, instr := range instructions {
		//fmt.Printf("From %d to %d, %d crates\n", instr[0], instr[1], instr[2])
		//fmt.Printf("Transferring %d crate(s) from stack #%d to stack #%d\n", instr[2], instr[0], instr[1])
		//fmt.Printf("> From stack: %s, start idx = %d, end idx = %d\n", stack[instr[0]], len(stack[instr[0]])-instr[2], len(stack[instr[0]]))
		els := stack[instr[0]][len(stack[instr[0]])-instr[2] : len(stack[instr[0]])]
		//fmt.Printf("> %s\n", els)
		//fmt.Printf("> Until idx=%d\n", len(stack[instr[0]])-instr[2])
		stack[instr[0]] = stack[instr[0]][:len(stack[instr[0]])-instr[2]]
		stack[instr[1]] = append(stack[instr[1]], els...)
		//fmt.Println(stack)
	}
}
