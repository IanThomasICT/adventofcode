package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
)

func main() {

	runTests()

	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println(lines[0])
	// fmt.Println("part 1:", part1(lines))
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

type Node struct {
	val               byte
	i                 int
	j                 int
	distanceFromStart int
	prev              *Node
	next              []Node
}

func (n *Node) move(next Node) {
	next.prev = n
	next.distanceFromStart = n.distanceFromStart + 1
}

func isValidSymbol(s byte) bool {
	symbols := []byte{'|', '-', 'L', 'J', '7', 'F'}
	valid := false
	for _, symbol := range symbols {
		if s == symbol {
			valid = true
		}
	}
	return valid
}

// func (n *Node) getPossibleMoves(grid [][]byte) {
// 	switch n.val {
// 	case '|':
// 		return [][]Node{{-1, 0}, {1, 0}}
// 	case '-':
// 		return [][]Node{{0, -1}, {0, 1}}
// 	case 'L':
// 		return [][]Node{{-1, 0}, {0, 1}}
// 	case 'J':
// 		return [][]Node{{-1, 0}, {0, -1}}
// 	case '7':
// 		return [][]Node{{1, 0}, {0, -1}}
// 	case 'F':
// 		return [][]Node{{1, 0}, {0, 1}}
// 	}
// 	return [][]int{{0, 0}}
// }

func getMaxDistanceFromPoint(n Node, grid [][]byte) {
}

func part1(lines []string) int {
	// function info
	grid := [][]byte{}

	return 0
}
