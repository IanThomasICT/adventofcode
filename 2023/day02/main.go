package main

import (
	"fmt"
	h "ianthomasict/adventofcode/helpers"
	"log"
	"strconv"
	"strings"
)

func main() {

	fileName := "input.txt"
	//fileName := "test.txt"
	lines, err := h.ReadLinesAsArray(fileName)
	if err != nil {
		log.Fatalf("failed to get lines as []string: %s", err)
	}

	sum := part1(lines)
	fmt.Println("part 1:", sum)
	sum = part2(lines)
	fmt.Println("part 2:", sum)

}

func part1(lines []string) int {

	gameSum := 0
	for _, line := range lines {
		//line = strings.ReplaceAll(line, " ", "")
		parts := strings.Split(line, ": ")
		header, game := parts[0], parts[1]

		possibleGame := checkIfPossibleGame(game)
		if possibleGame {
			gameNum, err := strconv.Atoi(strings.Split(header, " ")[1])
			if err != nil {
				log.Fatalf("Failed to convert %s into an integer", strings.Split(header, " ")[1])
			}
			fmt.Println("adding possible game", gameNum)
			gameSum += gameNum
		}
	}
	return gameSum
}

func checkIfPossibleGame(game string) bool {
	RED_MAX, GREEN_MAX, BLUE_MAX := 12, 13, 14
	//blocks := [3]int{0, 0, 0} // [r,g,b]

	draws := strings.Split(game, "; ")
	for _, draw := range draws {
		hands := strings.Split(draw, ", ")

		for _, hand := range hands {
			h := strings.Split(hand, " ")
			color := h[1]
			val, err := strconv.Atoi(h[0])
			if err != nil {
				log.Fatalf("Failed to convert %s into an integer", h[0])
			}

			fmt.Println(val, color)

			switch color {
			case "red":
				if val > RED_MAX {
					fmt.Printf("%d > %d\n", val, RED_MAX)
					return false
				}
			case "green":
				if val > GREEN_MAX {
					fmt.Printf("%d > %d\n", val, GREEN_MAX)
					return false
				}
			case "blue":
				if val > BLUE_MAX {
					fmt.Printf("%d > %d\n", val, BLUE_MAX)
					return false
				}
			}

		}
	}
	return true
}

func part2(lines []string) int {

	gameSum := 0
	for _, line := range lines {
		game := strings.Split(line, ": ")[1]

		powerVal := getPowerSetOfMinimumBlocks(game)
		gameSum += powerVal
	}
	return gameSum
}

func getPowerSetOfMinimumBlocks(game string) int {
	blocks := [3]int{0, 0, 0} // [r,g,b]

	draws := strings.Split(game, "; ")
	for _, draw := range draws {
		hands := strings.Split(draw, ", ")

		for _, hand := range hands {
			h := strings.Split(hand, " ")
			color := h[1]
			val, err := strconv.Atoi(h[0])
			if err != nil {
				log.Fatalf("Failed to convert %s into an integer", h[0])
			}

			switch color {
			case "red":
				if val > blocks[0] {
					blocks[0] = val
				}
			case "green":
				if val > blocks[1] {
					blocks[1] = val
				}
			case "blue":
				if val > blocks[2] {
					blocks[2] = val
				}
			}
		}
	}
	powerSet := blocks[0] * blocks[1] * blocks[2]
	fmt.Println("PowerSet", blocks, "->", powerSet)
	return powerSet
}
