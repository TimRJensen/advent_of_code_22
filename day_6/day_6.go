package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	if len(os.Args) < 3 || os.Args[1] != "part" || !strings.Contains("12", os.Args[2]) {
		log.Fatal("usage: part <1|2>")
	}

	if arg := os.Args[2]; arg == "1" {
		fmt.Println("result:", getMarkerPosition(buffer, 4))
	} else {
		fmt.Println("result:", getMarkerPosition(buffer, 14))
	}
}
