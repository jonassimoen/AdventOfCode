package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
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

type Robot struct {
	x, y   int
	vx, vy int
}

func parseInput(ls []string, w, h int) Grid {
	var rs []*Robot
	for _, l := range ls {
		st := strings.Split(l, " ")
		x, _ := strconv.Atoi(strings.Split(st[0], ",")[0][2:])
		y, _ := strconv.Atoi(strings.Split(st[0], ",")[1])
		vx, _ := strconv.Atoi(strings.Split(st[1], ",")[0][2:])
		vy, _ := strconv.Atoi(strings.Split(st[1], ",")[1])
		rs = append(rs, &Robot{x, y, vx, vy})
	}
	return Grid{w, h, rs}
}

type Grid struct {
	maxW, maxH int
	rs         []*Robot
}

func (g *Grid) oneSecondIteration() {
	for _, r := range g.rs {
		r.x += r.vx
		r.y += r.vy

		if r.x >= g.maxW {
			r.x %= g.maxW
		}
		if r.x < 0 {
			r.x = (r.x % g.maxW) + g.maxW
		}
		if r.y >= g.maxH {
			r.y %= g.maxH
		}
		if r.y < 0 {
			r.y = (r.y % g.maxH) + g.maxH
		}
	}
}

func (g *Grid) print() {
	pg := make([][]int, g.maxH)
	for i := range pg {
		pg[i] = make([]int, g.maxW)
	}
	for _, r := range g.rs {
		pg[r.y][r.x]++
	}
	for y := 0; y < g.maxH; y++ {
		for x := 0; x < g.maxW; x++ {
			if pg[y][x] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", pg[y][x])
			}
		}
		fmt.Println()
	}
}

func (g *Grid) countRobotsInQuadrants() int {
	xSplit := g.maxW / 2
	ySplit := g.maxH / 2
	tl, tr, bl, br := 0, 0, 0, 0
	for _, r := range g.rs {
		if r.x == xSplit || r.y == ySplit {
			continue
		}
		if r.x < xSplit {
			if r.y < ySplit {
				tl++
			} else {
				tr++
			}
		} else {
			if r.y < ySplit {
				bl++
			} else {
				br++
			}
		}
	}
	return tl * tr * bl * br
}

func part1(ls []string) int {
	g := parseInput(ls, 101, 103)
	for i := 0; i < 100; i++ {
		g.oneSecondIteration()
		//g.print()
		//fmt.Println()
	}
	return g.countRobotsInQuadrants()
}

func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Grid) createImage(i int) {
	white := color.RGBA{255, 255, 255, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{g.maxW, g.maxH}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for y := 0; y < g.maxH; y++ {
		for x := 0; x < g.maxW; x++ {
			img.Set(x, y, black)
		}
	}

	for _, r := range g.rs {
		img.Set(r.x, r.y, white)
	}
	f, _ := os.Create(fmt.Sprintf("day14/image%04d.png", i))
	png.Encode(f, img)
}

func (g *Grid) isInArray(x int, y int) bool {
	if x < 0 || x >= g.maxW || y < 0 || y >= g.maxH {
		return false
	}
	for _, r := range g.rs {
		if r.x == x && r.y == y {
			return true
		}
	}
	return false
}

func (g *Grid) checkChristmasTree() bool {
	for _, r := range g.rs {
		x, y := r.x, r.y
		below1 := g.isInArray(x-1, y-1) && g.isInArray(x, y-1) && g.isInArray(x+1, y-1)
		if !below1 {
			continue
		}
		below2 := g.isInArray(x-1, y-2) && g.isInArray(x-1, y-2) && g.isInArray(x, y-2) && g.isInArray(x+1, y-2) && g.isInArray(x+2, y-2)
		if !below2 {
			continue
		}
		return true
	}
	return false
}

func part2(ls []string) int {
	g := parseInput(ls, 101, 103)
	for i := 0; i < 9999; i++ {
		g.oneSecondIteration()

		if g.checkChristmasTree() {
			g.createImage(i + 1)
			return i + 1
		}

		//clearScreen()
		//fmt.Printf("=== %d ===\n\n", i+1)
		//g.print()
		//time.Sleep(time.Second)

		//g.createImage(i + 1)
	}
	return -1
}
