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
type GhostNode struct {
	startKey string
	key      string
	stepsToZ []int
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
func part2(lines []string) int {
	// function info
	fmt.Printf("\033[H\033[2J")
	instructions := lines[0]

	moveMap := make(map[string][2]string)
	trackedNodes := []GhostNode{}

	for _, line := range lines[2:] {
		parts := strings.Split(line, " = ")
		key := parts[0]
		// fmt.Println(key)
		left, right := getDirections(parts[1])
		moveMap[key] = [2]string{left, right}
		if key[2] == 'A' {
			trackedNodes = append(trackedNodes, GhostNode{key, key, []int{}}) // start tracked nodes at all ending in A
		}
	}

	i := 0
	for !allValuesEndInZ(trackedNodes) {
		for n := range trackedNodes {
			current := &trackedNodes[n]
			inst := instructions[i%(len(instructions)-1)]
			if inst == 'L' {
				current.key = moveMap[current.key][0]
			} else if inst == 'R' {
				current.key = moveMap[current.key][1]
			}

			if current.key[2] == 'Z' {
				if len(current.stepsToZ) == 0 {
					current.stepsToZ = append(current.stepsToZ, i+1)
				} else {
					current.stepsToZ = append(current.stepsToZ, i+1-current.stepsToZ[len(current.stepsToZ)-1])
				}
				fmt.Printf("\033[H\033[2J")
				for _, n := range trackedNodes {
					fmt.Println(n.startKey, n.key, n.stepsToZ[max(0, len(n.stepsToZ)-3):])
				}
			}
		}
		i++
	}

	return i
}

func printNodes(nodes []Node) {
	fmt.Printf("\033[0;0H")
	str := "["
	for _, n := range nodes {
		if n.key[2] == 'Z' {
			str += fmt.Sprintf(" \033[33;1m%s\033[0m", n.key)
		} else {
			str += fmt.Sprintf(" %s", n.key)
		}

	}
	str += "]"
	fmt.Println(str)

}

func getDirections(str string) (string, string) {
	d := strings.Split(str[:len(str)-1], "(")[1] // Trim ( and )
	dirs := strings.Split(d, ", ")
	// fmt.Println(dirs[0], dirs[1])
	return dirs[0], dirs[1]
}

func allValuesEndInZ(nodes []GhostNode) bool {
	for i, n := range nodes {
		if n.key[2] != 'Z' {
			if i > len(nodes)/2 {
				fmt.Println("Halfway there..", nodes)
			} else if i > len(nodes)-2 {
				fmt.Println("SO CLOSE..", nodes)
			}
			return false
		}
	}
	return true
}
