package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

/**
 * part_1 & part_2
 */
func maxCalories(lists [][]int, elves int) (result int) {
	sums := make([]int, len(lists))

	for i, list := range lists {
		sum := 0

		for _, val := range list {
			sum += val
		}

		sums[i] = sum
	}

	slices.SortFunc(sums, func(a int, b int) int {
		return b - a
	})

	for _, val := range sums[:elves] {
		result += val
	}

	return result
}

/**
 * driver
 */
func parseInput(buffer []byte) (result [][]int) {
	atoiWithDefault := func(s string) int {
		if val, err := strconv.Atoi(s); err == nil {
			return val
		}
		return 0
	}
	list := []int{}

	for _, line := range strings.Split(string(buffer), "\n") {
		if line == "" {
			result = append(result, list)
			list = []int{}

			continue
		}

		list = append(list, atoiWithDefault(line))
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
		fmt.Println("result:", maxCalories(parseInput(buffer), 1))
	} else {
		fmt.Println("result:", maxCalories(parseInput(buffer), 3))
	}
}
