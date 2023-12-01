package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func ReadLinesAsArray(fileName string) ([]string, error) {
	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var arr []string
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}
	return arr, nil
}

func GetSumOfTopNStars(lines []string, n int) int64 {
	var currVal int64
	var sums []int64
	for _, l := range lines {
		if l == "" {
			sums = append(sums, currVal)
			currVal = 0
		} else {
			val, err := strconv.ParseInt(l, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse %s into int64: %s", l, err)
			}

			currVal += val
		}
	}
	if currVal > 0 {
		sums = append(sums, currVal)
	}
	slices.Sort(sums)

	// Get summed totals for top N values where N <= number of elves
	var total int64
	for i := 1; i <= min(n, len(sums)); i++ {
		total += sums[len(sums)-i]
	}
	return total
}

func SliceToOccurrencesMap(arr []byte) map[byte]int {
	arrMap := make(map[byte]int)
	for _, s := range arr {
		arrMap[s]++
	}
	return arrMap
}

func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
}
