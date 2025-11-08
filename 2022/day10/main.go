package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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
	var out int
	if *part1ptr == 1 {
		out = part1(ls)
	} else {
		out = part2(ls)
	}
	fmt.Printf("Output: %d\n", out)
}

type cycleState struct {
	cycleNumber   int
	registerValue int
	sum           int
}

func readOperations(operations chan int, addition chan int) {
	for {
		op := <-operations
		//fmt.Printf("Operation: %d\n", op)
		switch {
		case op == 0:
			addition <- 0
		default:
			addition <- 0
			addition <- op
		}
	}
}

func processAddition(state *cycleState, addition chan int) {
	checkCycles := []int{20, 60, 100, 140, 180, 220}
	for {
		add := <-addition
		state.registerValue += add
		state.cycleNumber++
		//fmt.Printf("New cycle: %d, add %d ==> register = %d\n", state.cycleNumber, add, state.registerValue)
		for _, cc := range checkCycles {
			if cc == state.cycleNumber {
				state.sum += (state.registerValue * state.cycleNumber)
				fmt.Printf("Cycle: %d ==> %d ==> %d\n", cc, state.registerValue, state.sum)
			}
		}
	}
}

func part1(ls []string) int {
	operations := make(chan int) // 0 = noop, <>0 = addx
	addition := make(chan int)

	state := cycleState{cycleNumber: 1, registerValue: 1, sum: 0}
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		readOperations(operations, addition)
	}()

	go func() {
		defer wg.Done()
		processAddition(&state, addition)
	}()

	for _, l := range ls {
		s := strings.Split(l, " ")
		if s[0] == "addx" {
			n, _ := strconv.Atoi(s[1])
			operations <- n
		} else {
			operations <- 0
		}
	}
	return state.sum
}

func processCycle(state *cycleState, addition chan int, sb *strings.Builder) {
	for {
		sb.WriteRune(getPixel(state))
		fmt.Println(sb.String())
		fmt.Println("-")
		add := <-addition
		state.cycleNumber++
		state.registerValue += add
		if state.cycleNumber%40 == 1 {
			sb.WriteRune('\n')
		}
	}
}

func getPixel(state *cycleState) rune {
	pixelPosition := (state.cycleNumber - 1) % 40
	//fmt.Printf("Cycle: %d [ pixelpos =  %d, register = %d ]\n", state.cycleNumber, pixelPosition, state.registerValue)
	//fmt.Printf("\t Low bound = %d, upper bound = %d\n", max(0, state.registerValue-1), min(state.registerValue+1, 40))
	if (max(0, state.registerValue-1) <= pixelPosition) && (min(state.registerValue+1, 40) >= pixelPosition) {
		return '#'
	} else {
		return '.'
	}
}

func part2(ls []string) int {
	operations := make(chan int) // 0 = noop, <>0 = addx
	addition := make(chan int)

	state := cycleState{cycleNumber: 1, registerValue: 1, sum: 0}
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		readOperations(operations, addition)
	}()

	sb := strings.Builder{}
	go func() {
		defer wg.Done()
		processCycle(&state, addition, &sb)
	}()

	for _, l := range ls {
		s := strings.Split(l, " ")
		if s[0] == "addx" {
			n, _ := strconv.Atoi(s[1])
			operations <- n
		} else {
			operations <- 0
		}
	}
	fmt.Printf("Final string: \n%s\n", sb.String())
	return state.sum
}
