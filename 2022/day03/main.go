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
		compA, compB := l[:len(l)/2], l[len(l)/2:]
		//fmt.Printf("Compartiment A: %s\n", compA)
		//fmt.Printf("Compartiment B: %s\n", compB)

		for i := 0; i < len(compA); i++ {
			charA := rune(compA[i])
			if strings.ContainsRune(compB, charA) {
				intA := int(charA)
				if (intA >= int('A')) && (intA <= int('Z')) {
					//fmt.Printf("Found char %s, priority= %d\n", string(charA), (intA-int('A'))+27)
					sum += (intA - int('A') + 27)
				} else {
					//fmt.Printf("Found char %s, priority= %d\n", string(charA), (intA - int('a') + 1))
					sum += (intA - int('a') + 1)
				}
				break
				//charAsum += int(rune(charA) - )
			}
		}
	}

	return sum
}

func part2(ls []string) int {
	sum := 0
	for i := 0; i < len(ls); i += 3 {
		lineA := ls[i]
		lineB := ls[i+1]
		lineC := ls[i+2]

		for i := 0; i < len(lineA); i++ {
			charA := rune(lineA[i])
			if strings.ContainsRune(lineB, charA) && strings.ContainsRune(lineC, charA) {
				intA := int(charA)
				if (intA >= int('A')) && (intA <= int('Z')) {
					//fmt.Printf("Found char %s, priority= %d\n", string(charA), (intA-int('A'))+27)
					sum += (intA - int('A') + 27)
				} else {
					//fmt.Printf("Found char %s, priority= %d\n", string(charA), (intA - int('a') + 1))
					sum += (intA - int('a') + 1)
				}
				break
			}
		}
	}
	return sum
}
