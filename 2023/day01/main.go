package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strings"
)

func main() {

	fileName := "input.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	sum := part1(lines)
	fmt.Println("part 1:", sum)
	sum = part2(lines)

	fmt.Println("part 2:", sum)

}

func part1(lines []string) int {
	sum := 0

	// Iterate from front and back of line until a digit is found
	for _, line := range lines {
		lVal, rVal := -1, -1
		l, r := 0, len(line)-1

		for (lVal == -1 && rVal == -1) || (l < len(line) && r >= 0) {
			if lVal == -1 && h.IsDigit(line[l]) {
				lVal = int(line[l] - 48)
			}
			if rVal == -1 && h.IsDigit(line[r]) {
				rVal = int(line[r] - 48)
			}

			r--
			l++
		}
		sum += (lVal*10 + rVal)
	}
	return sum
}

func part2(lines []string) int {

	spelledNumbers := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	sum := 0
	for _, line := range lines {
		numsByIndex := make([]int, len(line))

		//	Get indexes of any spelled numbers
		for n, val := range spelledNumbers {
			f := strings.Index(line, n)
			if f != -1 {
				numsByIndex[f] = val
			}

			l := strings.LastIndex(line, n)
			if l != -1 {
				numsByIndex[l] = val
			}
		}

		// Find indexes of digits
		for i := range line {
			if h.IsDigit(line[i]) {
				numsByIndex[i] = int(line[i] - 48)
			}
		}

		lVal, rVal := -1, -1
		l, r := 0, len(line)-1
		for (lVal == -1 && rVal == -1) || (l < len(numsByIndex) && r >= 0) {
			if lVal == -1 && numsByIndex[l] != 0 {
				lVal = numsByIndex[l]
			}
			if rVal == -1 && numsByIndex[r] != 0 {
				rVal = numsByIndex[r]
			}

			r--
			l++
		}

		sum += (lVal*10 + rVal)
		//fmt.Println(sum)
	}
	return sum
}
