package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"strconv"
	"strings"
)

// [stoneVal] => totalNewStonesCreated
var lookup = map[string]map[int]int{}

func main() {
	lines, err := h.ReadLinesAsArray("sample")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := part1(strings.Split(lines[0], " "), 25)
	fmt.Printf("Part 1: %d\n", p1)

	// p2 := part1(strings.Split(lines[0], " "), 75)
	// fmt.Printf("Part 2: %d\n", p2)
}

func part1(input []string, blinks int) int {

	totalCount := len(input)
	for _, stone := range input {
		totalCount += processStone(stone, 0, 0, blinks)
	}

	fmt.Println(lookup)
	return totalCount

	// i := 0

	// for i < blinks {
	// 	fmt.Printf("Blink %d", i+1)
	// 	now := time.Now()

	// 	i++
	// 	fmt.Printf(" -- %vms (%d stones)\n", time.Now().Sub(now).Milliseconds(), len(s.stones))
	// }

	// return len(s.stones)
}

func handleBlink(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}

	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		l := trimLeadingZeros(stone[:mid])
		r := trimLeadingZeros(stone[mid:])

		return []string{l, r}
	}

	// Multiply stone by 2024
	si, _ := strconv.Atoi(stone)
	res := fmt.Sprintf("%d", si*2024)
	return []string{res}
}

func processStone(stone string, currBlink int, startBlink int, targetBlinks int) int {
	sMap, ok := lookup[stone]
	if !ok {
		sMap = map[int]int{}
	}

	// Check if stone was already processed
	newStoneCount := 0
	if s, ok := sMap[targetBlinks]; ok {
		return s
	}

	newStones := handleBlink(stone)
	if len(newStones) > 1 {
		newStoneCount++
	}

	if currBlink < targetBlinks {
		for _, s := range newStones {
			newStoneCount += processStone(s, currBlink+1, startBlink, targetBlinks)
		}
	}

	sMap[targetBlinks] = newStoneCount
	lookup[stone] = sMap
	// fmt.Printf("stone (%s) has %d new stones at %d depth\n", stone, newStoneCount, startBlink)
	return newStoneCount

}

func trimLeadingZeros(s string) string {
	trimmed := strings.TrimLeft(s, "0")
	if trimmed == "" {
		return "0"
	}

	return trimmed
}

// func multiply(num1, num2 string) string {
// 	h.Assert(len(num1) != 0 && len(num2) != 0, fmt.Sprintf("multiply(%s,%s): Failed due to length being 0", num1, num2))
// 	// fmt.Println("Multiplying ", num1, num2)
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}

// 	// Initialize result array to store intermediate results
// 	n1, n2 := len(num1), len(num2)
// 	result := make([]int, n1+n2)

// 	// Perform multiplication digit by digit
// 	for i := n1 - 1; i >= 0; i-- {
// 		for j := n2 - 1; j >= 0; j-- {
// 			mul := int(num1[i]-'0') * int(num2[j]-'0')
// 			sum := mul + result[i+j+1]

// 			result[i+j+1] = sum % 10
// 			result[i+j] += sum / 10
// 		}
// 	}

// 	// Convert the result array back to a string
// 	var sb strings.Builder
// 	for _, val := range result {
// 		if !(sb.Len() == 0 && val == 0) { // Skip leading zeros
// 			sb.WriteByte(byte(val + '0'))
// 		}
// 	}

// 	return sb.String()
// }
