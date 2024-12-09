package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"slices"
)

func getUpdatesByType(orderRules map[int][]int, updateOps [][]int) [2][][]int {
	goodUpdates := [][]int{}
	badUpdates := [][]int{}

ValidateUpdates:
	for _, updates := range updateOps {
		for i, page := range updates {
			mustComeAfter, ok := orderRules[page]
			if !ok {
				continue
			}

			if i <= len(updates)-1 {
				for _, val := range mustComeAfter {
					// Iterate through values that page MUST come AFTER
					// and see if those value appear in the remaining updates
					if slices.Contains(updates[i+1:], val) {
						badUpdates = append(badUpdates, updates)
						continue ValidateUpdates
					}
				}
			}
		}
		goodUpdates = append(goodUpdates, updates)
	}

	return [2][][]int{goodUpdates, badUpdates}
}

func part1(lines []string) int {
	// map[x] -> [intsXMustComeAfter]
	orderRules := map[int][]int{}
	updateOps := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		switch line[2] {
		case '|':
			rules := h.SplitToIntSlice(line, "|")
			h.Assert(len(rules) == 2, fmt.Sprintf("Ordering rules must only have two rules: %v", rules))
			orderRules[rules[1]] = append(orderRules[rules[1]], rules[0])
		case ',':
			updates := h.SplitToIntSlice(line, ",")
			updateOps = append(updateOps, updates)
		}
		// fmt.Println(line)
	}

	res := getUpdatesByType(orderRules, updateOps)
	goodUpdates, _ := res[0], res[1]

	fmt.Printf("Good updates (%d):\n", len(goodUpdates))
	for _, l := range goodUpdates {
		fmt.Println(l)
	}

	midpoints := []int{}
	for _, vals := range goodUpdates {
		m := len(vals) / 2
		midpoints = append(midpoints, vals[m])
	}
	fmt.Printf("Midpoints: %v\n", midpoints)

	return h.Sum(midpoints)
}

func reorder(orderRules map[int][]int, updates []int) []int {
	reordered := updates
	for i := 0; i < len(updates); i++ {
		val := updates[i]
		mustComeAfter, ok := orderRules[val]
		if !ok {
			continue
		}

		// Find last element in updates that val should come after
		lastFoundIdx := -1
		for _, v := range mustComeAfter {
			if i < len(updates)+1 && slices.Index(updates[i+1:], v) != -1 {
				valIdx := lastIndex(updates, v)
				if valIdx > i && valIdx > lastFoundIdx {
					lastFoundIdx = valIdx
				}
			}
		}

		if lastFoundIdx != -1 {
			// Move val after last element it must come after
			arr := moveAfter(updates, i, lastFoundIdx)
			// fmt.Printf("\tMoved %d after %d in %v -> %v\n", val, lastFoundIdx, updates, arr)
			return reorder(orderRules, arr)
		}
	}
	return reordered
}

func lastIndex(slice []int, val int) int {
	lastIdx := -1
	for i, v := range slice {
		if v == val {
			lastIdx = i
		}
	}
	return lastIdx
}

func moveAfter(slice []int, fromIdx int, afterIdx int) []int {
	arr := []int{}
	val := slice[fromIdx]
	for i := range slice {
		if i == fromIdx {
			continue
		}

		arr = append(arr, slice[i])
		if i == afterIdx {
			arr = append(arr, val)
		}
	}
	return arr
}

func part2(lines []string) int {
	// map[x] -> [intsXMustComeAfter]
	orderRules := map[int][]int{}
	updateOps := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		switch line[2] {
		case '|':
			rules := h.SplitToIntSlice(line, "|")
			h.Assert(len(rules) == 2, fmt.Sprintf("Ordering rules must only have two rules: %v", rules))
			orderRules[rules[1]] = append(orderRules[rules[1]], rules[0])
		case ',':
			updates := h.SplitToIntSlice(line, ",")
			updateOps = append(updateOps, updates)
		}
		// fmt.Println(line)
	}

	res := getUpdatesByType(orderRules, updateOps)
	badUpdates := res[1]

	fmt.Printf("Bad updates (%d):\n", len(badUpdates))
	for i, l := range badUpdates {
		fmt.Printf("Reordering update...\t%v", l)
		reordered := reorder(orderRules, badUpdates[i])
		fmt.Printf("-> %v\n", reordered)
		badUpdates[i] = reordered
	}

	midpoints := []int{}
	for _, vals := range badUpdates {
		m := len(vals) / 2
		midpoints = append(midpoints, vals[m])
	}
	fmt.Printf("Midpoints: %v\n", midpoints)

	return h.Sum(midpoints)
}

func main() {
	lines, err := h.ReadLinesAsArray("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	p1 := part1(lines)
	fmt.Printf("Part 1: %d\n\n", p1)

	p2 := part2(lines)
	fmt.Printf("Part 2: %d\n\n", p2)

}

// 4613 is too low
// 5091 -- CORRECT
