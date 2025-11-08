package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	sum := 0
	for _, l := range ls {
		s := strings.Split(l, " ")
		runeA := rune(s[0][0])
		runeB := rune(s[1][0])
		valueA := int(runeA) - int('A') + 1
		valueB := int(runeB) - int('W')
		result := 0

		switch valueB - valueA {
		case 0:
			result = 3
			//fmt.Printf("[%s] [%d %d] draw => %d + %d = %d\n", l, valueA, valueB, valueB, result, result+valueB)
		case 1, -2:
			result = 6
			//fmt.Printf("[%s] [%d %d] draw => %d + %d = %d\n", l, valueA, valueB, valueB, result, result+valueB)
		default:
			//fmt.Printf("[%s] [%d %d] A wins => %d + %d = %d\n", l, valueA, valueB, valueB, result, result+valueB)
		}

		sum += (result + valueB)
	}
	return sum
}

func part2(ls []string) int {
	sum := 0
	for _, l := range ls {
		s := strings.Split(l, " ")
		runeA := rune(s[0][0])
		runeB := rune(s[1][0])
		valueA := int(runeA) - int('A') + 1
		//valueB := int(runeB) - int('W')

		chosen := 0
		switch runeB {
		case 'X':
			chosen = valueA - 1
			if chosen == 0 {
				chosen = 3
			}
			sum += chosen
			//fmt.Printf("[%s] [%d %d] ==> %d\n", l, valueA, valueB, chosen)
			// a 1 -> z 3
			// b 2 -> x 1
			// c 3 -> y 2
		case 'Y':
			sum += (valueA + 3)
			//fmt.Printf("[%s] [%d %d] ==> %d\n", l, valueA, valueB, chosen)
			// a 1 -> x 1
			// b 2 -> y 2
			// c 3 -> z 3

		case 'Z':
			chosen = valueA + 1
			if chosen == 4 {
				chosen = 1
			}
			sum += (chosen + 6)
			//fmt.Printf("[%s] [%d %d] ==> %d\n", l, valueA, valueB, chosen)
			// a 1 -> b 2
			// b 2 -> c 3
			// c 3 -> a 1

		}
	}
	return sum
}
