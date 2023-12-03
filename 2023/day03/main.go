package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strconv"
)

type Part struct {
	value    int
	location [2]int // row, j
	length   int
}

type Symbol struct {
	value    byte
	location [2]int
}

func main() {

	// fileName := "input.txt"
	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	part1(lines)
	// fmt.Println("part 1:", sum)
	// sum = part2(lines)
	// fmt.Println("part 2:", sum)

}

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

func part1(rows []string) {
	parts := []Part{}
	symbols := []Symbol{}

	for i, row := range rows {
		for j := 0; j < len(row)-1; j++ {
			char := row[j]
			if char == '.' {
				continue
			} else if h.IsDigit(char) {
				startIdx := j
				for h.IsDigit(row[j]) && j < len(row)-1 {
					j++
				}

				valStr := row[startIdx:j]
				val, err := strconv.Atoi(valStr)
				if err != nil {
					log.Fatalf("Failed to convert %s into an integer", valStr)
				}

				parts = append(parts, Part{value: val, location: [2]int{i, startIdx}, length: len(valStr)})
			} else {
				symbols = append(symbols, Symbol{value: char, location: [2]int{i, j}})
			}
		}
	}

	// fmt.Println(parts)
	// fmt.Println(symbols)

}
