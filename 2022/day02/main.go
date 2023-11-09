package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
)

type Move string

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0

	// Rock-Paper-Scissor moves
	ROCK     Move = "ROCK"
	PAPER    Move = "PAPER"
	SCISSORS Move = "SCISSORS"
)

type Combination struct {
	beats Move
	loses Move
	score int
}

func main() {
	lines, err := h.ReadLinesAsArray("input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	combinations := map[Move]Combination{
		ROCK:     Combination{beats: SCISSORS, loses: PAPER, score: 1},
		PAPER:    Combination{beats: ROCK, loses: SCISSORS, score: 2},
		SCISSORS: Combination{beats: PAPER, loses: ROCK, score: 3},
	}

	roundOneScore := PlayRockPaperScissors(lines, combinations, 1)
	roundTwoScore := PlayRockPaperScissors(lines, combinations, 2)

	fmt.Println("Total roundOneScore:", roundOneScore)
	fmt.Println("Total roundTwoScore:", roundTwoScore)
}

func PlayRockPaperScissors(lines []string, combinations map[Move]Combination, roundNum int) int {
	score := 0
	for _, l := range lines {
		if len(l) != 3 {
			log.Fatalln("line \"%s\" has more than 3 chars", l)
		}

		lMove, _ := getMove(l[0])
		rMove, _ := getMove(l[2])
		right := combinations[rMove]

		if roundNum == 1 {
			if right.beats == lMove {
				score += right.score + WIN
			} else if right.loses == lMove {
				score += right.score + LOSE
			} else {
				score += right.score + DRAW
			}
		} else if roundNum == 2 {
			desiredOutcome := getDesiredOutcome(l[2])
			opponent := combinations[lMove]
			if desiredOutcome == WIN {
				score += combinations[opponent.loses].score + desiredOutcome
			} else if desiredOutcome == LOSE {
				score += combinations[opponent.beats].score + desiredOutcome
			} else if desiredOutcome == DRAW {
				score += opponent.score + desiredOutcome
			}

		}

	}
	return score
}

func getMove(char byte) (Move, error) {
	switch char {
	case 'A', 'X':
		return ROCK, nil
	case 'B', 'Y':
		return PAPER, nil
	case 'C', 'Z':
		return SCISSORS, nil
	default:
		return "", fmt.Errorf("An invalid char '%s' was provided", char)
	}
}

func getDesiredOutcome(char byte) int {
	switch char {
	case 'X':
		return LOSE
	case 'Y':
		return DRAW
	case 'Z':
		return WIN
	default:
		return 0
	}
}
