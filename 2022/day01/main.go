package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
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

func part1(ls []string) int {
	idx, max_calories, curr_calories := 0, 0, 0
	for _, l := range ls {
		if l == "" {
			//fmt.Printf("Elf %d --> calories = %d\n", idx+1, elf_max_cal)
			if curr_calories > max_calories {
				max_calories = curr_calories
			}
			curr_calories = 0
			idx++
		} else {
			nr, err := strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
			//fmt.Printf("Elf %d --> calories += %d [total=%d]\n", idx+1, nr, curr_calories)
			curr_calories += nr
		}
	}
	if curr_calories > max_calories {
		max_calories = curr_calories
	}
	return max_calories
}

func part2(ls []string) int {
	var caloriesPerElf []int
	curr, idx := 0, 0
	for _, l := range ls {
		if l == "" {
			caloriesPerElf = append(caloriesPerElf, curr)
			curr = 0
			idx++
		} else {
			nr, err := strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
			curr += nr
		}
	}
	slices.Sort(caloriesPerElf)
	fmt.Println(caloriesPerElf)
	sum := 0
	for idx, val := range slices.Backward(caloriesPerElf) {
		if idx < len(caloriesPerElf)-3 {
			break
		}
		sum += val
		fmt.Println(idx, val)
	}
	return sum
}
