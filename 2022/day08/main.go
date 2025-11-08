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

type Tree struct {
	l, r, t, d     int
	value          int
	dL, dR, dT, dd int
}

func parseInput(ls []string) ([][]Tree, [][]int) {
	ss := make([][]Tree, 0)
	ssInt := make([][]int, 0)
	for li, l := range ls {
		ss = append(ss, make([]Tree, 0))
		ssInt = append(ssInt, make([]int, 0))
		for _, r := range l {
			ss[li] = append(ss[li], Tree{value: int(r - '0')})
			ssInt[li] = append(ssInt[li], int(r-'0'))
		}
	}
	return ss, ssInt
}

func convertToTreeSlice(slc [][]Tree) [][]Tree {
	//
	for x := 0; x < len(slc); x++ {
		slc[0][x].t = slc[0][x].value
		slc[len(slc)-1][x].d = slc[len(slc)-1][x].value
	}
	for y := 0; y < len(slc); y++ {
		slc[y][0].l = slc[y][0].value
		slc[y][len(slc[y])-1].r = slc[y][len(slc[y])-1].value
	}
	for y := 1; y < len(slc)-1; y++ {
		for x := 1; x < len(slc[y])-1; x++ {
			slc[y][x].t = max(slc[y-1][x].value, slc[y-1][x].t)
			slc[y][x].l = max(slc[y][x-1].value, slc[y][x-1].l)
		}
	}
	for y := len(slc) - 2; y > 0; y-- {
		for x := len(slc[y]) - 2; x > 0; x-- {
			slc[y][x].d = max(slc[y+1][x].value, slc[y+1][x].d)
			slc[y][x].r = max(slc[y][x+1].value, slc[y][x+1].r)
		}
	}
	return slc
}

func processTrees(treeSlice [][]Tree) int {
	total := (len(treeSlice)-1)*2 + (len(treeSlice[0])-1)*2
	for y := 1; y < len(treeSlice)-1; y++ {
		for x := 1; x < len(treeSlice[y])-1; x++ {
			t := treeSlice[y][x]
			value := t.value
			visible := (t.t < value) ||
				(t.d < value) ||
				(t.l < value) ||
				(t.r < value)
			if visible {
				total++
			}
		}
	}
	return total
}

func part1(ls []string) int {
	slc, _ := parseInput(ls)
	tslc := convertToTreeSlice(slc)
	return processTrees(tslc)
}

func calculateScenicScore(slice [][]int, x, y int) int {
	//fmt.Printf("%d %d = %d \n", x, y, slice[y][x])
	val := slice[y][x]
	left, top, right, bottom := 0, 0, 0, 0
	stop := false
	for xi := x - 1; (xi >= 0) && !stop; xi-- {
		//fmt.Printf("\t\tL: %d \n", slice[y][xi])
		left++
		if slice[y][xi] >= val {
			stop = true
		}
	}
	stop = false
	for yi := y - 1; (yi >= 0) && !stop; yi-- {
		//fmt.Printf("\t\tT: %d \n", slice[yi][x])
		top++
		if slice[yi][x] >= val {
			stop = true
		}
	}
	stop = false
	for xi := x + 1; (xi < len(slice[0])) && !stop; xi++ {
		//fmt.Printf("\t\tR: [%d, %d] %d %v\n", xi, y, slice[y][xi], (slice[y][xi] <= val))
		right++
		if slice[y][xi] >= val {
			stop = true
		}
	}
	stop = false
	for yi := y + 1; (yi < len(slice)) && !stop; yi++ {
		//fmt.Printf("\t\tB: %d \n", slice[yi][x])
		bottom++
		if slice[yi][x] >= val {
			stop = true
		}
	}
	//fmt.Println("\t", left, top, right, bottom)
	return left * bottom * top * right
}

func processTrees2(treeSlice [][]int) int {
	maxi := 0
	for y := 0; y < len(treeSlice); y++ {
		for x := 0; x < len(treeSlice[y]); x++ {
			maxi = max(maxi, calculateScenicScore(treeSlice, x, y))
		}
	}
	return maxi
}

func part2(ls []string) int {
	_, slc := parseInput(ls)
	//tslc := convertToTreeSlice2(slc)
	return processTrees2(slc)
}
