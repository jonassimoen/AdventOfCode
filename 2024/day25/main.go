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

type Input struct {
	keys   [][]int
	locks  [][]int
	height int
}

func (i *Input) convertObject(object []int, objectType string) {
	if object != nil {
		if objectType == "lock" {
			i.locks = append(i.locks, object)
		} else {
			for i, v := range object {
				object[i] = v - 1
			}
			i.keys = append(i.keys, object)
		}
	}
}

func parseInput(ls []string) Input {
	input := Input{}
	var object []int
	var objectType string
	startHeight := 0
	for idx, line := range ls {
		if line == "" {
			if startHeight == 0 {
				startHeight = idx
				input.height = idx
			}
			input.convertObject(object, objectType)
			object = nil
			objectType = ""
			continue
		}

		if object == nil {
			isLock := true
			for _, c := range line {
				if c != '#' {
					isLock = false
					break
				}
			}
			if isLock {
				objectType = "lock"
			}
			object = make([]int, len(line))
		} else {
			// Currently building an object
			for i, value := range line {
				if value == '#' {
					object[i] += 1
				}
			}
		}
	}
	input.convertObject(object, objectType)
	return input
}

func (i *Input) checkOverlapForLock(lockIdx int) int {
	lock := i.locks[lockIdx]
	numOverlap := 0
	for _, key := range i.keys {
		fmt.Printf("Comparing key %v with lock %v\n", key, lock)
		overlap := true
		for p := range lock {
			overlap = overlap && ((key[p] + lock[p]) <= (i.height - 2))
			fmt.Printf("\tChecking %d: %d + %d = %d // %d ==> %t\n", p, key[p], lock[p], key[p]+lock[p], i.height-2, overlap)
		}
		if overlap {
			fmt.Printf("Overlap found for lock %v and key %v\n", lock, key)
			numOverlap++
		}
	}
	return numOverlap
}

func part1(ls []string) int {
	input := parseInput(ls)
	for _, k := range input.keys {
		fmt.Println("key", k)
	}
	for _, l := range input.locks {
		fmt.Println("lock", l)
	}
	fmt.Println("height", input.height)

	totalOverlap := 0
	for lockIdx := range input.locks {
		totalOverlap += input.checkOverlapForLock(lockIdx)
	}
	return totalOverlap
}

func part2(ls []string) int {
	return -1
}
