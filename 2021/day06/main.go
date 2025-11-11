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

func parseInput(ls []string) []int {
	internalTimers := make([]int, 9)
	for _, l := range ls {
		for _, ns := range strings.Split(l, ",") {
			n, _ := strconv.Atoi(ns)
			internalTimers[n] += 1
		}
	}
	return internalTimers
}

func simulateDay(timers []int) []int {
	newly := timers[0]
	if newly > 0 {
		timers[0] = 0
	}
	for i := 1; i <= 8; i++ {
		timers[i-1] = timers[i]
		timers[i] = 0
	}
	timers[8] += newly
	timers[6] += newly
	return timers
}

func countSum(timers []int) int {
	sum := 0
	for _, t := range timers {
		sum += t
	}
	return sum
}

func simulateDays(timers []int, days int) []int {
	for i := 0; i < days; i++ {
		timers = simulateDay(timers)
		//fmt.Println("Day", i+1, "timers:", timers)
	}
	return timers
}

func part1(ls []string) int {
	internalTimers := parseInput(ls)
	//fmt.Println("Internal timers:", internalTimers)
	endTimers := simulateDays(internalTimers, 80)
	return countSum(endTimers)
}

func part2(ls []string) int {
	internalTimers := parseInput(ls)
	//fmt.Println("Internal timers:", internalTimers)
	endTimers := simulateDays(internalTimers, 256)
	return countSum(endTimers)
}
