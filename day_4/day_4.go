package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type section struct {
	min int
	max int
}

func (a *section) contains(b *section) bool {
	return a.min <= b.min && a.max >= b.max
}

func (a *section) intersects(b *section) bool {
	return a.min <= b.max && a.max >= b.min
}

/**
 * part_1
 */
func getContains(list []*section) (result int) {
	for i := 0; i < len(list); i += 2 {
		if list[i].contains(list[i+1]) || list[i+1].contains(list[i]) {
			result++
		}
	}

	return result
}

/**
 * part_2
 */
func getIntersects(list []*section) (result int) {
	for i := 0; i < len(list); i += 2 {
		if list[i].intersects(list[i+1]) {
			result++
		}
	}

	return result
}

/**
 * driver
 */
func getInput(buffer []byte) (result []*section) {
	for _, line := range strings.Split(string(buffer), "\n") {
		a := &section{}
		b := &section{}

		if i, _ := fmt.Sscanf(line, "%d-%d,%d-%d", &a.min, &a.max, &b.min, &b.max); i == 4 {
			result = append(result, a, b)
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
		fmt.Println("result:", getContains(getInput(buffer)))
	} else {
		fmt.Println("result:", getIntersects(getInput(buffer)))
	}
}
