package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

type Box struct {
	x, y, z int
	id      int
}

type Distance struct {
	i, j     int
	distance float64
}

func parseInput(ls []string) []*Box {
	bs := make([]*Box, len(ls))
	for i, l := range ls {
		split := strings.Split(l, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		bs[i] = &Box{x, y, z, i}
	}
	return bs
}

func calculateDistance(p1, p2 *Box) float64 {
	dx := float64(p2.x - p1.x)
	dy := float64(p2.y - p1.y)
	dz := float64(p2.z - p1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func checkAlreadyJoined(joins [][]int, box int) int {
	for i, join := range joins {
		if contains(join, box) {
			return i
		}
	}
	return -1
}

func part1(ls []string) int {
	boxes := parseInput(ls)

	dists := make([][]Distance, len(boxes))
	for i := range boxes {
		currDists := []Distance{}
		for j := i + 1; j < len(boxes); j++ {
			dist := calculateDistance(boxes[i], boxes[j])
			currDists = append(currDists, Distance{i, j, dist})
		}
		dists[i] = currDists
	}

	max := 10
	if len(ls) > 100 {
		max = 1000
	}
	conns := [][]int{}
	for i := 0; i < max; i++ {
		minDistance := math.MaxFloat64
		from := -1
		to := -1
		idx := -1

		for cFrom := 0; cFrom < len(dists); cFrom++ {
			for cIdx, distEntry := range dists[cFrom] {
				if distEntry.distance < minDistance {
					minDistance = distEntry.distance
					from = cFrom
					to = distEntry.j
					idx = cIdx
				}
			}
		}

		//fmt.Printf("Joining %v and %v\n", boxes[from], boxes[to])
		dists[from][idx] = Distance{from, to, math.MaxFloat64}

		a := checkAlreadyJoined(conns, from)
		b := checkAlreadyJoined(conns, to)

		if a == -1 && b == -1 {
			conns = append(conns, []int{from, to})
		} else if a == b {
			continue
		} else if a == -1 && b != -1 {
			conns[b] = append(conns[b], from)
		} else if a != -1 && b == -1 {
			conns[a] = append(conns[a], to)
		} else if a != -1 && b != -1 {
			// Both points already in groups - merge them
			conns[a] = append(conns[a], conns[b]...)
			conns = append(conns[:b], conns[b+1:]...)
		}

		//fmt.Printf("Connections: %v\n", conns)
	}
	sort.Slice(conns, func(i, j int) bool {
		return len(conns[i]) > len(conns[j])
	})

	// Multiply sizes of top 3
	result := len(conns[0]) * len(conns[1]) * len(conns[2])
	return result
}

func part2(ls []string) int {
	boxes := parseInput(ls)

	dists := make([][]Distance, len(boxes))
	for i := range boxes {
		currDists := []Distance{}
		for j := i + 1; j < len(boxes); j++ {
			dist := calculateDistance(boxes[i], boxes[j])
			currDists = append(currDists, Distance{i, j, dist})
		}
		dists[i] = currDists
	}

	conns := [][]int{}
	latestXA, latestXB := -1, -1
	for len(conns) != 1 || (len(conns) == 1 && len(conns[0]) != len(boxes)) {
		minDistance := math.MaxFloat64
		from := -1
		to := -1
		idx := -1

		for cFrom := 0; cFrom < len(dists); cFrom++ {
			for cIdx, distEntry := range dists[cFrom] {
				if distEntry.distance < minDistance {
					minDistance = distEntry.distance
					from = cFrom
					to = distEntry.j
					idx = cIdx
				}
			}
		}

		//fmt.Printf("Joining %v and %v\n", boxes[from], boxes[to])
		dists[from][idx] = Distance{from, to, math.MaxFloat64}
		latestXA = boxes[from].x
		latestXB = boxes[to].x

		a := checkAlreadyJoined(conns, from)
		b := checkAlreadyJoined(conns, to)

		if a == -1 && b == -1 {
			conns = append(conns, []int{from, to})
		} else if a == b {
			continue
		} else if a == -1 && b != -1 {
			conns[b] = append(conns[b], from)
		} else if a != -1 && b == -1 {
			conns[a] = append(conns[a], to)
		} else if a != -1 && b != -1 {
			// Both points already in groups - merge them
			conns[a] = append(conns[a], conns[b]...)
			conns = append(conns[:b], conns[b+1:]...)
		}
		//fmt.Printf("Connections: %v\n", conns)
	}

	fmt.Printf("Latest XA: %d, Latest XB: %d\n", latestXA, latestXB)
	return latestXA * latestXB
}
