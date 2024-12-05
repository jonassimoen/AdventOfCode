package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func parseMul(l string) int {
	//fmt.Println("Multiplying", l)
	mul := 1
	r := regexp.MustCompile(`[0-9]{1,3}`)
	for _, m := range r.FindAllString(l, -1) {
		g, _ := strconv.Atoi(m)
		mul *= g
	}
	return mul
}

func part1(ls []string) int {
	sum := 0
	for _, l := range ls {
		r := regexp.MustCompile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
		for _, m := range r.FindAllString(l, -1) {
			sum += parseMul(m)
		}
	}
	return sum
}

func part2(ls []string) int {
	// mul\([0-9]{1,3},[0-9]{1,3}\)|(don't\(\)|do\(\))
	sum := 0
	enabled := true
	for _, l := range ls {
		r := regexp.MustCompile("(mul\\([0-9]{1,3},[0-9]{1,3}\\))|(don't\\(\\)|do\\(\\))")
		//fmt.Println(r.FindAllString(l, -1))
		for _, m := range r.FindAllString(l, -1) {
			if m == "don't()" {
				enabled = false
				continue
			} else if m == "do()" {
				enabled = true
				continue
			}
			if enabled {
				sum += parseMul(m)
				//fmt.Println(l, sum)
			}
		}
	}
	return sum
}
