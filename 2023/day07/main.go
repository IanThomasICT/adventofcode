package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"sort"
	"strings"
)

func main() {

	runTest()
	fileName := "input"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	// fmt.Println(lines[0])
	fmt.Println("part 1:", playPoker(lines))
	fmt.Println("part 2:", playPokerWithJokers(lines))

}

func runTest() {

	fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	fmt.Println("test 1:", playPoker(lines))
	fmt.Println("test 2:", playPokerWithJokers(lines))
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	ALL_UNIQUE      = 1
)

type Hand struct {
	hand     string
	handType int
	bid      int
}

func playPoker(lines []string) int64 {
	// function info
	hands := []Hand{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand, bid := parts[0], h.ParseInt(parts[1])

		cards := make([]int, 15)
		for i := range hand {
			cards[getCardValue(hand[i])]++
		}
		handType := calculateHandType(cards)

		hands = append(hands, Hand{hand: hand, bid: bid, handType: handType})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return compareHands(hands[i].hand, hands[j].hand)
		} else {
			return hands[i].handType < hands[j].handType
		}
	})

	// fmt.Println(hands)
	sum := int64(0)
	for i, hand := range hands {
		sum += int64(hand.bid * (i + 1))
		// fmt.Println(sum)
	}

	return sum
}

func playPokerWithJokers(lines []string) int64 {
	// function info
	hands := []Hand{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand, bid := parts[0], h.ParseInt(parts[1])

		cards := make([]int, 15)
		for i := range hand {
			cards[getCardValueWithJoker(hand[i])]++
		}
		jokers := cards[1]

		handType := calculateHandType(cards[2:])
		jokeHandType := addJokers(handType, jokers)

		hands = append(hands, Hand{hand: hand, bid: bid, handType: jokeHandType})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return compareHandsWithJokers(hands[i].hand, hands[j].hand)
		} else {
			return hands[i].handType < hands[j].handType
		}
	})

	// fmt.Println(hands)
	sum := int64(0)
	for i, hand := range hands {
		sum += int64(hand.bid * (i + 1))
		// fmt.Println(sum)
	}

	return sum
}

func addJokers(handType int, jokers int) int {
	if jokers == 0 {
		return handType
	}

	if handType == FOUR_OF_A_KIND {
		return FIVE_OF_A_KIND
	} else if handType == THREE_OF_A_KIND {
		return THREE_OF_A_KIND + jokers + 1 // 4 of kind is +2 above 3 of kind, so we need to add +1
	} else if handType == TWO_PAIR {
		return FULL_HOUSE
	} else if handType == ONE_PAIR {
		switch jokers {
		case 1:
			return THREE_OF_A_KIND
		case 2:
			return FOUR_OF_A_KIND
		case 3:
			return FIVE_OF_A_KIND
		}
	}

	// Add jokers when there are no current matches
	switch jokers {
	case 1:
		return ONE_PAIR
	case 2:
		return THREE_OF_A_KIND
	case 3:
		return FOUR_OF_A_KIND
	default:
		return FIVE_OF_A_KIND
	}
}

func compareHands(h1 string, h2 string) bool {
	for i := range h1 {
		v1, v2 := getCardValue(h1[i]), getCardValue(h2[i])
		if v1 != v2 {
			return v1 < v2
		}
	}
	return false
}
func compareHandsWithJokers(h1 string, h2 string) bool {
	for i := range h1 {
		v1, v2 := getCardValueWithJoker(h1[i]), getCardValueWithJoker(h2[i])
		if v1 != v2 {
			return v1 < v2
		}
	}
	return false
}
func getCardValue(c byte) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		log.Fatalf("'%v' is not a valid card value", c)
	}
	return -1
}
func getCardValueWithJoker(c byte) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 1 // Jokers are less than 2
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		log.Fatalf("'%v' is not a valid card value", c)
	}
	return -1
}
func calculateHandType(cards []int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
	// fmt.Println(cards)

	handType := ALL_UNIQUE
	for _, occ := range cards {
		if occ == 5 {
			handType = FIVE_OF_A_KIND
			break
		} else if occ == 4 {
			handType = FOUR_OF_A_KIND
			break
		} else if occ == 3 && handType == TWO_PAIR || occ == 2 && handType == THREE_OF_A_KIND {
			handType = FULL_HOUSE
			break
		} else if occ == 3 {
			handType = THREE_OF_A_KIND
		} else if occ == 2 && handType == ONE_PAIR {
			handType = TWO_PAIR
			break
		} else if occ == 2 {
			handType = ONE_PAIR
		}
	}
	return handType
}
