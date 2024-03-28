package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	modulus   = 40
	remainder = 20
	litCell   = "#"
	darkCell  = "."
)

/**
 * part_1
 */
func signalStrength(immediates []int) (result int) {
	deferred, i := 0, 0
	cycle, x := 1, 1

	for i < len(immediates) || deferred != 0 {
		if cycle%modulus == remainder { //magic ints
			result += cycle * x
		}

		if deferred != 0 {
			x += deferred
			deferred = 0
		} else {
			if i < len(immediates) && immediates[i] != 0 {
				deferred = immediates[i]
			}

			i++
		}

		cycle++
	}

	return result
}

/**
 * part_2
 */
func draw(immediates []int) (result string) {
	deferred, i := 0, 0
	cycle, x := 1, 2

	for i < len(immediates) || deferred != 0 {
		mod := cycle % modulus

		switch {
		case mod == 0:
			result += "\n"
		case x-1 <= mod && x+1 >= mod:
			result += litCell
		default:
			result += darkCell
		}

		if deferred != 0 {
			x += deferred
			deferred = 0
		} else {
			if i < len(immediates) && immediates[i] != 0 {
				deferred = immediates[i]
			}

			i++
		}

		cycle++
	}

	return result
}

/**
 * driver
 */
func parseInput(buffer []byte) (result []int) {
	for _, line := range strings.Split(string(buffer), "\n") {
		if fields := strings.Fields(line); len(fields) == 2 {
			if val, err := strconv.Atoi(fields[1]); err == nil {
				result = append(result, val)
			}
		} else {
			result = append(result, 0)
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
		fmt.Println("result:", signalStrength(parseInput(buffer)))
	} else {
		fmt.Println("result:")
		fmt.Println(draw(parseInput(buffer)))
	}
}
