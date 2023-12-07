package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strconv"
)

func main() {

	testLines, err := h.ReadLinesAsArray("test.txt")
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}
	fmt.Println("test p1:", sumOfAdjacentParts(testLines))
	fmt.Println("test p2:", sumOfGearRatios(testLines))

	fileName := "input.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("part 1:", sumOfAdjacentParts(lines))
	fmt.Println("part 2:", sumOfGearRatios(lines))

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

func sumOfAdjacentParts(rows []string) int {

	parts := map[string]int{} // key: ij, val: partNum
	for i, row := range rows {
		for j, char := range row {
			// Skip if char is not a symbol
			if char == '.' || h.IsDigit(byte(char)) {
				continue
			}

			// char is a symbol
			coords := GetSurroundingCoords(i, j, len(rows), len(row))

			for _, coord := range coords {
				y, x := coord[0], coord[1]
				// If a surrounding coordinate contains a digit,
				// find the full number and add to map by number's first idx
				if h.IsDigit(rows[y][x]) {
					part, pIdx := revealNumberInString(x, rows[y])
					partLoc := fmt.Sprintf("%d%d-%d", y, pIdx, part)
					parts[partLoc] = part
				}
			}
		}
	}

	sum := 0
	for _, p := range parts {
		sum += p
	}
	return sum

}

// Part 2
func sumOfGearRatios(rows []string) int {
	GEAR := '*'

	ratioSum := 0
	for i, row := range rows {
		for j, char := range row {
			if char != GEAR {
				continue
			}
			// char is a symbol
			coords := GetSurroundingCoords(i, j, len(rows), len(row))
			adjParts := map[string]int{}

			for _, coord := range coords {
				y, x := coord[0], coord[1]
				// If a surrounding coordinate contains a digit,
				// find the full number and add to map by number's first idx
				if h.IsDigit(rows[y][x]) {
					part, pIdx := revealNumberInString(x, rows[y])
					partLoc := fmt.Sprintf("%d%d-%d", y, pIdx, part)
					adjParts[partLoc] = part
				}
			}

			if len(adjParts) == 2 {
				ratio := 1
				for _, v := range adjParts {
					ratio *= v
				}

				ratioSum += ratio
			}
		}
	}

	return ratioSum

}
func revealNumberInString(startIdx int, line string) (int, int) {
	if !h.IsDigit(line[startIdx]) {
		return 0, startIdx
	}

	l, r := startIdx, startIdx
	// Scan to left edge of number
	for l > 0 {
		if !h.IsDigit(line[l-1]) {
			break
		}
		l--
	}

	for r < len(line)-1 {
		if !h.IsDigit(line[r+1]) {
			break
		}
		r++
	}

	fullNumber, err := strconv.Atoi(line[l : r+1])
	if err != nil {
		log.Fatalf("Failed to convert %s to an int", line[l:r+1])
	}
	return fullNumber, l
}

func GetSurroundingCoords(i int, j int, arrHeight int, arrWidth int) [][]int {
	dirs := [][]int{
		{0, -1},  // left
		{-1, -1}, // top left
		{-1, 0},  // top
		{-1, 1},  // top right
		{0, 1},   // right
		{1, 1},   // bottom right
		{1, 0},   // bottom
		{1, -1},  // bottom left
	}

	coords := [][]int{}
	for _, dir := range dirs {
		y, x := i+dir[0], j+dir[1]
		// Add coordinate if it's within bounds of array
		if y >= 0 && y < arrHeight && x >= 0 && x < arrWidth {
			coords = append(coords, []int{y, x})
		}
	}

	return coords
}
