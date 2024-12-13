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

// Reads a file and returns a slice of strings where each element is a line in the file
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

// Returns the sum of the top N values from the provided lines
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

// Converts a slice of bytes to a map with the count of each byte
func SliceToOccurrencesMap(arr []byte) map[byte]int {
	arrMap := make(map[byte]int)
	for _, s := range arr {
		arrMap[s]++
	}
	return arrMap
}

// Converts a slice of strings to a map with the count of each string
func SliceToMap(arr []string) map[string]int {
	arrMap := make(map[string]int)
	for _, s := range arr {
		arrMap[s]++
	}
	return arrMap
}

// Checks if a byte is a digit
func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
}

// Splits a string into substrings of length n
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

func Assert(condition bool, message string) {
	if !condition {
		panic(fmt.Sprintf("Failed Assertion: %s", message))
	}
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Splits a string by a delimiter and converts the substrings to a slice of ints
func SplitToIntSlice(s string, del string) []int {
	splitElems := strings.Split(s, del)
	nums := []int{}

	for _, str := range splitElems {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failed to parse '%s' to an int while splitting string %s", str, s)
		}
		nums = append(nums, val)
	}

	return nums
}

// Splits a string by a delimiter and converts the substrings to a slice of int64s
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

// Parses a string to an int
func ParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to parse %s to an int", s)
	}
	return val
}

// Parses a string to an int64
func ParseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %s to an int64", s)
	}
	return val
}

// Returns the coordinates of the surrounding cells in a 2D array
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

// Split array into arrays of length subLen
func SplitToSubArrays(s []string, subLen int) [][]string {
	sl := len(s)
	if sl > 25 {
		sl = 25
	}
	Assert(subLen > 0, fmt.Sprintf("SplitToSubArray(%v, %d): Sub array list must be greater than 0", s[:sl], subLen))

	if len(s) <= subLen {
		return [][]string{s}
	}

	var res [][]string
	for i := 0; i < len(s); i += subLen {
		end := i + subLen
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[i:end])
	}
	return res
}
