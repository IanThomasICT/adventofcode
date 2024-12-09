package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"slices"
)

type Guard struct {
	pos [2]int
	dir string
}

// Returns if the guard can move to the next position
func (g *Guard) move(m [][]byte) bool {
	switch g.dir {
	case "up":
		if g.pos[0] == 0 {
			return false
		}

		nextPos := [2]int{g.pos[0] - 1, g.pos[1]}
		if m[nextPos[0]][nextPos[1]] == '#' {
			g.dir = "right"
		} else {
			g.pos = nextPos
		}
	case "down":
		if g.pos[0] == len(m)-1 {
			return false
		}

		nextPos := [2]int{g.pos[0] + 1, g.pos[1]}
		if m[nextPos[0]][nextPos[1]] == '#' {
			g.dir = "left"
		} else {
			g.pos = nextPos
		}
	case "left":
		if g.pos[1] == 0 {
			return false
		}

		nextPos := [2]int{g.pos[0], g.pos[1] - 1}
		if m[nextPos[0]][nextPos[1]] == '#' {
			g.dir = "up"
		} else {
			g.pos = nextPos
		}

	case "right":
		if g.pos[1] == len(m[0])-1 {
			return false
		}

		nextPos := [2]int{g.pos[0], g.pos[1] + 1}
		if m[nextPos[0]][nextPos[1]] == '#' {
			g.dir = "down"
		} else {
			g.pos = nextPos
		}
	}

	return true
}

func main() {
	lines, err := h.ReadLinesAsArray("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load 2D map and create Guard
	m := make([][]byte, len(lines))
	var guard Guard
	for i, l := range lines {
		for j := range l {
			m[i] = append(m[i], l[j])
			if l[j] == '^' {
				guard = Guard{pos: [2]int{i, j}, dir: "up"}
			}
		}
	}

	p1 := part1(guard, m)
	fmt.Printf("Guard visited %d distinct positions\n", p1)

}

func part1(guard Guard, m [][]byte) int {
	visited := append([][2]int{}, guard.pos)
	for guard.move(m) {
		if !slices.Contains(visited, guard.pos) {
			visited = append(visited, guard.pos)
		}
	}
	return len(visited)
}

func part2(guard Guard, m [][]byte) int {
	juncPoints := append([][2]int{}, guard.pos)
	lastDir := guard.dir
	for guard.move(m) {
		if guard.dir != lastDir {
			juncPoints = append(juncPoints, guard.pos)
			lastDir = guard.dir
		}
	}
	return len(visited)
}
