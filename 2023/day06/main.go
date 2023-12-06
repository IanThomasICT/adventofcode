package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"math"
	"strconv"
)

func main() {

	// fileName := "input"
	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))

}

func part1(lines []string) int {
	// function info
	times := getNumbers(lines[0])
	distances := getNumbers(lines[1])

	totalAboveRecord := 1
	for i := 0; i < len(times); i++ {
		totalTime := times[i]
		recordDistance := distances[i]

		totalAboveRecord *= getNumOfProductPairsAboveValue(0, totalTime, recordDistance)
	}

	fmt.Println(times, distances)

	return totalAboveRecord
}

func part2(lines []string) int {
	// function info
	time := concatAsString(getNumbers(lines[0]))
	distance := concatAsString(getNumbers(lines[1]))

	// getNumOfProductPairsAboveValue(0, totalTime, recordDistance)

	fmt.Println(time, distance)

	return 0
}
func getNumbers(line string) []int {
	numbers := []int{}
	// Get times
	num := ""
	for i := range line {
		c := line[i]
		if !h.IsDigit(c) && num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Failed to parse %s into int", num)
			}
			num = ""
			numbers = append(numbers, n)
		}

		if h.IsDigit(c) {
			num += string(c)
		}
	}
	if num != "" {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Failed to parse %s into int", num)
		}
		num = ""
		numbers = append(numbers, n)
	}
	return numbers
}

func getNumOfProductPairsAboveValue(l int, h int, limit int) int {

	step := int(math.Sqrt(float64(h)))

	i := step
	fmt.Println("Step", step)
	// Find general left edge
	for ; i < h; i += step {
		if (l+i)*(h-i) > limit {
			break
		}
	}
	i -= step

	// Get leading fine edge
	for (l+i)*(h-i) < limit {
		i++
	}
	fmt.Println("leading edge at", i, (l+i)*(h-i))

	aboveLimit := 0
	for ; i < h; i += step {
		if (l+i)*(h-i) < limit {
			break
		} else {
			aboveLimit += step
		}
	}
	i -= step
	aboveLimit -= step

	for (l+i)*(h-i) > limit {
		i++
		aboveLimit++
	}
	fmt.Println("limit", aboveLimit)

	return aboveLimit

}

func concatAsString(arr []int) string {
	str := ""
	for _, num := range arr {
		str += strconv.FormatInt(int64(num), 10)
	}
	return str
}
