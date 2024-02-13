package main

import (
	"fmt"
	"log"
	"os"
)

/**
 * part_1 & part_2
 */
func getMarkerPosition(buffer []byte, length int) int {
	for i := 0; i < len(buffer)-length; i++ {
		set := map[byte]byte{buffer[i]: buffer[i]}
		j := i + 1

		for ; j < i+length; j++ {
			if _, ok := set[buffer[j]]; ok {
				break
			}

			set[buffer[j]] = buffer[j]
		}

		if len(set) == length {
			return j
		}
	}

	return 0
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
		fmt.Println("result:", getMarkerPosition(buffer, 4))
	} else {
		fmt.Println("result:", getMarkerPosition(buffer, 14))
	}
}
