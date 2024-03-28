package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

const (
	up = iota
	left
	down
	right
	dirs
	dimensions = 2
	corners    = dimensions * 2
)

var (
	dxs = [...]int{0, -1, 0, 1}
	dys = [...]int{-1, 0, 1, 0}
)

func nextPoint(x int, y int, d int) (int, int) {
	return x + dxs[d], y + dys[d]
}

/**
 * part_1
 */
func visibleTrees(list [][]int) (result int) {
	max_y := len(list)
	max_x := len(list[0])

	for y := 1; y < max_y-1; y++ {
		for x := 1; x < max_x-1; x++ {
			for d := 0; d < dirs; d++ {
				flag := true
				dx, dy := nextPoint(x, y, d)

				for !(dx < 0 || dx == max_x || dy < 0 || dy == max_y) {
					if list[dy][dx] >= list[y][x] {
						flag = false
						break
					}

					dx, dy = nextPoint(dx, dy, d)
				}

				if flag {
					result++
					break
				}
			}
		}
	}

	return dimensions*(max_y+max_x) - corners + result
}

/**
 * part_2
 */
func scenicScore(list [][]int) (result int) {
	max_y := len(list)
	max_x := len(list[0])
	products := []int{}

	for y := 1; y < max_y-1; y++ {
		for x := 1; x < max_x-1; x++ {
			product := 1

			for d := 0; d < dirs; d++ {
				dx, dy := nextPoint(x, y, d)
				viewDist := 0

				for !(dx < 0 || dx == max_x || dy < 0 || dy == max_y) {
					viewDist++

					if list[dy][dx] >= list[y][x] {
						break
					}

					dx, dy = nextPoint(dx, dy, d)
				}

				product *= viewDist
			}

			if product > 0 {
				products = append(products, product)
			}
		}
	}

	slices.SortFunc(products, func(a int, b int) int {
		return b - a
	})

	return products[0]
}

/**
 * driver
 */
func parseInput(buffer []byte) (result [][]int) {
	for _, line := range strings.Split(string(buffer), "\n") {
		l := len(result)
		result = append(result, []int{})

		for _, c := range line {
			if unicode.IsDigit(c) {
				result[l] = append(result[l], int(c-'0'))
			}
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
		fmt.Println("result:", visibleTrees(parseInput(buffer)))
	} else {
		fmt.Println("result:", scenicScore(parseInput(buffer)))
	}
}
