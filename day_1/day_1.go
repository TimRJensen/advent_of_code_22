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
 * part_1
 */
func getMaxValue(lists [][]int) (result int) {
	for _, list := range lists {
		sum := 0

		for _, val := range list {
			sum += val
		}

		if sum > result {
			result = sum
		}
	}

	return result
}

/**
 * part_2
 */
func getMaxValues(lists [][]int, n int) (result int) {
	sums := make([]int, len(lists))

	for i, list := range lists {
		sum := 0

		for _, val := range list {
			sum += val
		}

		sums[i] = sum
	}

	slices.Sort(sums)

	for _, val := range sums[len(lists)-n:] {
		result += val
	}

	return result
}

/**
 * driver
 */
const (
	input = "input.txt"
)

func getInput(buffer []byte) (result [][]int) {
	result = make([][]int, 128, 256)
	length := 0

	for _, line := range strings.Split(string(buffer), "\n") {
		if line == "" {
			if length++; length == len(result) {
				result = append(result, []int{})
			}

			continue
		}

		val, _ := strconv.Atoi(line)
		result[length] = append(result[length], val)
	}

	return result
}

func main() {
	buffer, err := os.ReadFile(input)

	if err != nil {
		log.Fatal(err)
	}

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getMaxValue(getInput(buffer)))
	} else {
		fmt.Println("result:", getMaxValues(getInput(buffer), 3))
	}
}
