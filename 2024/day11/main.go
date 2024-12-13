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

func parseInt(ls string) []int {
	var s []int
	for _, l := range strings.Split(ls, " ") {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		s = append(s, i)
	}
	return s
}

func numberOfDigits(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func splitInTwo(n int, dig int) (int, int) {
	s := strconv.Itoa(n)
	sa, _ := strconv.Atoi(s[:dig/2])
	sb, _ := strconv.Atoi(s[dig/2:])
	return sa, sb
}

func blink(s []int) []int {
	var sn []int
	for _, si := range s {
		if si == 0 {
			sn = append(sn, 1)
			continue
		}
		dig := numberOfDigits(si)
		if dig%2 == 0 {
			a, b := splitInTwo(si, dig)
			sn = append(sn, a)
			sn = append(sn, b)
			continue
		}
		sn = append(sn, si*2024)
	}
	return sn
}

func part1(ls []string) int {
	s := parseInt(ls[0])
	//fmt.Println(s)
	i := 0
	for i < 25 {
		s = blink(s)
		//fmt.Println(s)
		i++
	}
	return len(s)
}

type XT struct {
	nr    int
	times int
}

func loopTimesNumber(xt XT, memoization *map[XT]int) int {
	if (*memoization)[xt] != 0 {
		fmt.Println(xt.nr, xt.times, (*memoization)[xt])
		return (*memoization)[xt]
	}
	ret := 1
	dig := numberOfDigits(xt.nr)
	if xt.times == 0 {
		ret = 1
	} else if xt.nr == 0 {
		ret = loopTimesNumber(XT{1, xt.times - 1}, memoization)
	} else if dig%2 == 0 {
		a, b := splitInTwo(xt.nr, dig)
		ret = (loopTimesNumber(XT{a, xt.times - 1}, memoization) +
			loopTimesNumber(XT{b, xt.times - 1}, memoization))
	} else {
		ret = loopTimesNumber(XT{xt.nr * 2024, xt.times - 1}, memoization)
	}
	(*memoization)[xt] = ret
	return ret
}

func sum(m *map[XT]int) int {
	s := 0
	for _, l := range *m {
		s += l
	}
	return s
}

func part2(ls []string) int {
	s := parseInt(ls[0])
	m := make(map[XT]int)
	ss := 0
	for _, si := range s {
		ss += loopTimesNumber(XT{
			si, 75,
		}, &m)
	}
	return ss
}
