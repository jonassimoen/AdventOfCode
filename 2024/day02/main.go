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
	sum := 0
	for _, l := range ls {
		split := strings.Split(l, " ")
		prev, _ := strconv.Atoi(split[0])
		cur, _ := strconv.Atoi(split[1])
		inc := cur >= prev
		//fmt.Printf("Line %d is inc: %b\n", li, inc)
		for i := 1; i < len(split); i++ {
			cur, _ = strconv.Atoi(split[i])
			diff := math.Abs(float64(cur - prev))
			if diff > 3 || diff < 1 || ((inc && (cur < prev)) || (!inc && (cur > prev))) {
				//fmt.Printf("%d - %d = NOT OK\n", prev, cur)
				break
			} else if i == len(split)-1 {
				sum += 1
			}
			//fmt.Printf("%d - %d = OK\n", prev, cur)
			prev = cur
		}
	}
	return sum
}

func isValid(a, b int, increasing bool) bool {
	diff := math.Abs(float64(a) - float64(b))
	return (diff <= 3) && (diff >= 1) && ((increasing && (b > a)) || (!increasing && (a > b)))
}

func remove(slice []int, s int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:s]...)
	return append(ret, slice[s+1:]...)
}

func checkSequence(arr []int, idxA, idxB int, increasing, badLevelPossible bool, idx int) bool {
	if len(arr) == idxA || len(arr) == idxB {
		return true
	}

	a := arr[idxA]
	b := arr[idxB]
	val := isValid(a, b, increasing)
	if !val {
		if badLevelPossible {
			return checkSequence(remove(arr, idxA+1), 0, 1, increasing, false, idx+1) ||
				checkSequence(remove(arr, idxA), 0, 1, increasing, false, idx+1)
		}
		return false
	}

	return checkSequence(arr, idxA+1, idxB+1, increasing, badLevelPossible, idx+1)
}

func part2(ls []string) int {
	sum := 0
	for _, l := range ls {
		var arr []int
		for _, s := range strings.Split(l, " ") {
			val, _ := strconv.Atoi(s)
			arr = append(arr, val)
		}
		if checkSequence(arr, 0, 1, true, true, 0) ||
			checkSequence(arr, 0, 1, false, true, 0) {
			sum += 1
		}
	}
	return sum
}
