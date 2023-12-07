package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"math"
	"strconv"
)

func main() {

	fileName := "input"
	// fileName := "test.txt"
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

	timeStr := concatAsString(getNumbers(lines[0]))
	distanceStr := concatAsString(getNumbers(lines[1]))

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	fmt.Println(time, distance)

	// 71503
	return getNumOfProductPairsAboveValue(0, time, distance)
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

func getNumOfProductPairsAboveValue(l int, h int, record int) int {

	step := int(math.Sqrt(float64(h)))

	i := step
	var limit int64 = int64(record)

	// Find general left edge
	aboveLimit := 0
	for ; i < h; i++ {
		product, err := MultiplyExact(int64(l+i), int64(h-i))
		if err != nil {
			log.Fatalln(err)
		}
		if err != nil || product > limit {
			aboveLimit++
		}
	}
	return aboveLimit

}

func concatAsString(arr []int) string {
	str := ""
	for _, num := range arr {
		str += strconv.FormatInt(int64(num), 10)
	}
	return str
}

const mostNegative = -(mostPositive + 1)
const mostPositive = 1<<63 - 1

func MultiplyExact(a, b int64) (int64, error) {
	result := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return result, nil
	}
	if a == mostNegative || b == mostNegative {
		return result, fmt.Errorf("Overflow multiplying %v and %v", a, b)
	}
	if result/b != a {
		return result, fmt.Errorf("Overflow multiplying %v and %v", a, b)
	}
	return result, nil
}
