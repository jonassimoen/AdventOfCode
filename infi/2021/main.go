package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

type Program struct {
	parts        map[string]Parts
	missingParts int

	partsCount map[string]int
	subparts   map[string]bool
	mainparts  []string
}

type Parts map[string]int

func parseInput(ls []string) *Program {
	p := &Program{
		parts:      map[string]Parts{},
		partsCount: map[string]int{},
		subparts:   map[string]bool{},
		mainparts:  []string{},
	}

	missingParts, err := strconv.Atoi(strings.Split(ls[0], " ")[0])
	if err != nil {
		panic(err)
	}
	p.missingParts = missingParts

	for _, l := range ls[1:] {
		splitLine := strings.Split(l, ": ")
		play := splitLine[0]
		splitSubParts := strings.Split(splitLine[1], ", ")

		p.parts[play] = map[string]int{}
		for _, subpart := range splitSubParts {
			subpartSplit := strings.Split(subpart, " ")
			subpartCount, err := strconv.Atoi(subpartSplit[0])
			subpartName := subpartSplit[1]
			if err != nil {
				panic(err)
			}
			p.parts[play][subpartName] = subpartCount
			p.subparts[subpartName] = true
		}
		//fmt.Println(l)
	}

	for part := range p.parts {
		if _, ok := p.subparts[part]; !ok {
			p.mainparts = append(p.mainparts, part)
		}
	}

	return p
}

func (p *Program) calculateParts(part string) int {
	if count, ok := p.partsCount[part]; ok {
		return count
	}

	total := 0
	for subPartName, subPartCount := range p.parts[part] {
		if _, ok := p.parts[subPartName]; !ok {
			total += subPartCount
		} else {
			total += subPartCount * p.calculateParts(subPartName)
		}
	}

	p.partsCount[part] = total
	return total
}

func part1(p *Program) {
	maxAantal := 0
	maxAantalPart := ""
	for part := range p.parts {
		count := p.calculateParts(part)
		if count > maxAantal {
			maxAantal = count
			maxAantalPart = part
		}
	}
	fmt.Println(maxAantalPart, maxAantal)
}

func (p *Program) calculatePresents(remainingPresents, usedParts int) ([]rune, bool) {
	if remainingPresents == 0 {
		ok := usedParts == 0
		return []rune{}, ok
	}
	for _, part := range p.mainparts {
		if p.partsCount[part] <= usedParts {
			//fmt.Printf("Used %d parts for %s (%d parts for %d presents remains)\n", usedParts, part, usedParts-p.partsCount[part], remainingPresents-1)
			remainingPresents, ok := p.calculatePresents(remainingPresents-1, usedParts-p.partsCount[part])
			if ok {
				return append([]rune{rune(part[0])}, remainingPresents...), true
			}
		}
	}
	return []rune{}, false
}

func part2(p *Program) {
	sortedParts := p.mainparts
	sort.Slice(p.mainparts, func(i, j int) bool {
		return p.partsCount[p.mainparts[i]] > p.partsCount[p.mainparts[j]]
	})
	p.mainparts = sortedParts
	parts, ok := p.calculatePresents(20, p.missingParts)
	slices.Sort(parts)
	fmt.Println(string(parts), ok)
}

func solve(inputFile string) {
	fmt.Println(inputFile)
	ls, err := readFile(inputFile)
	if err != nil {
		panic(err)
	}

	p := parseInput(ls)
	part1(p)
	part2(p)
}

func main() {
	//solve("./in_test")
	solve("./in")
}
