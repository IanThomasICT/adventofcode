package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
)

func main() {

	// fileName := "input.txt"
	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("part 1:", part1(lines))
	// fmt.Println("part 2:", part1(lines))

}

func part1(lines []string) int {
	// function info
	return 0
}
