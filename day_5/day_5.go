package main

import (
	"common"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const (
	move = iota
	from
	to
)

var (
	regex = regexp.MustCompile(`[\[\]A-Z1-9 ]{3,4}`)
)

/**
 * part_1 & part_2
 */
func topSupplies(supplies []*common.Stack[string], instructions [][]int, preserveOrder bool) (result string) {
	for _, instruction := range instructions {
		list := []string{}

		for i := 0; i < instruction[move]; i++ {
			if val, ok := supplies[instruction[from]].Pop(); ok {
				list = append(list, val)
			}
		}

		if preserveOrder {
			for i := len(list) - 1; i >= 0; i-- {
				supplies[instruction[to]].Push(list[i : i+1][0])
			}
		} else {
			for i := 0; i < len(list); i++ {
				supplies[instruction[to]].Push(list[i : i+1][0])
			}
		}
	}

	for _, stack := range supplies {
		if s, ok := stack.Pop(); ok {
			result += s
		}
	}

	return result
}

/**
 * driver
 */
func parseInstructions(lines []string) (result [][]int) {
	for _, line := range lines {
		i := len(result)
		result = append(result, []int{})

		for j, field := range strings.Fields(line) {
			val, err := strconv.Atoi(field)

			if err != nil {
				continue
			}

			if j > 1 {
				result[i] = append(result[i], val-1)
			} else {
				result[i] = append(result[i], val)
			}
		}
	}

	return result
}

func parseSupplies(lines []string) (result []*common.Stack[string]) {
	for i := len(lines) - 1; i >= 0; i-- {
		for j, match := range regex.FindAllString(lines[i : i+1][0], -1) {
			if s := strings.Trim(match, "[] "); s != "" {
				if unicode.IsDigit(rune(s[0])) {
					result = append(result, common.NewStack[string]())
				} else {
					result[j].Push(s)
				}
			}
		}
	}

	return result
}

func parseInput(buffer []byte) (supplies []*common.Stack[string], instructions [][]int) {
	lines := strings.Split(string(buffer), "\n")

	if i := slices.Index(lines, ""); i == -1 {
		return supplies, instructions
	} else {
		return parseSupplies(lines[:i]), parseInstructions(lines[i+1:])
	}
}

func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 3 || os.Args[1] != "part" || !strings.Contains("12", os.Args[2]) {
		log.Fatal("usage: part <1|2>")
	}

	supplies, instructions := parseInput(buffer)

	if arg := os.Args[2]; arg == "1" {
		fmt.Println("result:", topSupplies(supplies, instructions, false))
	} else {
		fmt.Println("result:", topSupplies(supplies, instructions, true))
	}
}
