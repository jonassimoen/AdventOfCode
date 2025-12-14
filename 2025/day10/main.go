package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

type Machine struct {
	indicators []bool
	buttons    [][]int
	joltages   []int
}

func parseInput(ls []string) []*Machine {
	var machines []*Machine

	// Updated regex to capture the different parts
	regex := regexp.MustCompile(`^\[([\.#]+)\]\s+(.+?)\s+\{([0-9,]+)\}$`)

	for _, l := range ls {
		matches := regex.FindStringSubmatch(l)
		if matches == nil {
			continue
		}

		machine := &Machine{}

		// Parse indicators: [.##.]
		indicatorStr := matches[1]
		machine.indicators = make([]bool, len(indicatorStr))
		for i, ch := range indicatorStr {
			machine.indicators[i] = (ch == '#')
		}

		// Parse buttons: (3) (1,3) (2) etc.
		buttonsStr := matches[2]
		buttonRegex := regexp.MustCompile(`\(([0-9,]+)\)`)
		buttonMatches := buttonRegex.FindAllStringSubmatch(buttonsStr, -1)

		machine.buttons = make([][]int, len(buttonMatches))
		for i, buttonMatch := range buttonMatches {
			numStrs := strings.Split(buttonMatch[1], ",")
			machine.buttons[i] = make([]int, len(numStrs))
			for j, numStr := range numStrs {
				num, _ := strconv.Atoi(numStr)
				machine.buttons[i][j] = num
			}
		}
		joltageStr := matches[3]
		joltageStrs := strings.Split(joltageStr, ",")
		machine.joltages = make([]int, len(joltageStrs))
		for i, js := range joltageStrs {
			machine.joltages[i], _ = strconv.Atoi(js)
		}

		machines = append(machines, machine)
	}

	return machines
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func minButtonPresses(machine *Machine) int {
	n := len(machine.indicators)
	m := len(machine.buttons)

	// Build augmented matrix [A|b]
	// A[i][j] = 1 if button j toggles light i
	// b[i] = 1 if light i needs to be on (target state from the pattern)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m+1)
		if machine.indicators[i] {
			matrix[i][m] = 1 // This light needs to be ON in target state
		}
	}

	for j, button := range machine.buttons {
		for _, light := range button {
			if light < n {
				matrix[light][j] = 1
			}
		}
	}

	// Gaussian elimination over GF(2)
	pivotCols := make([]int, 0)
	pivotRow := 0

	for col := 0; col < m && pivotRow < n; col++ {
		// Find pivot
		foundPivot := false
		for row := pivotRow; row < n; row++ {
			if matrix[row][col] == 1 {
				matrix[pivotRow], matrix[row] = matrix[row], matrix[pivotRow]
				foundPivot = true
				break
			}
		}

		if !foundPivot {
			continue
		}

		pivotCols = append(pivotCols, col)

		// Eliminate column
		for row := 0; row < n; row++ {
			if row != pivotRow && matrix[row][col] == 1 {
				for c := 0; c <= m; c++ {
					matrix[row][c] ^= matrix[pivotRow][c]
				}
			}
		}

		pivotRow++
	}

	// Check for inconsistency
	for row := pivotRow; row < n; row++ {
		if matrix[row][m] == 1 {
			return -1 // No solution
		}
	}

	// Identify free variables
	isPivot := make(map[int]bool)
	for _, col := range pivotCols {
		isPivot[col] = true
	}

	freeVars := make([]int, 0)
	for col := 0; col < m; col++ {
		if !isPivot[col] {
			freeVars = append(freeVars, col)
		}
	}

	// Try all combinations of free variables to find minimum
	minPresses := math.MaxInt32
	numFree := len(freeVars)

	for mask := 0; mask < (1 << numFree); mask++ {
		solution := make([]int, m)

		// Set free variables
		for i, freeVar := range freeVars {
			if (mask & (1 << i)) != 0 {
				solution[freeVar] = 1
			}
		}

		// Back-substitute to find pivot variables
		for i := len(pivotCols) - 1; i >= 0; i-- {
			col := pivotCols[i]
			row := i

			val := matrix[row][m]
			for c := col + 1; c < m; c++ {
				if matrix[row][c] == 1 {
					val ^= solution[c]
				}
			}
			solution[col] = val
		}

		// Count total presses
		presses := 0
		for _, v := range solution {
			presses += v
		}

		if presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses
}

func part1(ls []string) int {
	machines := parseInput(ls)
	total := 0

	for _, machine := range machines {
		presses := minButtonPresses(machine)
		total += presses
	}
	return total
}

func minJoltagePresses(machine *Machine) int {
	n := len(machine.joltages) // number of counters
	m := len(machine.buttons)  // number of buttons

	// Build matrix A where A[i][j] = 1 if button j affects counter i
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, m+1)
		matrix[i][m] = float64(machine.joltages[i])
	}

	for j, button := range machine.buttons {
		for _, counter := range button {
			if counter < n {
				matrix[counter][j] = 1
			}
		}
	}

	// Gaussian elimination
	pivotCols := make([]int, 0)
	pivotRow := 0

	for col := 0; col < m && pivotRow < n; col++ {
		// Find pivot
		maxRow := pivotRow
		maxVal := math.Abs(matrix[pivotRow][col])
		for row := pivotRow + 1; row < n; row++ {
			if math.Abs(matrix[row][col]) > maxVal {
				maxVal = math.Abs(matrix[row][col])
				maxRow = row
			}
		}

		if maxVal < 1e-10 {
			continue
		}

		matrix[pivotRow], matrix[maxRow] = matrix[maxRow], matrix[pivotRow]
		pivotCols = append(pivotCols, col)

		// Normalize pivot row
		pivot := matrix[pivotRow][col]
		for c := 0; c <= m; c++ {
			matrix[pivotRow][c] /= pivot
		}

		// Eliminate
		for row := 0; row < n; row++ {
			if row != pivotRow && math.Abs(matrix[row][col]) > 1e-10 {
				factor := matrix[row][col]
				for c := 0; c <= m; c++ {
					matrix[row][c] -= factor * matrix[pivotRow][c]
				}
			}
		}

		pivotRow++
	}

	// Check for inconsistency
	for row := pivotRow; row < n; row++ {
		if math.Abs(matrix[row][m]) > 1e-10 {
			return -1
		}
	}

	// Identify free variables
	isPivot := make(map[int]bool)
	for _, col := range pivotCols {
		isPivot[col] = true
	}

	freeVars := make([]int, 0)
	for col := 0; col < m; col++ {
		if !isPivot[col] {
			freeVars = append(freeVars, col)
		}
	}

	// Calculate upper bound for free variables based on max joltage value
	maxJoltage := 0
	for _, j := range machine.joltages {
		if j > maxJoltage {
			maxJoltage = j
		}
	}
	maxFreeVal := maxJoltage * 2 // More generous upper bound

	if len(freeVars) > 4 {
		maxFreeVal = maxJoltage // Reduce for many free vars
	}

	minTotal := math.MaxInt32

	var tryFreeVars func(int, []int)
	tryFreeVars = func(idx int, freeVals []int) {
		if idx == len(freeVars) {
			solution := make([]int, m)
			for i, v := range freeVals {
				solution[freeVars[i]] = v
			}

			// Back-substitute for pivot variables
			valid := true
			for i := len(pivotCols) - 1; i >= 0; i-- {
				col := pivotCols[i]
				row := i

				val := matrix[row][m]
				for c := col + 1; c < m; c++ {
					val -= matrix[row][c] * float64(solution[c])
				}

				intVal := int(math.Round(val))
				if intVal < 0 || math.Abs(val-float64(intVal)) > 0.01 {
					valid = false
					break
				}
				solution[col] = intVal
			}

			if !valid {
				return
			}

			// Verify solution
			for i := 0; i < n; i++ {
				sum := 0
				for j := 0; j < m; j++ {
					for _, counter := range machine.buttons[j] {
						if counter == i {
							sum += solution[j]
						}
					}
				}
				if sum != machine.joltages[i] {
					return
				}
			}

			// Count total
			total := 0
			for _, v := range solution {
				total += v
			}
			if total < minTotal {
				minTotal = total
			}
			return
		}

		// Prune: if current partial solution already exceeds min, skip
		currentSum := 0
		for i := 0; i < idx; i++ {
			currentSum += freeVals[i]
		}
		if currentSum >= minTotal {
			return
		}

		for v := 0; v <= maxFreeVal; v++ {
			tryFreeVars(idx+1, append(freeVals, v))
		}
	}

	tryFreeVars(0, []int{})

	if minTotal == math.MaxInt32 {
		return -1
	}
	return minTotal
}

func part2(ls []string) int {
	machines := parseInput(ls)
	total := 0

	for _, machine := range machines {
		presses := minJoltagePresses(machine)
		if presses > 0 {
			total += presses
		}
	}
	return total
}
