package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
 * part_1
 */
const (
	win  int = 6
	draw int = 3
)

var (
	hands = [...]string{"AX", "BY", "CZ"}
)

func handToInt(s string) (int, error) {
	for i := 0; i < len(hands); i++ {
		if strings.Contains(hands[i], s) {
			return i, nil
		}
	}

	return 0, errors.New("Invalid input")
}

func getScore(list []int) (result int) {
	for i := 0; i < len(list); i += 2 {
		switch {
		case (list[i+1]+2)%len(hands) == list[i]:
			result += list[i+1] + win
		case list[i+1] == list[i]:
			result += list[i+1] + draw
		default:
			result += list[i+1]
		}

		result++
	}

	return result
}

/**
 * part_2
 */
const (
	x int = iota
	y
	z
)

func getScoreForced(list []int) (result int) {
	for i := 0; i < len(list); i += 2 {
		switch list[i+1] {
		case z:
			result += (list[i]+1)%len(hands) + win
		case y:
			result += list[i] + draw
		default:
			result += (list[i] + 2) % len(hands)
		}

		result++
	}

	return result
}

/**
 * driver
 */
func getInput(buffer []byte) (result []int) {
	lines := strings.Split(string(buffer), "\n")

	for _, line := range lines {
		round := strings.Split(line, " ")

		for _, hand := range round {
			if val, err := handToInt(hand); err == nil {
				result = append(result, val)
			}
		}
	}

	return result
}

func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getScore(getInput(buffer)))
	} else {
		fmt.Println("result:", getScoreForced(getInput(buffer)))
	}
}
