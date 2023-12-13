package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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

func SliceToMap(arr []string) map[string]int {
	arrMap := make(map[string]int)
	for _, s := range arr {
		arrMap[s]++
	}
	return arrMap
}
func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
}

func SplitEveryN(s string, n int) []string {
	if n <= 0 || n >= len(s) {
		return []string{s}
	}

	var arr []string
	str := ""
	for _, c := range s {
		str += string(c)
		if len(str) == n {
			arr = append(arr, str)
			str = ""
		}
	}
	if len(str) > 0 {
		arr = append(arr, str)

	}

	return arr
}

func SplitToIntSlice(s string, del string) []int {
	splitElems := strings.Split(s, del)
	nums := []int{}

	for _, str := range splitElems {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failed to parse %s to an int while splitting string %s", str, s)
		}
		nums = append(nums, val)
	}

	return nums
}

func SplitToInt64Slice(s string, del string) []int64 {
	splitElems := strings.Split(s, del)
	nums := []int64{}

	for _, str := range splitElems {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s to an int64 while splitting string %s", str, s)
		}
		nums = append(nums, val)
	}

	return nums
}

func ParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to parse %s to an int", s)
	}
	return val
}

func ParseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %s to an int64", s)
	}
	return val
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
