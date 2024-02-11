package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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

func getInput(file *os.File) (result [][]int) {
	result = make([][]int, 128, 256)
	scanner := bufio.NewScanner(file)
	length := 0

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			if length++; length == len(result) {
				result = append(result, []int{})
			}

			continue
		}

		val, _ := strconv.Atoi(s)
		result[length] = append(result[length], val)
	}

	return result
}

func main() {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getMaxValue(getInput(file)))
	} else {
		fmt.Println("result:", getMaxValues(getInput(file), 3))
	}
}
