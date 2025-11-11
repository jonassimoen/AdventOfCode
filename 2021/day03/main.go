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

func part1(ls []string) int {
	bitCount := make([]int, len(ls[0]))
	for _, l := range ls {
		for idx, c := range l {
			if c == '1' {
				bitCount[idx] += 1
			}
		}
	}
	gamma := 0
	epsilon := 0
	for i := 0; i < len(bitCount); i++ {
		if bitCount[i] > len(ls)/2 {
			gamma += 1 << (len(bitCount) - i - 1)
		}
	}
	epsilon = ^gamma & ((1 << len(bitCount)) - 1)
	return gamma * epsilon
}

func calculateBitCount(nrs []int, mask int) int {
	var count int
	for _, n := range nrs {
		if (n & mask) != 0 {
			count += 1
		}
	}
	return count
}

func keepNumbers(nrs []int, isOne bool, mask int) []int {
	var numbers []int
	for _, n := range nrs {
		hasOne := (n & mask) != 0
		//fmt.Printf("%05b // %05b, %t == %t\n", n, mask, isOne, hasOne)
		if isOne == hasOne {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func printBitNumbers(nrs []int) {
	for _, n := range nrs {
		fmt.Printf("%05b  ", n)
	}
	fmt.Println()
}

func rating(nrs []int, sizeNumber int, mostCommon bool) int {
	checkNumbers := nrs
	bitCount := 0
	i := 0
	for len(checkNumbers) > 1 {
		bitMask := 1 << (sizeNumber - i - 1)
		if len(checkNumbers) == 1 {
			return checkNumbers[0]
		}
		bitCount = calculateBitCount(checkNumbers, bitMask)
		complyBitMask := float64(bitCount) >= float64(len(checkNumbers))/2
		checkNumbers = keepNumbers(checkNumbers, (mostCommon && complyBitMask) || (!mostCommon && !complyBitMask), bitMask)
		i++
	}
	if len(checkNumbers) != 1 {
		panic("Something went wrong")
	}
	return checkNumbers[0]
}

func part2(ls []string) int {
	var numbers []int
	for _, l := range ls {
		number, err := strconv.ParseInt(l, 2, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(number))
	}
	numberSize := len(ls[0])
	return rating(numbers, numberSize, true) * rating(numbers, numberSize, false)
}
