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

func NewGuard(pos [2]int, dir string) Guard {
	return Guard{pos, dir}
}

// Returns if the guard can move to the next position
func (g *Guard) move(m [][]byte) bool {
	// Get next position
	nextPos := g.getNextPos(m)
	if nextPos == nil {
		return false
	}

	if m[nextPos[0]][nextPos[1]] == '#' {
		g.rotate(m)
	} else {
		g.pos = *nextPos
	}

	return true
}

func (g *Guard) getNextPos(m [][]byte) *[2]int {
	switch g.dir {
	case "up":
		if g.pos[0] == 0 {
			return nil
		}
		return &[2]int{g.pos[0] - 1, g.pos[1]}
	case "down":
		if g.pos[0] == len(m)-1 {
			return nil
		}
		return &[2]int{g.pos[0] + 1, g.pos[1]}
	case "left":
		if g.pos[1] == 0 {
			return nil
		}
		return &[2]int{g.pos[0], g.pos[1] - 1}
	case "right":
		if g.pos[1] == len(m[0])-1 {
			return nil
		}
		return &[2]int{g.pos[0], g.pos[1] + 1}
	}

	return nil
}

func (g *Guard) rotate(m [][]byte) {
	switch g.dir {
	case "up":
		g.dir = "right"
	case "down":
		g.dir = "left"
	case "left":
		g.dir = "up"
	case "right":
		g.dir = "down"
	}
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
	fmt.Printf("Part 1: Guard visited %d distinct positions\n", p1)

	p2 := part2(guard, m)
	fmt.Printf("Part 2: %d distinct junction positions\n", p2)

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

func dirNum(dir string) int {
	return slices.Index([]string{"up", "right", "down", "left"}, dir)
}

func part2(guard Guard, m [][]byte) int {
	startPos, startDir := guard.pos, guard.dir

	loopMakingPoints := 0
	for i, row := range m {
		for j := range row {
			oldVal := m[i][j]
			m[i][j] = '#'
			temp := NewGuard(startPos, startDir)

			stopPoints := map[string]bool{
				fmt.Sprint(temp.pos[0], temp.pos[1], dirNum(temp.dir)): true,
			}

			lastDir := temp.dir
			for temp.move(m) {
				if temp.dir != lastDir {
					stoppedOn := fmt.Sprint(temp.pos[0], temp.pos[1], dirNum(temp.dir))
					if stopPoints[stoppedOn] {
						// A loop was formed
						loopMakingPoints++
						break
					}

					stopPoints[stoppedOn] = true
					lastDir = temp.dir
				}
			}

			// Reset map
			m[i][j] = oldVal
		}
	}

	return loopMakingPoints
}
