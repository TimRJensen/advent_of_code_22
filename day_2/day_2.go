package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock int = iota
	paper
	scissors
	shapes int = 3
	win    int = 6
	draw   int = 3
)

var (
	atohand = map[string]int{"AX": rock, "BY": paper, "CZ": scissors}
)

/**
 * part_1
 */
func score(list []int) (result int) {
	for i := 0; i < len(list); i += 2 {
		switch {
		case (list[i+1]+2)%shapes == list[i]:
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
func scoreForced(list []int) (result int) {
	for i := 0; i < len(list); i += 2 {
		switch list[i+1] {
		case scissors:
			result += (list[i]+1)%shapes + win
		case paper:
			result += list[i] + draw
		default:
			result += (list[i] + 2) % shapes
		}
		result++
	}

	return result
}

/**
 * driver
 */
func parseInput(buffer []byte) (result []int) {
	atohandWithDefault := func(s string) int {
		for key, val := range atohand {
			if strings.Contains(key, s) {
				return val
			}
		}
		return -1
	}

	for _, line := range strings.Split(string(buffer), "\n") {
		if fields := strings.Fields(line); len(fields) == 2 {
			for _, hand := range fields {
				result = append(result, atohandWithDefault(hand))
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

	if len(os.Args) < 3 || os.Args[1] != "part" || !strings.Contains("12", os.Args[2]) {
		log.Fatal("usage: part <1|2>")
	}

	if arg := os.Args[2]; arg == "1" {
		fmt.Println("result:", score(parseInput(buffer)))
	} else {
		fmt.Println("result:", scoreForced(parseInput(buffer)))
	}
}
