package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	up = iota
	left
	down
	right
	dirs
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
func getVisibleTrees(list [][]int) (result int) {
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

	return 2*(max_y+max_x-2) + result //magic int
}

/**
 * part_2
 */
func getScenicScore(list [][]int) (result int) {
	max_y := len(list)
	max_x := len(list[0])
	products := make([]int, 0, 128)

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
		return a - b
	})

	return products[len(products)-1]
}

/**
 * driver
 */
func getInput(buffer []byte) (result [][]int) {
	for _, line := range strings.Split(string(buffer), "\n") {
		l := len(result)
		result = append(result, make([]int, 0, 128))

		for _, c := range line {
			result[l] = append(result[l], int(c-'0'))
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
		fmt.Println("result:", getVisibleTrees(getInput(buffer)))
	} else {
		fmt.Println("result:", getScenicScore(getInput(buffer)))
	}
}
