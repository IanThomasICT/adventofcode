package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := h.ReadLinesAsArray("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := part1(lines)
	fmt.Printf("Part 1: %d\n", p1)

	p2 := part2(lines)
	fmt.Printf("Part 2: %d\n", p2)
	// 426214133343273 -- Too high

}

func part1(lines []string) int {
	calResult := 0
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Failed to parse '%s' to an int\n", parts[0])
			return -1
		}
		nums := h.SplitToIntSlice(parts[1], " ")

		if searchCombinations(nums[1:], nums[0], target) {
			// fmt.Printf("Success -> %s\n", l)
			calResult += target
		}
	}

	return calResult
}

func part2(lines []string) int {
	calResult := 0
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Failed to parse '%s' to an int\n", parts[0])
			return -1
		}
		nums := h.SplitToIntSlice(parts[1], " ")

		if searchWithConcat(nums[1:], nums[0], target) {
			// fmt.Printf("Success -> %s\n", l)
			calResult += target
		}
	}

	return calResult
}

func searchCombinations(nums []int, current int, target int) bool {
	if len(nums) == 0 && current == target {
		return true
	}

	for i, val := range nums {
		if searchCombinations(nums[i+1:], current*val, target) {
			return true
		} else if searchCombinations(nums[i+1:], current+val, target) {
			return true
		}
	}

	return false
}

func searchWithConcat(nums []int, current int, target int) bool {
	if len(nums) == 0 && current == target {
		return true
	}

	for i, val := range nums {
		if searchWithConcat(nums[i+1:], current*val, target) {
			return true
		} else if searchWithConcat(nums[i+1:], current+val, target) {
			return true
		}

		concat, err := strconv.Atoi(fmt.Sprintf("%d%d", current, val))
		if err != nil {
			fmt.Printf("Failed to parse '%d%d' to an int:\n", current, val)
			os.Exit(1)
		}

		if searchWithConcat(nums[i+1:], concat, target) {
			return true
		}
	}

	return false
}
