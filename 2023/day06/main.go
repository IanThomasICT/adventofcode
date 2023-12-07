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
	var product int64
	var limit int64 = int64(record)

	// Find general left edge
	for ; i < h; i += step {
		product = int64(l+i) * int64(h-i)
		if product > limit {
			break
		}
	}
	i -= step

	// Get leading fine edge
	product = int64(l+i) * int64(h-i)
	for product <= limit {
		product = int64(l+i) * int64(h-i)
		i++
	}
	fmt.Printf("leading edge at Lower: %d, Higher: %d, Record: %d => %d\n", l+i, h-i, limit, (l+i)*(h-i))

	aboveLimit := 1
	for ; i < h; i += step {
		product = int64(l+i) * int64(h-i)
		if product < limit {
			break
		} else {
			aboveLimit += step
		}
	}
	i -= step
	aboveLimit -= step

	product = int64(l+i) * int64(h-i)
	for product > limit {
		product = int64(l+i) * int64(h-i)
		i++
		aboveLimit++
	}
	aboveLimit--
	fmt.Println("aboveLimit", aboveLimit)

	return aboveLimit

}

func concatAsString(arr []int) string {
	str := ""
	for _, num := range arr {
		str += strconv.FormatInt(int64(num), 10)
	}
	return str
}
