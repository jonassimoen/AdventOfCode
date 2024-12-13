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

func parseInput(ls []string) ([][]int, []int) {
	var results []int
	var values [][]int
	for _, l := range ls {
		lVal := []int{}
		resValSplit := strings.Split(l, ": ")
		res, _ := strconv.Atoi(resValSplit[0])
		valSplit := strings.Split(resValSplit[1], " ")
		for _, v := range valSplit {
			val, _ := strconv.Atoi(v)
			lVal = append(lVal, val)
		}
		values = append(values, lVal)
		results = append(results, res)
		//fmt.Println(l, results, values)
	}
	return values, results
}

func findSolution(res int, values []int) bool {
	if len(values) == 1 && values[0] == res {
		return true

	}

	if len(values) == 2 {
		multiply := (values[0]*values[1] == res)
		add := (values[0]+values[1] == res)
		return multiply || add
	}

	newValues := append([]int{values[0] * values[1]}, values[2:]...)
	if findSolution(res, newValues) {
		return true
	}

	newValues = append([]int{values[0] + values[1]}, values[2:]...)
	if findSolution(res, newValues) {
		return true
	}
	return false
}

func part1(ls []string) int {
	sum := 0
	v, r := parseInput(ls)
	for i := 0; i < len(r); i++ {
		vi := v[i]
		ri := r[i]
		//fmt.Printf("Checking: %d, %v\n", ri, vi)
		ok := findSolution(ri, vi)
		if ok {
			sum += ri
		}
		fmt.Printf("Ended searching for %d, %v: %t\n", ri, vi, ok)
	}
	return sum
}

func hasSolution(res int, values []int) bool {
	if len(values) == 1 && values[0] == res {
		return true
	}

	con, _ := strconv.Atoi(fmt.Sprintf("%d%d", values[0], values[1]))
	if len(values) == 2 {
		multiply := (values[0]*values[1] == res)
		add := (values[0]+values[1] == res)
		concat := con == res
		return multiply || add || concat
	}

	newValuesMultiply := append([]int{values[0] * values[1]}, values[2:]...)
	if hasSolution(res, newValuesMultiply) {
		return true
	}

	newValuesAdd := append([]int{values[0] + values[1]}, values[2:]...)
	if hasSolution(res, newValuesAdd) {
		return true
	}

	newValuesConcat := append([]int{con}, values[2:]...)
	if hasSolution(res, newValuesConcat) {
		return true
	}

	return false
}

func part2(ls []string) int {
	sum := 0
	v, r := parseInput(ls)
	for i := 0; i < len(r); i++ {
		vi := v[i]
		ri := r[i]
		ok := hasSolution(ri, vi)
		if ok {
			sum += ri
		}
		fmt.Printf("Ended searching for %d, %v: %t\n", ri, vi, ok)
	}
	return sum
}
