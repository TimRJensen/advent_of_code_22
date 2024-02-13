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
 * stack
 */
type stack[T rune | string] struct {
	items []T
}

func newStack[T rune | string]() *stack[T] {
	return &stack[T]{items: make([]T, 0, 256)}
}

func (stack *stack[T]) isEmpty() bool {
	return len(stack.items) == 0
}

func (stack *stack[T]) push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *stack[T]) pop() (result T, flag bool) {
	if stack.isEmpty() {
		return *new(T), false
	}
	result = stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return result, true
}

func (stack *stack[T]) String() string {
	return fmt.Sprintf("%T%v", stack, stack.items)
}

/**
 * part_1
 */
func getTopSuppliesPart1(supplies []*stack[string], instructions [][]int) (result string) {
	for _, instruction := range instructions {
		for i := 0; i < instruction[move]; i++ {
			if line, ok := supplies[instruction[from]].pop(); ok {
				supplies[instruction[to]].push(line)
			}
		}
	}

	for _, stack := range supplies {
		if s, ok := stack.pop(); ok {
			result += s
		}
	}

	return result
}

/**
 * part_2
 */
func getTopSuppliesPart2(supplies []*stack[string], instructions [][]int) (result string) {
	for _, instruction := range instructions {
		stack := newStack[string]()

		for i := 0; i < instruction[move]; i++ {
			if line, ok := supplies[instruction[from]].pop(); ok {
				stack.push(line)
			}
		}

		for !stack.isEmpty() {
			if line, ok := stack.pop(); ok {
				supplies[instruction[to]].push(line)
			}
		}
	}

	for _, stack := range supplies {
		if s, ok := stack.pop(); ok {
			result += s
		}
	}

	return result
}

/**
 * driver
 */
const (
	move      = 0
	from      = 1
	to        = 2
	increment = 4
)

func getInstructions(lines []string) (result [][]int) {
	for _, line := range lines {
		i := len(result)
		result = append(result, make([]int, 0, 256))

		for j, instruction := range strings.Split(line, " ") {
			if j%2 == 1 {
				val, _ := strconv.Atoi(instruction)

				if j > 1 {
					result[i] = append(result[i], val-1)
				} else {
					result[i] = append(result[i], val)
				}
			}
		}
	}

	return result
}

func getSupplies(lines []string) (result []*stack[string]) {
	stack := newStack[string]()

	for _, line := range lines {
		stack.push(line + " ")
	}

	for !stack.isEmpty() {
		if line, ok := stack.pop(); ok {
			for i := 0; i < len(line); i += increment {
				if s := strings.Trim(line[i:i+increment], " []"); s != "" {
					if s[0] >= 48 && s[0] <= 57 {
						result = append(result, newStack[string]())
					} else {
						result[i/increment].push(s)
					}
				}
			}
		}
	}

	return result
}

func getInput(buffer []byte) (supplies []*stack[string], instructions [][]int) {
	lines := strings.Split(string(buffer), "\n")
	i := slices.Index(lines, "")

	return getSupplies(lines[:i]), getInstructions(lines[i+1:])
}

func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	supplies, instructions := getInput(buffer)

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getTopSuppliesPart1(supplies, instructions))
	} else {
		fmt.Println("result:", getTopSuppliesPart2(supplies, instructions))
	}
}
