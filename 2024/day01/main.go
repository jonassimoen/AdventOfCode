package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
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

func part1(ls []string) int {
	var aa, bb []int
	for _, l := range ls {
		split := strings.Split(l, "   ")
		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		aa = append(aa, a)
		bb = append(bb, b)
	}
	if len(aa) != len(bb) {
		panic("Wrong input")
	}
	slices.Sort(aa)
	slices.Sort(bb)

	sum := 0
	for i := 0; i < len(aa); i++ {
		a := aa[i]
		b := bb[i]
		sum += int(math.Abs(float64(a - b)))
	}
	return sum
}

func part2(ls []string) int {
	sum := 0
	var aa, bb []int
	for _, l := range ls {
		split := strings.Split(l, "   ")
		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		aa = append(aa, a)
		bb = append(bb, b)
	}

	bbb := make(map[int]int)
	for _, b := range bb {
		bbb[b]++
	}

	for _, a := range aa {
		sum += a * bbb[a]
	}

	return sum
}
