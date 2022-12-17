package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed inputs.txt
var file string
var lines []string
var score int

func init() {
	// do this in init (not main) so test file has same input
	file = strings.TrimRight(file, "\n")
	lines = strings.Split(file, "\n")
}

// Part 1
// func main() {
// 	for _, l := range lines {
// 		mid := len(l) / 2
// 		side1 := l[:mid]
// 		side2 := l[mid:]
// 		if len(side1) != len(side2) {
// 			log.Fatal("uh oh ")
// 		}
// 		seenIn1 := make(map[rune]bool)
// 		for _, c := range side1 {
// 			seenIn1[c] = true
// 		}
// 		found := false
// 		for _, c := range side2 {
// 			if found {
// 				continue
// 			}
// 			_, ok := seenIn1[c]
// 			if ok {
// 				score += priorityScore(c)
// 				found = true
// 			}
// 		}
// 	}
// 	fmt.Println(score)
// }

// Part 2
func main() {
	if len(lines)%3 != 0 {
		log.Fatal("Not a clean number of groups")
	}

	for i := 0; i+2 < len(lines); i += 3 {
		a := lines[i]
		b := lines[i+1]
		c := lines[i+2]

		seen := make(map[rune]int)
		// Mark all characters in first line as seen
		for _, char := range a {
			seen[char] = 1
		}

		// Mark all new characters as seen, mark the ones
		// from previous line that are missing from second
		// as not seen
		for _, char := range b {
			_, ok := seen[char]
			if ok {
				seen[char] = 2
			}
		}

		// Iterate through third in gropu
		for _, char := range c {
			val, ok := seen[char]
			if ok && val == 2 {
				score += priorityScore(char)
				break
			}
		}
	}
	fmt.Println(score)
}

func priorityScore(s rune) int {
	// a-z is 97-122: priority 1-26
	if s >= 97 && s <= 122 {
		return int(s) - 96
	}
	// A-Z is 65-90: priority 27-52
	if s >= 65 && s <= 90 {
		return int(s) - 64 + 26
	}
	// 0-9 is 48-57
	return 0
}
