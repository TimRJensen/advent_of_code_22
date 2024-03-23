package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var (
	offsets = map[bool]map[rune]int{false: {'a': 1}, true: {'A': 27}}
)

/**
 * part_1
 */
func itemPriorities(lines []string) (result int) {
	for _, line := range lines {
		mid := len(line) / 2
		a := line[:mid]
		b := line[mid:]

		for _, c := range a {
			if !strings.ContainsRune(b, c) {
				continue
			}

			for key, val := range offsets[unicode.IsUpper(c)] {
				result += int(c-key) + val
			}

			break
		}
	}

	return result
}

/**
 * part_2
 */
func badgePriorities(lines []string, elves int) (result int) {
	for i := 0; i < len(lines); i += elves {
		seen := map[rune]complex64{}
		flag := false
		a := lines[i]

		for _, c := range a {
			if _, ok := seen[c]; !ok {
				for key, val := range offsets[unicode.IsUpper(c)] {
					seen[c] = complex(float32(int(c-key)+val), 1)
				}
			}

			for j := 1; j < elves; j++ {
				b := lines[i+j]

				if !strings.ContainsRune(b, c) {
					continue
				}

				if val := seen[c]; int(imag(val)) == j {
					val += 0 + 1i

					if r, i := real(val), imag(val); int(i) == elves {
						result += int(r)
						flag = true

						break
					}

					seen[c] = val
				}
			}

			if flag {
				break
			}
		}
	}

	return result
}

/**
 * driver
 */
func parseInput(buffer []byte) (result []string) {
	result = append(result, strings.Split(string(buffer), "\n")...)
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
		fmt.Println("result:", itemPriorities(parseInput(buffer)))
	} else {
		fmt.Println("result:", badgePriorities(parseInput(buffer), 3))
	}
}
