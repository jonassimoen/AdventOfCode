package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

type Cell struct {
	walkable bool
	x, y     int
	time     int
}

type Grid struct {
	Grid  [][]Cell
	Start []int
	End   []int
}

func parseInput(ls []string) *Grid {
	grid := Grid{
		Grid:  [][]Cell{},
		Start: []int{},
		End:   []int{},
	}
	for y, l := range ls {
		grid.Grid = append(grid.Grid, make([]Cell, len(l)))
		for x, c := range l {
			if c == 'S' {
				grid.Start = []int{x, y}
				grid.Grid[y][x] = Cell{walkable: true, x: x, y: y, time: 0}
			} else if c == 'E' {
				grid.End = []int{x, y}
				grid.Grid[y][x] = Cell{walkable: true, x: x, y: y, time: -1}
			} else if c == '.' {
				grid.Grid[y][x] = Cell{walkable: true, x: x, y: y, time: -1}
			} else if c == '#' {
				grid.Grid[y][x] = Cell{walkable: false, x: x, y: y, time: -1}
			}
		}
	}
	return &grid
}

func (g *Grid) getWalkableNeighbors(c Cell, distancePresent bool) []Cell {
	cells := make([]Cell, 0)
	x, y := c.x, c.y
	for _, cell := range [][]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	} {
		nX, nY := cell[0], cell[1]
		if nX < 0 || nX >= len(g.Grid[0]) || nY < 0 || nY >= len(g.Grid) {
			continue
		}
		if !g.Grid[nY][nX].walkable {
			continue
		}
		if distancePresent && g.Grid[nY][nX].time == -1 {
			continue
		}
		cells = append(cells, g.Grid[nY][nX])
	}
	return cells
}

func (g *Grid) calculateNormalWalk() int {
	queue := make([]Cell, 0)
	queue = append(queue, g.Grid[g.Start[1]][g.Start[0]])
	for len(queue) > 0 {
		cellToCheck := queue[0]
		cell := g.Grid[cellToCheck.y][cellToCheck.x]
		//fmt.Println(cell)
		queue = queue[1:]

		for _, nb := range g.getWalkableNeighbors(cell, false) {
			if cell.time+1 < nb.time || nb.time == -1 {
				g.Grid[nb.y][nb.x].time = cell.time + 1
				queue = append(queue, nb)
			}
			if nb.x == g.End[0] && nb.y == g.End[1] {
				return cell.time + 1
			}
		}
	}
	return g.Grid[g.End[1]][g.End[0]].time
}

func (g *Grid) findShortcuts() int {
	mapWonTime := make(map[int]int)
	shortcutsThatSave100ps := 0
	for _, row := range g.Grid {
		for _, cell := range row {
			if !cell.walkable {
				neighbors := g.getWalkableNeighbors(cell, true)
				if len(neighbors) < 2 {
					continue
				}
				sort.Slice(neighbors, func(i, j int) bool {
					return neighbors[i].time < neighbors[j].time
				})
				cellBeforeShortcut := neighbors[0]

				for _, cellAfterShortcut := range neighbors[1:] {
					wonTime := cellAfterShortcut.time - cellBeforeShortcut.time - 2
					if _, ok := mapWonTime[wonTime]; !ok {
						mapWonTime[wonTime] = 0
					}
					mapWonTime[wonTime]++
					if wonTime >= 100 {
						shortcutsThatSave100ps++
					}

				}
			}
		}
	}
	//for wonTime, count := range mapWonTime {
	//	fmt.Printf("Won time %d: %d\n", wonTime, count)
	//}
	return shortcutsThatSave100ps
}

func part1(ls []string) int {
	g := parseInput(ls)
	normalWalk := g.calculateNormalWalk()
	fmt.Println("Normal walk:", normalWalk)
	shortcutsThatSave100ps := g.findShortcuts()
	return shortcutsThatSave100ps
}

func (g *Grid) getWalkableNeighborsWithinDistance(c Cell, dist int) []Cell {
	cells := make([]Cell, 0)
	x, y := c.x, c.y

	for dy := -dist; dy <= dist; dy++ {
		for dx := -dist; dx <= dist; dx++ {
			nX, nY := x+dx, y+dy
			if nX < 0 || nX >= len(g.Grid[0]) || nY < 0 || nY >= len(g.Grid) {
				continue
			}
			steps := int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
			if steps > dist || steps == 0 {
				continue
			}
			if !g.Grid[nY][nX].walkable {
				continue
			}
			cells = append(cells, g.Grid[nY][nX])
		}
	}
	return cells
}

func (g *Grid) findShortcutsBis() int {
	mapWonTime := make(map[int]int)
	shortcutsThatSave100ps := 0
	for _, row := range g.Grid {
		for _, cell := range row {
			if cell.walkable {
				neighbors := g.getWalkableNeighborsWithinDistance(cell, 20)
				for _, nb := range neighbors {
					dx := nb.x - cell.x
					dy := nb.y - cell.y
					steps := int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
					// Possible to win time
					wonTime := nb.time - cell.time - steps

					if _, ok := mapWonTime[wonTime]; !ok {
						mapWonTime[wonTime] = 0
					}
					mapWonTime[wonTime]++
					if wonTime >= 100 {
						shortcutsThatSave100ps++
					}
				}
			}
		}
	}
	//wonTimes := make([]int, 0)
	//for wonTime := range mapWonTime {
	//	if wonTime >= 50 {
	//		wonTimes = append(wonTimes, wonTime)
	//	}
	//}
	//slices.Sort(wonTimes)
	//for _, wonTime := range wonTimes {
	//	fmt.Printf("Won time %d: %d\n", wonTime, mapWonTime[wonTime])
	//}
	return shortcutsThatSave100ps
}

func part2(ls []string) int {
	g := parseInput(ls)
	normalWalk := g.calculateNormalWalk()
	fmt.Println("Normal walk:", normalWalk)
	shortcutsThatSave100ps := g.findShortcutsBis()
	return shortcutsThatSave100ps
}
