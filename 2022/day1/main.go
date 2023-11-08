package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fileName := "input.txt"
	lines, err := GetLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	part1 := getSumOfTopNStars(lines, 1)

	fmt.Printf("max stars: %d", part1)

}

func getSumOfTopNStars(lines []string, n int) int64 {
	//TODO: Implement n functionality
	var currVal, maxVal int64
	for _, l := range lines {
		switch l {
		case "":
			currVal = 0
		default:
			val, err := strconv.ParseInt(l, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse %s into int64: %s", l, err)
			}

			currVal += val
			if currVal > maxVal {
				maxVal = currVal
			}
		}
	}
	return maxVal
}

func GetLinesAsArray(fileName string) ([]string, error) {
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
