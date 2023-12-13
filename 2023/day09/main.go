package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"slices"
)

func main() {

	runTests()

	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println(lines[0])
	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))

}

func runTests() {

	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("test 1:", part1(lines))
	fmt.Println("test 2:", part2(lines))
}

type OasisReport struct {
	readings [][]int
}

func (report *OasisReport) calculateDifferences() {
	allZero, depth := false, 1
	valueList := report.readings[0]
	for !allZero {
		report.readings = append(report.readings, []int{})

		allZero = true
		for i := 0; i < len(valueList)-1; i++ {
			val := valueList[i+1] - valueList[i]
			if val != 0 {
				allZero = false
			}
			report.readings[depth] = append(report.readings[depth], val)
		}
		valueList = report.readings[depth]
		depth++
	}
}

func (report *OasisReport) Print() {
	for row := range report.readings {
		fmt.Println(report.readings[row])
	}
}

func (report *OasisReport) extrapolateForward() {
	diff := 0
	numOfRows := len(report.readings)
	for r := range report.readings {
		row := report.readings[numOfRows-1-r]
		lastElem := row[len(row)-1]

		// Extend each readings row by the diff of the row below it
		report.readings[numOfRows-r-1] = append(row, lastElem+diff)
		diff = lastElem + diff
	}
}
func (report *OasisReport) extrapolateBackward() {
	diff := 0
	numOfRows := len(report.readings)
	for r := range report.readings {
		row := report.readings[numOfRows-1-r]
		firstElem := row[0]

		// Extend each readings row by the diff of the row below it
		slices.Insert(report.readings[numOfRows-r-1], 0, firstElem-diff)
		diff = firstElem - diff
	}
}
func part1(lines []string) int {
	// function info

	// get differences
	reports := []OasisReport{}
	for _, line := range lines {
		nums := h.SplitToIntSlice(line, " ")
		l := OasisReport{[][]int{nums}}
		l.calculateDifferences()
		l.extrapolateForward()
		reports = append(reports, l)
	}

	sum := 0
	for _, report := range reports {
		row := report.readings[0]
		sum += row[len(row)-1]

		// layers[l].Print()
		// fmt.Println("")
	}

	return sum
}

func part2(lines []string) int {
	// function info

	// get differences
	reports := []OasisReport{}
	for _, line := range lines {
		nums := h.SplitToIntSlice(line, " ")
		l := OasisReport{[][]int{nums}}
		l.calculateDifferences()
		l.extrapolateForward()
		l.extrapolateBackward()
		reports = append(reports, l)
	}

	sum := 0
	for _, report := range reports {
		row := report.readings[0]
		sum += row[0]

		// report.Print()
		// fmt.Println("")
	}

	return sum
}
