package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// fileName := "test.txt"
	// lines, err := h.ReadLinesAsArray(fileName)
	// if err != nil {
	// 	log.Fatalf("failed to get lines as []string: %s", err)
	// }
	//
	// fmt.Println("test:", part2(lines))

	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}
	// fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))

}

type SeedRange struct {
	Start int64
	End   int64
}

type OffsetRange struct {
	Start  int64
	End    int64
	Offset int64
}

func part1(lines []string) int64 {
	// read seeds
	seedLine := strings.Split(lines[0], ": ")[1]
	seeds := []int64{}
	for _, seed := range strings.Split(seedLine, " ") {
		val := ParseInt64(seed)
		seeds = append(seeds, val)
	}

	fmt.Println("Seed count:", len(seeds))

	stages := getOffsetInstructions(lines)

	for i, stage := range stages {
		fmt.Println("Stage", i+1)
		for s := range seeds {
			switch s {
			case len(seeds) / 4:
				fmt.Println("Converted 25%")
			case len(seeds) / 2:
				fmt.Println("Converted 50%")
			case (len(seeds) / 4) * 3:
				fmt.Println("Converted 75%")
			default:
			}

			rangeIdx := slices.IndexFunc(stage, func(ri OffsetRange) bool {
				if seeds[s] >= ri.Start && seeds[s] <= ri.End {
					return true
				}
				return false
			})
			if rangeIdx == -1 {
				continue
			}

			ri := stage[rangeIdx]
			seeds[s] += ri.Offset
		}
	}

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i] < seeds[j]
	})

	var lowestLocation int64 = seeds[0]

	return lowestLocation
}

type Seed struct {
	original int64
	val      int64
}

func part2(lines []string) int64 {
	// read seeds
	seedLine := strings.Split(lines[0], ": ")[1]
	seedParts := strings.Split(seedLine, " ")

	seedRanges := [][]int64{}
	for s := 0; s < len(seedParts)-1; s += 2 {
		pair := seedParts[s : s+2]
		startVal, numOfSeeds := ParseInt64(pair[0]), ParseInt64(pair[1])
		seedRanges = append(seedRanges, []int64{startVal, startVal + numOfSeeds})
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i][0] < seedRanges[j][0]
	})

	// Where is this? 1212778760
	for f := range seedRanges {
		fmt.Println(seedRanges[f])
	}

	var lowestLocation int64 = 200_000_000_000

	stages := getOffsetInstructions(lines)

	for _, seedRange := range seedRanges {
		fmt.Printf("Checking from %d to %d\n", seedRange[0], seedRange[1])
		// iterate by 25
		for i := seedRange[0]; i < seedRange[1]; i++ {
			loc := convertSeedToLocation(Seed{i, i}, stages)

			// if location is low, iterate through last 25
			if loc < lowestLocation {
				fmt.Println("Found lower location:", loc)
				lowestLocation = min(loc, lowestLocation)
				// for j := min(i-1, seedRange[0]); j < i+1; j++ {
				// 	locJ := convertSeedToLocation(Seed{j, j}, stages)
				// 	if locJ < lowestLocation {
				// 		fmt.Println("Found even lower location!", locJ, "from seed:", j)
				// 	}
				// 	lowestLocation = min(locJ, lowestLocation)
				// }
			}
		}
	}

	// Wrong answers:
	/**
		1
	-22
	-7
	1
	2
	3
	78775052
	*/
	return lowestLocation
}

func convertSeedToLocation(seed Seed, stages [7][]OffsetRange) int64 {
	for _, stage := range stages {
		rangeIdx := slices.IndexFunc(stage, func(ri OffsetRange) bool {
			if seed.val >= ri.Start && seed.val <= ri.End {
				return true
			}
			return false
		})
		if rangeIdx == -1 {
			continue
		}

		ri := stage[rangeIdx]
		seed.val += ri.Offset
	}

	return seed.val
}

func getOffsetsForSeedRange(seedRng SeedRange, stage []OffsetRange) []OffsetRange {
	newRanges := []OffsetRange{}

	if (seedRng.Start < stage[0].Start && seedRng.End < stage[0].Start) || (seedRng.Start > stage[len(stage)-1].End && seedRng.End > stage[len(stage)-1].End) {
		return []OffsetRange{{Start: seedRng.Start, End: seedRng.End, Offset: 0}}
	}

	// For half before and half inside first offset range, offset by 0
	if seedRng.Start < stage[0].Start && seedRng.End > stage[0].Start {
		newRanges = append(newRanges, OffsetRange{Start: seedRng.Start, End: stage[0].Start - 1, Offset: 0})
	}

	for _, ri := range stage {
		// Start is fully inside offset range
		if seedRng.Start > ri.Start && seedRng.Start < ri.End {
			newRanges = append(newRanges, OffsetRange{Start: seedRng.Start, End: min(ri.End, seedRng.End), Offset: ri.Offset})
		} else if seedRng.End < ri.End {
			newRanges = append(newRanges, OffsetRange{Start: ri.Start, End: seedRng.End, Offset: ri.Offset})
		} else {
			newRanges = append(newRanges, ri)
		}
	}

	// For seed range after last stage, offset by 0
	if seedRng.End > stage[len(stage)-1].End && seedRng.Start < stage[len(stage)-1].End {
		newRanges = append(newRanges, OffsetRange{Start: stage[len(stage)-1].End + 1, End: seedRng.End, Offset: 0})
	}

	fmt.Println("Split range", seedRng, "into", newRanges)

	return newRanges
}

// region Helpers
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

func ParseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %s to an int64", s)
	}
	return val
}

func getOffsetInstructions(lines []string) [7][]OffsetRange {
	stages := [7][]OffsetRange{}

	stage := 0
	for l := 1; l < len(lines); l++ {
		if IsBlank(lines[l]) {
			continue
		}

		if strings.Contains(lines[l], "map") {
			l++

			instructions := []OffsetRange{}
			// parse map until reaching a blank line or end of the file
			for l < len(lines) && !IsBlank(lines[l]) {
				inst := SplitToInt64Slice(lines[l], " ")
				to, from, rng := inst[0], inst[1], inst[2]
				instructions = append(instructions, OffsetRange{Start: from, End: from + rng, Offset: to - from})
				l++
			}

			// Sort instructions by Start value
			sort.Slice(instructions, func(i, j int) bool {
				return instructions[i].Start < instructions[j].Start
			})

			stages[stage] = instructions
			stage++
		}
	}

	return stages
}

//endregion
