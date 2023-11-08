package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
)

func main() {

	fileName := "input.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	part1 := h.GetSumOfTopNStars(lines, 1)
	part2 := h.GetSumOfTopNStars(lines, 3)

	fmt.Println("max stars:", part1)
	fmt.Println("max stars for top3:", part2)

}
