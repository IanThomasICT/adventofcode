package main

import (
	"cmp"
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// fileName := "input.txt"
	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("part 1:", part1(lines))
	// fmt.Println("part 2:", part1(lines))

}

// seeds: 79 14 55 13
//
// seed-to-soil map:
// 50 98 2
// 52 50 48
//
// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

type RangeInstruction struct {
	Start    int64
	End      int64
	ChangeBy int
}

type Seed struct {
	originalVal int64
	val         int64
}

func part1(lines []string) int {
	// read seeds
	seedLine := strings.Split(lines[0], ": ")[1]
	seeds := []Seed{}
	for _, seed := range strings.Split(seedLine, " ") {
		val, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s to an int64", seed)
		}
		seeds = append(seeds, Seed{originalVal: val, val: val})
	}
	slices.SortFunc(seeds, func(a Seed, b Seed) int {
		return cmp.Compare(a.originalVal, b.originalVal)
	})
	fmt.Println(seeds)

	stages := [7][]RangeInstruction{}

	l, stage := 2, 0
	for l < len(lines) {
		l++
		if IsBlank(lines[l]) {
			continue
		}

		if strings.Contains(lines[l], "map") {
			l++

			instructions := []RangeInstruction{}
			// parse map until reaching a blank line or end of the file
			for l < len(lines) && !IsBlank(lines[l]) {
				inst := SplitToInt64Slice(lines[l], " ")
				to, from, rng := inst[0], inst[1], inst[2]
				instructions = append(instructions, RangeInstruction{Start: from, End: from + rng, ChangeBy: int(to - from)})
				l++
			}
			slices.SortFunc(instructions, func(a RangeInstruction, b RangeInstruction) int {
				return cmp.Compare(a.Start, b.Start)
			})
			// fmt.Println(instructions)

			stages[stage] = instructions
			stage++
		}
	}
	for _, stage := range stages {
		fmt.Println(stage)
	}

	for _, stage := range stages {
		for _, seed := range seeds {
			rangeIdx := slices.IndexFunc(stage, func(ri RangeInstruction) bool {
				if seed.val >= ri.Start && seed.val <= ri.End {
					return true
				}
				return false
			})
			if rangeIdx == -1 {
				log.Fatalf("Unable to place seed value %d within stage ranges %v", seed.val, stage)
			}

			ri := stage[rangeIdx]
			seed.val = seed.val + int64(ri.ChangeBy)
		}
	}

	fmt.Println("Seeds after convesion:", seeds)

	return 0
}

func IsBlank(s string) bool {
	blank := s == "" || s == " " || s == "\n"
	if blank {
		return blank
	}

	for _, b := range s {
		if b != ' ' {
			return false
		}
	}
	return false
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
