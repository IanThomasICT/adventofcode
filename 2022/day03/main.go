package main

import (
	"cmp"
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"slices"
)

func main() {
	lines, err := h.ReadLinesAsArray("input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	p1Total := getSumOfPrioritiesInCompartments(lines)
	p2Total := getSumOfPrioritiesAcross3Rucksacks(lines)

	fmt.Println("Summed priorities:", p1Total)
	fmt.Println("Summed priorities across 3 rucksacks:", p2Total)
}

func getSumOfPrioritiesInCompartments(lines []string) int {
	total := 0
	for _, l := range lines {
		comp1 := l[:len(l)/2]
		comp2 := l[len(l)/2:]

		compartmentOneLetters := h.SliceToOccurrencesMap([]byte(comp1))

		sharedLetters := make(map[byte]int)
		for _, letter := range comp2 {
			char := byte(letter)
			if _, exists := compartmentOneLetters[char]; exists {
				sharedLetters[char]++
			}
		}
		//fmt.Printf("%s %s [ ", comp1, comp2)
		for char := range sharedLetters {
			total += getPriority(char)
			//fmt.Printf("%d ", getPriority(char))
		}
		//fmt.Println("]")
	}
	return total
}

func getSumOfPrioritiesAcross3Rucksacks(lines []string) int {

	const GROUP_SIZE = 3
	total := 0

	for i := 0; i < len(lines)-(GROUP_SIZE-1); i += GROUP_SIZE {
		group := lines[i:(i + GROUP_SIZE)]

		var lineLetters []map[byte]int
		for _, rs := range group {
			lineLetters = append(lineLetters, h.SliceToOccurrencesMap([]byte(rs)))
		}

		// Sort lines by number of unique letters
		slices.SortFunc(lineLetters, func(m1, m2 map[byte]int) int {
			return cmp.Compare(len(m1), len(m2))
		})

		// Set sharedLetters as largest map of unique letters
		sharedLetters := lineLetters[0]

		// Remove letters from sharedLetters that don't exist in other lines
		for char := range lineLetters[0] {
			for _, m := range lineLetters[1:] {
				if _, exists := m[char]; !exists '-'{
					delete(sharedLetters, char)
				}
			}
		}

		//fmt.Printf("%s %s [ ", group, group2)
		for char := range sharedLetters {
			total += getPriority(char)
			//fmt.Printf("%d ", getPriority(char))
		}
		//fmt.Println("]")
	}
	return total
}

func getPriority(char byte) int {
	// Lowercase letter (priority: 1-26)
	if char > 96 && char < 123 {
		return int(char) - 96
	}

	// Uppercase letter (priority: 27-52)
	if char > 64 && char < 91 {
		return int(char) - 64 + 26
	}

	return 0
}
