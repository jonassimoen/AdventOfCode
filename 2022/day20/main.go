package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

type MixElement struct {
	value int
	mixed bool
	prev  *MixElement
	next  *MixElement
}

func parseSlice(lines []string) ([]*MixElement, *MixElement) {
	s := make([]*MixElement, 0)
	var start, curr, zeroNode *MixElement

	for _, l := range lines {
		x, _ := strconv.Atoi(l)
		newEl := &MixElement{value: x, mixed: false, prev: curr}
		if start == nil {
			start = newEl
			curr = newEl
		} else {
			curr.next = newEl
			curr = curr.next
		}
		if x == 0 {
			zeroNode = newEl
		}
		s = append(s, curr)
	}
	start.prev = curr
	curr.next = start
	return s, zeroNode
}

func displaySlice(arr []*MixElement) {
	curr := arr[0]
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", curr.value)
		curr = curr.next
	}
	fmt.Printf("\n")
}

func displaySliceExtended(arr []*MixElement) {
	curr := arr[0]
	fmt.Println("Current situation:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("\t[%11d] (%11d, %11d)\n", curr.value, curr.prev.value, curr.next.value)
		curr = curr.next
	}
	fmt.Printf("\n")
}

func mixElement(curr *MixElement, max int, arr []*MixElement) {
	steps := curr.value
	steps %= (max - 1)
	//fmt.Printf("[%d] %d \n", curr.value, steps)
	if steps < 0 {
		steps += (max - 1)
	}
	//fmt.Printf("[%d] %d \n", curr.value, steps)

	for steps > 0 {
		//fmt.Printf("STEP [%d] ", steps)
		currPrev, currNext, currNextNext := curr.prev, curr.next, curr.next.next
		curr.next = currNext.next
		curr.prev = currNext
		currPrev.next = currNext
		currNext.prev = currPrev
		currNext.next = curr
		currNextNext.prev = curr
		//fmt.Println("MIXING DONE")
		//displaySliceExtended(arr)

		//fmt.Printf("\nFinal state; \t")
		//displaySlice(arr)
		steps--
	}
	return
}

func findAfterZero(n *MixElement, amount int) int {
	curr := n
	for range amount {
		curr = curr.next
	}
	return curr.value
}

func part1(ls []string) int {
	arr, zeroEl := parseSlice(ls)

	for _, el := range arr {
		mixElement(el, len(arr), arr)
	}

	return findAfterZero(zeroEl, 1000) +
		findAfterZero(zeroEl, 2000) +
		findAfterZero(zeroEl, 3000)
}

func part2(ls []string) int {
	arr, zeroEl := parseSlice(ls)

	for i := 0; i < len(arr); i++ {
		arr[i].value *= 811589153
	}

	for range 10 {
		//displaySlice(arr)
		for _, el := range arr {
			mixElement(el, len(arr), arr)
		}
		//fmt.Println()
	}

	return findAfterZero(zeroEl, 1000) +
		findAfterZero(zeroEl, 2000) +
		findAfterZero(zeroEl, 3000)
}
