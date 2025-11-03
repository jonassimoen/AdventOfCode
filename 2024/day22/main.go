package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
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

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func calculateSecretNumber(x int) int {
	res1 := prune(mix(x, x*64))
	res2 := prune(mix(res1, int(math.Floor(float64(res1)/float64(32)))))
	res3 := prune(mix(res2, res2*2048))
	return res3
}

func tests() {
	fmt.Printf("Test 1: mix 42/15: %d\n", mix(42, 15))
	fmt.Printf("Test 2: prune 100000000: %d\n", prune(100000000))
	secret := 123
	for i := range 10 {
		secret = calculateSecretNumber(secret)
		fmt.Printf("Test 3.%d: secret 123: %d\n", i, secret)
	}
}

func generateSecretNumbers(secret, n int) int {
	newSecret := secret
	for range n {
		newSecret = calculateSecretNumber(newSecret)
	}
	return newSecret
}

func part1(ls []string) int {
	//tests()
	total := 0
	for _, numberLine := range ls {
		number, err := strconv.Atoi(numberLine)
		if err != nil {
			panic(err)
		}
		secret := generateSecretNumbers(number, 2000)
		total += secret
		//fmt.Printf("%d: %d\n", number, secret)
	}
	return total
}

func toSequence(n []int) string {
	if len(n) != 4 {
		panic("Invalid sequence")
	}
	return fmt.Sprintf("%d%d%d%d", n[0], n[1], n[2], n[3])
}

func part2(ls []string) int {
	total := 0
	totalBananasPerSequence := map[string]int{}
	for _, numberLine := range ls {
		foundSeqForSecret := map[string]bool{}
		number, err := strconv.Atoi(numberLine)
		if err != nil {
			panic(err)
		}
		secret := number
		//fmt.Printf("%8d: %2d\n", secret, secret%10)

		sequence := []int{}
		for n := range 2000 {
			next := calculateSecretNumber(secret)
			diff := next%10 - secret%10
			sequence = append(sequence, diff)
			//fmt.Printf("%8d: %2d (%d)\n", next, next%10, diff)
			//fmt.Printf("\t> %v\n", sequence)
			secret = next

			if n >= 3 {
				seq := toSequence(sequence)
				if _, ok := foundSeqForSecret[seq]; !ok {
					if _, ok := totalBananasPerSequence[seq]; !ok {
						totalBananasPerSequence[seq] = 0
					}
					totalBananasPerSequence[seq] += secret % 10
					//fmt.Printf("\t> %s: %d\n", seq, totalBananasPerSequence[seq])
					foundSeqForSecret[seq] = true
				}
				sequence = sequence[1:]
			}

		}
		total += secret
	}
	maxValue := 0
	for _, v := range totalBananasPerSequence {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
