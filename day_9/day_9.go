package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	up = iota
	right
	down
	left
	dirs
	gap = 2
)

var (
	dxs    = [dirs]int{0, 1, 0, -1}
	dys    = [dirs]int{1, 0, -1, 0}
	atodir = map[string]int{"U": up, "L": left, "D": down, "R": right}
)

type vector struct {
	x int
	y int
}

type knot struct {
	v    vector
	next *knot
}

func (k *knot) move(x int, y int) {
	k.v.x += 1 * x
	k.v.y += 1 * y
}

type instruction struct {
	dir  int
	step int
}

/**
 * part_1
 */
func traverse(instructions []*instruction, knots int) int {
	visisted := map[complex64]bool{}
	visisted[0+0i] = true

	tail := &knot{v: vector{x: 0, y: 0}, next: nil}
	head := &knot{v: vector{x: 0, y: 0}, next: tail}

	for i := 2; i < knots; i++ {
		next := &knot{v: vector{x: 0, y: 0}, next: head}
		head = next
	}

	for _, instruction := range instructions {
		for i := 0; i < instruction.step; i++ {
			knot := head
			knot.move(dxs[instruction.dir], dys[instruction.dir])

			for next := knot.next; next != nil; next = knot.next {
				dx := knot.v.x - next.v.x
				dy := knot.v.y - next.v.y

				switch {
				case max(dx, -dx) == gap && max(dy, -dy) == gap:
					next.move(dx/gap, dy/gap)
				case max(dx, -dx) == gap:
					next.move(dx/gap, dy)
				case max(dy, -dy) == gap:
					next.move(dx, dy/gap)
				default:
					knot = tail
					continue
				}

				if next == tail {
					key := complex(float32(next.v.x), float32(next.v.y))
					if _, ok := visisted[key]; !ok {
						visisted[key] = true
					}
				}

				knot = next
			}
		}
	}

	return len(visisted)
}

/**
 * part_2
 */

/**
 * driver
 */
func getInput(buffer []byte) (instructions []*instruction) {
	atoiWithDefault := func(s string) int {
		if val, err := strconv.Atoi(s); err == nil {
			return val
		}
		return 0
	}
	atodirWithDefault := func(s string) int {
		if val, ok := atodir[s]; ok {
			return val
		}
		return up
	}

	for _, line := range strings.Split(string(buffer), "\n") {
		if fields := strings.Fields(line); len(fields) == 2 {
			instructions = append(instructions, &instruction{
				dir:  atodirWithDefault(fields[0]),
				step: atoiWithDefault(fields[1]),
			})
		}
	}

	return instructions
}

func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", traverse(getInput(buffer), 2))
	} else {
		fmt.Println("result:", traverse(getInput(buffer), 10))
	}
}
