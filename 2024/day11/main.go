package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"strings"
	"time"
)

func main() {
	lines, err := h.ReadLinesAsArray("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := part1(strings.Split(lines[0], " "), 25)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := part1(strings.Split(lines[0], " "), 75)
	fmt.Printf("Part 2: %d\n", p2)
}

func part1(input []string, blinks int) int {
	i := 0
	stones := input
	for i < blinks {
		fmt.Printf("Blink %d", i+1)
		now := time.Now()
		newStones := []string{}
		for _, stone := range stones {
			nextStones := handleBlink(stone)
			newStones = append(newStones, strings.Split(nextStones, " ")...)
		}
		stones = newStones
		i++
		fmt.Printf(" -- %vms\n", time.Now().Sub(now).Milliseconds())
	}

	return len(stones)
}

func part2(input []string, blinks int) int {
	i := 0

	lengths := []int{}
	for _, val := range input {
		// Treat "0" as having 0 length
		if val == "0" {
			lengths = append(lengths, 0)
		} else {
			lengths = append(lengths, len(val))
		}
	}

	for i < blinks {
		// fmt.Printf("Blink %d", i+1)
		// now := time.Now()

		newStoneLens := []int{}
		for _, sLen := range lengths {
			nextLens := handleBlinkUsingLen(sLen)
			newStoneLens = append(newStoneLens, nextLens...)
		}
		lengths = newStoneLens
		i++

		// fmt.Printf(" -- %vms\n", time.Now().Sub(now).Milliseconds())
	}

	return len(lengths)
}

type Stone struct {
	val string
	len int
}

func handleBlinkUsingLen(sLen int) []int {
	if sLen == 0 {
		return []int{1}
	}

	if sLen%2 == 0 {
		mid := sLen / 2

		return []int{mid, mid}
	}

	// Multiply sLen by 2024

	return []int{4 * sLen}
}

func handleBlink(stone string) string {
	if stone == "0" {
		return "1"
	} else if len(stone)%2 == 0 {
		mid := len(stone) / 2
		l := trimLeadingZeros(stone[:mid])
		r := trimLeadingZeros(stone[mid:])

		return fmt.Sprintf("%s %s", l, r)
	}

	// Multiply stone by 2024
	return multiply(stone, "2024")
}

func trimLeadingZeros(s string) string {
	trimmed := strings.TrimLeft(s, "0")
	if trimmed == "" {
		return "0"
	}

	return trimmed
}

func multiply(num1, num2 string) string {
	h.Assert(len(num1) != 0 && len(num2) != 0, fmt.Sprintf("multiply(%s,%s): Failed due to length being 0", num1, num2))
	// fmt.Println("Multiplying ", num1, num2)
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// Initialize result array to store intermediate results
	n1, n2 := len(num1), len(num2)
	result := make([]int, n1+n2)

	// Perform multiplication digit by digit
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + result[i+j+1]

			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}

	// Convert the result array back to a string
	var sb strings.Builder
	for _, val := range result {
		if !(sb.Len() == 0 && val == 0) { // Skip leading zeros
			sb.WriteByte(byte(val + '0'))
		}
	}

	return sb.String()
}
