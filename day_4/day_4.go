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
func contains(list []*section) (result int) {
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
func intersects(list []*section) (result int) {
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
func parseInput(buffer []byte) (result []*section) {
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

	if len(os.Args) < 3 || os.Args[1] != "part" || !strings.Contains("12", os.Args[2]) {
		log.Fatal("usage: part <1|2>")
	}

	if arg := os.Args[2]; arg == "1" {
		fmt.Println("result:", contains(parseInput(buffer)))
	} else {
		fmt.Println("result:", intersects(parseInput(buffer)))
	}
}
