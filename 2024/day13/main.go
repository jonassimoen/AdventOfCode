package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

type ClawMachine struct {
	btnA  []int
	btnB  []int
	price []int
}

func parseInput(ls []string, add int) []ClawMachine {
	var clawMachines []ClawMachine
	for i := 0; i < len(ls)-1; i += 4 {
		xA, yA := parseXY(ls[i])
		xB, yB := parseXY(ls[i+1])
		xP, yP := parseXY(ls[i+2])
		cm := ClawMachine{
			btnA:  []int{xA, yA},
			btnB:  []int{xB, yB},
			price: []int{xP + add, yP + add},
		}
		clawMachines = append(clawMachines, cm)
	}
	return clawMachines
}

func parseXY(l string) (int, int) {
	s := r.FindAllString(l, -1)
	ss := strings.Split(s[0], ", ")
	sA, _ := strconv.Atoi(ss[0][2:])
	sB, _ := strconv.Atoi(ss[1][2:])
	return sA, sB
}

func calculateCheapestCombination(cm ClawMachine) (int, int) {
	pushesA := (cm.price[1]*cm.btnB[0] - cm.price[0]*cm.btnB[1]) / (cm.btnA[1]*cm.btnB[0] - cm.btnB[1]*cm.btnA[0])
	pushesB := (cm.price[1]*cm.btnA[0] - cm.price[0]*cm.btnA[1]) / (cm.btnB[1]*cm.btnA[0] - cm.btnA[1]*cm.btnB[0])
	return pushesA, pushesB
}

func checkClawMachines(cm []ClawMachine, maxPresses int) int {
	sum := 0
	for _, m := range cm {
		pA, pB := calculateCheapestCombination(m)
		tokens := 0
		valid := ((pA*m.btnA[0] + pB*m.btnB[0]) == m.price[0]) && ((pA*m.btnA[1] + pB*m.btnB[1]) == m.price[1])
		if valid {
			if maxPresses != -1 && (pA > 0 && pA < 100) && (pB > 0 && pB < 100) {
				tokens = pA*3 + pB
			} else if maxPresses == -1 {
				tokens = pA*3 + pB
			}
		}
		sum += tokens
	}
	return sum
}

func checkClawMachinesIterative(cm []ClawMachine) int {
	sum := 0
	for _, m := range cm {
		pAmin, pBmin, tokensMin := 0, 0, 0
		for pA := 0; pA < 100; pA++ {
			for pB := 0; pB < 100; pB++ {
				xx := pA*m.btnA[0] + pB*m.btnB[0]
				yy := pA*m.btnA[1] + pB*m.btnB[1]
				if xx == m.price[0] && yy == m.price[1] {
					if pAmin != 0 || pBmin != 0 {
						fmt.Println("Duplicate")
					}
					pAmin = pA
					pBmin = pB
					tokensMin = pA*3 + pB
				}
			}
		}
		sum += tokensMin
	}
	return sum
}

var r = regexp.MustCompile("X[+|=]([0-9]*), Y[+|=]([0-9]*)")

func part1(ls []string) int {
	c := parseInput(ls, 0)
	return checkClawMachines(c, 100)
}

func part2(ls []string) int {
	c := parseInput(ls, 10000000000000)
	return checkClawMachines(c, -1)
}
