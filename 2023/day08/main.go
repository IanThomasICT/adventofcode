package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strings"
)

type Node struct {
	key   string
	left  string
	right string
}

func main() {

	runTests()

	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	// fmt.Println(lines[0])
	fmt.Println("part 1:", part1(lines))
	// fmt.Println("part 2:", part2(lines))

}

func runTests() {

	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("test 1:", part1(lines))
	// fmt.Println("test 2:", part2(lines))
}

func part1(lines []string) int {
	// function info
	instructions := lines[0]

	moveMap := make(map[string]Node)

	for _, line := range lines[2:] {
		parts := strings.Split(line, " = ")
		key := parts[0]
		// fmt.Println(key)
		left, right := getDirections(parts[1])
		moveMap[key] = Node{key, left, right}
	}

	current := moveMap["AAA"]
	i := 0
	for current.key != "ZZZ" {
		inst := instructions[i%len(instructions)]
		if inst == 'L' {
			current = moveMap[current.left]
		} else if inst == 'R' {
			current = moveMap[current.right]
		}
		i++
	}

	return i
}

func getDirections(str string) (string, string) {
	d := strings.Split(str[:len(str)-1], "(")[1] // Trim ( and )
	dirs := strings.Split(d, ", ")
	// fmt.Println(dirs[0], dirs[1])
	return dirs[0], dirs[1]
}
