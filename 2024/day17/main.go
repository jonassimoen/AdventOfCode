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

type Registers struct {
	a, b, c      int
	instructions []int
	instr_idx    int
	output       []int
}

func parseInput(ls []string) *Registers {
	regA, _ := strconv.Atoi(strings.Split(ls[0], ": ")[1])
	regB, _ := strconv.Atoi(strings.Split(ls[1], ": ")[1])
	regC, _ := strconv.Atoi(strings.Split(ls[2], ": ")[1])

	p := ls[4]
	p = p[9:]
	s := strings.Split(p, ",")
	i := make([]int, 0)
	for _, v := range s {
		f, _ := strconv.Atoi(v)
		i = append(i, f)
	}
	return &Registers{
		a:            regA,
		b:            regB,
		c:            regC,
		instructions: i,
		instr_idx:    0,
		output:       make([]int, 0),
	}
}

func (r *Registers) parseComboOperand(c int) int {
	switch c {
	case 4:
		return r.a
	case 5:
		return r.b
	case 6:
		return r.c
	case 7:
		//panic("Invalid program, contains combo operand 7")
		return 7
	default:
		return c
	}
}

func (r *Registers) parseOpcode() {
	opCode := r.instructions[r.instr_idx]
	operandCode := r.instructions[r.instr_idx+1]
	operand := r.parseComboOperand(operandCode)
	//fmt.Printf("Opcode: %d, operand: %d (%d)\n", opCode, operand, operandCode)
	switch opCode {
	case 0:
		r.adv(operand)
		r.instr_idx += 2
		return
	case 1:
		r.bxl(operandCode)
		r.instr_idx += 2
		return
	case 2:
		r.bst(operand)
		r.instr_idx += 2
		return
	case 3:
		r.jnz(operandCode)
		return
	case 4:
		r.bxc(operand)
		r.instr_idx += 2
		return
	case 5:
		r.out(operand)
		r.instr_idx += 2
		return
	case 6:
		r.bdv(operand)
		r.instr_idx += 2
		return
	case 7:
		r.cdv(operand)
		r.instr_idx += 2
		return
	}
}

func (r *Registers) runProg() {
	for r.instr_idx < len(r.instructions) {
		r.parseOpcode()
		//fmt.Println(r)
	}
}

// 0
func (r *Registers) adv(operand int) {
	r.a = r.dv(operand)
}

// 1
func (r *Registers) bxl(operand int) {
	r.b = r.b ^ operand
}

// 2
func (r *Registers) bst(operand int) {
	r.b = operand % 8
}

// 3
func (r *Registers) jnz(operand int) {
	if r.a != 0 {
		r.instr_idx = operand
	} else {
		r.instr_idx += 2
	}
}

// 4
func (r *Registers) bxc(operand int) {
	r.b = r.b ^ r.c
}

// 5
func (r *Registers) out(operand int) {
	operand = operand % 8
	r.output = append(r.output, operand)
}

// 6
func (r *Registers) bdv(operand int) {
	r.b = r.dv(operand)
}

// 7
func (r *Registers) cdv(operand int) {
	r.c = r.dv(operand)
}

func (r *Registers) dv(operand int) int {
	return int(float64(r.a) / math.Pow(2, float64(operand)))

}

func part1_tests() {
	regs := []Registers{
		{c: 9, instructions: []int{2, 6}},
		{a: 10, instructions: []int{5, 0, 5, 1, 5, 4}},
		{a: 2024, instructions: []int{0, 1, 5, 4, 3, 0}},
		{b: 29, instructions: []int{1, 7}},
		{b: 2024, c: 43690, instructions: []int{4, 0}},
	}

	for _, r := range regs {
		fmt.Println(r)
		r.runProg()
		fmt.Println(r)
		fmt.Println("=====")
	}
}

func part1(ls []string) int {
	//part1_tests()
	reg := parseInput(ls)
	//fmt.Println(reg)
	reg.runProg()
	//fmt.Println(reg)
	b := strings.Builder{}
	for _, o := range reg.output {
		b.WriteString(strconv.Itoa(o))
		b.WriteRune(',')
	}
	fmt.Println(b.String())

	return -1
}

func (r *Registers) isCopy() bool {
	if len(r.output) != len(r.instructions) {
		return false
	}
	for idx, o := range r.instructions {
		if o != r.output[idx] {
			return false
		}
	}
	return true
}

func part2(ls []string) int {
	// TODO
	return -1
}
