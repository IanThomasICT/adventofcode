package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
)

func main() {

	runTests()

	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println(lines[0])
	// fmt.Println("part 1:", part1(lines))
	// fmt.Println("part 2:", part2(lines))

}

func runTests() {

	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("test 1:", part1(lines))
	// fmt.Println("test 2:", part2(lines))
}

func part1(lines []string) int {
	// function info
	return 0
}
