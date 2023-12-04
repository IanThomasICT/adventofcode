package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strings"
)

func main() {

	fileName := "input"
	// fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("part 1:", parseScratchCards(lines))
	fmt.Println("part 2:", totalDuplicatedScratchCards(lines))

}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

func parseScratchCards(lines []string) int {

	totalScore := 0
	for _, card := range lines {
		card := strings.Split(card, ":")[1]
		parts := strings.Split(card, " |")

		winningSide, rightSide := parts[0], parts[1]

		// get winning numbers
		wMap := map[string]bool{}
		numbers := []string{}
		i := 1
		for i < len(rightSide)+1 {
			if i < len(winningSide)+1 && i%3 == 0 {
				wMap[winningSide[i-2:i]] = true
			}
			if i%3 == 0 {
				numbers = append(numbers, rightSide[i-2:i])
			}
			i++
		}

		// fmt.Println(wMap, numbers)

		score := 0
		for _, num := range numbers {
			if _, ok := wMap[num]; ok {
				if score == 0 {
					score = 1
				} else {
					score += score
				}
			}
		}
		// fmt.Printf("Card %d -> %d points\n", n+1, score)
		totalScore += score
	}

	return totalScore
}
func totalDuplicatedScratchCards(lines []string) int {

	cardCopies := make([]int, len(lines))
	for n, card := range lines {
		cardCopies[n]++ // Count original copy

		// Scratch original and duplicates
		for c := 0; c < cardCopies[n]; c++ {
			card := strings.Split(card, ":")[1]
			parts := strings.Split(card, " |")

			winningSide, rightSide := parts[0], parts[1]

			// get winning numbers
			wMap := map[string]bool{}
			numbers := []string{}
			i := 1
			for i < len(rightSide)+1 {
				if i < len(winningSide)+1 && i%3 == 0 {
					wMap[winningSide[i-2:i]] = true
				}
				if i%3 == 0 {
					numbers = append(numbers, rightSide[i-2:i])
				}
				i++
			}

			// fmt.Println(wMap, numbers)

			wins := 0
			for _, num := range numbers {
				if _, ok := wMap[num]; ok {
					wins++
				}
			}

			// Create copies for next w cards
			for w := 1; w <= wins; w++ {
				cardCopies[n+w]++
			}
		}
	}

	// Calculate total copies
	totalCopies := 0
	for _, copies := range cardCopies {
		totalCopies += copies
	}

	return totalCopies
}
