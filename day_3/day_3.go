package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

var (
	starts = [...]int{int('a'), 1, int('A'), 27}
)

type entry struct {
	value int
	count int
}

func (entry entry) String() string {
	return fmt.Sprintf("value: %d; count: %d", entry.value, entry.count)
}

func getPriorities(m map[rune]*entry) (result int) {
	for _, item := range m {
		result += item.count * item.value
	}

	return result
}

/**
 * part_1
 */
func getInputPart1(buffer []byte) (result map[rune]*entry) {
	result = make(map[rune]*entry)

	for _, line := range strings.Split(string(buffer), "\n") {
		m := len(line) / 2
		a := line[:m]
		b := line[m:]

		for _, c := range a {
			if !strings.ContainsRune(b, c) {
				continue
			}

			if item, ok := result[c]; ok {
				item.count++
			} else {
				i := 0

				if unicode.IsUpper(c) {
					i = 2
				}

				result[c] = &entry{

					value: (int(c) - starts[i]) + starts[i+1],
					count: 1,
				}
			}

			break
		}
	}

	return result
}

/**
 * part_2
 */
func getInputPart2(buffer []byte) (result map[rune]*entry) {
	result = make(map[rune]*entry)
	lines := strings.Split(string(buffer), "\n")

	for i := 0; i < len(lines); i += 3 {
		m := make(map[rune]*entry)

		for _, line := range lines[i+1 : i+3] {
			cs := make([]rune, len(lines[i]))

			for _, c := range lines[i] {
				if !strings.ContainsRune(line, c) {
					continue
				}

				if slices.Contains(cs, c) {
					continue
				}

				cs = append(cs, c)

				if item, ok := m[c]; ok {
					item.count++
				} else {
					m[c] = &entry{count: 1}
				}
			}
		}

		for key, item := range m {
			if item.count < 2 {
				continue
			}

			if item, ok := result[key]; ok {
				item.count++
			} else {
				i := 0

				if unicode.IsUpper(key) {
					i = 2
				}

				result[key] = &entry{value: int(key) - starts[i] + starts[i+1], count: 1}
			}

			break
		}
	}

	return result
}

/**
 * driver
 */
func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getPriorities(getInputPart1(buffer)))
	} else {
		fmt.Println("result:", getPriorities(getInputPart2(buffer)))
	}
}
