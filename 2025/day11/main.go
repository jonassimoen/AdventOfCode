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

type Graph map[string]map[string]bool

func parseInput(ls []string) *Graph {
	g := make(Graph)
	for _, l := range ls {
		splitted := strings.Split(l, ": ")
		if len(splitted) != 2 {
			panic("Invalid input")
		}
		g[splitted[0]] = map[string]bool{}
		splitConns := strings.Split(splitted[1], " ")
		for _, c := range splitConns {
			g[splitted[0]][c] = true
		}
	}
	return &g
}

func (g *Graph) countPaths(current, end string, visited map[string]bool) int {
	if current == end {
		return 1
	}

	visited[current] = true

	paths := 0

	if nb, ok := (*g)[current]; ok {
		for n := range nb {
			if !visited[n] {
				paths += g.countPaths(n, end, visited)
			}
		}
	}

	visited[current] = false

	return paths
}

func part1(ls []string) int {
	graph := parseInput(ls)

	start := "you"
	end := "out"

	// Count paths
	count := graph.countPaths(start, end, map[string]bool{})

	fmt.Printf("Finding all paths from '%s' to '%s': %d\n\n", start, end, count)
	return count
}

func (g *Graph) countPathsBis(current, end string, memo map[string]int) int {
	if memo == nil {
		memo = make(map[string]int)
	}
	var ret int
	if current == end {
		return 1
	}
	if ret, ok := memo[current]; ok {
		return ret
	}
	for v := range (*g)[current] {
		ret += g.countPathsBis(v, end, memo)
	}

	memo[current] = ret
	return ret
}

func part2(ls []string) int {
	graph := parseInput(ls)

	start := "svr"
	end := "out"

	count := graph.countPathsBis("svr", "dac", nil) * graph.countPathsBis("dac", "fft", nil) * graph.countPathsBis("fft", "out", nil)
	countRev := graph.countPathsBis("svr", "fft", nil) * graph.countPathsBis("fft", "dac", nil) * graph.countPathsBis("dac", "out", nil)

	total := count + countRev
	fmt.Printf("Finding all paths from '%s' to '%s' that visit both dac and fft: %d\n\n", start, end, total)
	return total
}
