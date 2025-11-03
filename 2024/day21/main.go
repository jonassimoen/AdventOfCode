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

var numpad = [4][3]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}

var dirpad = [2][3]rune{
	{'X', '^', 'A'},
	{'<', 'v', '>'},
}

type Path = [2]rune

type ConvertGraph = map[Path]string

type Converter struct {
	codes    []string
	graphNum ConvertGraph
	graphDir ConvertGraph
	bis      map[int]map[string]int
}

func diffOrNull(a, b int) int {
	return int(math.Max(float64(a-b), float64(0)))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (c *Converter) createNumGraph() {
	c.graphNum = make(map[Path]string)
	for y1, row1 := range numpad {
		for x1, c1 := range row1 {
			for y2, row2 := range numpad {
				for x2, c2 := range row2 {
					if c1 == 'X' || c2 == 'X' {
						continue
					}
					left := strings.Repeat("<", diffOrNull(x1, x2))
					right := strings.Repeat(">", diffOrNull(x2, x1))
					up := strings.Repeat("^", diffOrNull(y1, y2))
					down := strings.Repeat("v", diffOrNull(y2, y1))
					totalPath := fmt.Sprintf("%s%s%s%s", left, down, up, right)
					if (y1 == 3 && x2 == 0) || (x1 == 3 && y2 == 0) {
						totalPath = reverse(totalPath)
					}

					c.graphNum[Path{c1, c2}] = totalPath + "A"
				}
			}
		}
	}
}
func (c *Converter) createDirGraph() {
	c.graphDir = make(map[Path]string)
	for y1, row1 := range dirpad {
		for x1, c1 := range row1 {
			for y2, row2 := range dirpad {
				for x2, c2 := range row2 {
					if c1 == 'X' || c2 == 'X' {
						continue
					}
					right := strings.Repeat(">", diffOrNull(x2, x1))
					left := strings.Repeat("<", diffOrNull(x1, x2))
					up := strings.Repeat("^", diffOrNull(y1, y2))
					down := strings.Repeat("v", diffOrNull(y2, y1))
					totalPath := fmt.Sprintf("%s%s%s%s", left, down, up, right)
					if y1 == 0 && x2 == 0 || y2 == 0 && x1 == 0 {
						totalPath = reverse(totalPath)
					}
					c.graphDir[Path{c1, c2}] = totalPath + "A"
				}
			}
		}
	}
}

func parseInput(ls []string) *Converter {
	c := Converter{}
	c.createNumGraph()
	c.createDirGraph()
	c.codes = ls
	return &c
}

func (*Converter) convertCode(seq string, usedGraph ConvertGraph) string {
	//for path, res := range usedGraph {
	//	fmt.Printf("%c ==> %c = %v\n", path[0], path[1], res)
	//}
	//fmt.Printf("Converting %s\n", seq)
	convertedSequence := ""
	origin := 'A'
	for _, s := range seq {
		p := Path{origin, s}
		//fmt.Printf("%c ==> %c = %v\n", origin, s, p)
		if path, exists := usedGraph[p]; exists {
			//fmt.Printf("%c ==> %c = %v\n", origin, s, p)
			convertedSequence += path
			origin = s
		}
	}
	return convertedSequence
}

func (c *Converter) convertAll(countDirConvertion int) int {
	totalComplexity := 0
	for _, code := range c.codes {
		convertedSequence := c.convertCode(code, c.graphNum)
		for _ = range countDirConvertion {
			convertedSequence = c.convertCode(convertedSequence, c.graphDir)
		}
		i, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		totalComplexity += len(convertedSequence) * i
	}
	return totalComplexity
}

func part1(ls []string) int {
	c := parseInput(ls)
	return c.convertAll(2)
}

func (c *Converter) getLength(seq string, iterations int, first bool) int {
	if iterations == 0 {
		return len(seq)
	}
	if _, exists := c.bis[iterations]; !exists {
		c.bis[iterations] = make(map[string]int)
	}
	prev := 'A'
	total_length := 0
	graph := c.graphDir
	if first {
		graph = c.graphNum
	}
	for _, s := range seq {
		p := Path{prev, s}
		if path, exists := graph[p]; exists {
			if l, exists := c.bis[iterations][path]; exists {
				total_length += l
				prev = s
				continue
			}
			l := c.getLength(path, iterations-1, false)
			c.bis[iterations][path] = l
			total_length += l
		}
		prev = s
	}
	return total_length
}

func (c *Converter) convertAllBis() int {
	c.bis = make(map[int]map[string]int)
	totalComplexity := 0
	for _, code := range c.codes {
		i, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			panic(err)
		}
		length := c.getLength(code, 26, true)
		totalComplexity += length * i
	}
	return totalComplexity
}

func part2(ls []string) int {
	c := parseInput(ls)
	return c.convertAllBis()
}
