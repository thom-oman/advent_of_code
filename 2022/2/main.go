package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs.txt
var inputs string
var rounds []string
var score int
var scores map[string]int

func init() {
	// do this in init (not main) so test file has same input
	inputs = strings.TrimRight(inputs, "\n")
	if len(inputs) == 0 {
		panic("empty input.txt file")
	}
	rounds = strings.Split(inputs, "\n")
	scores = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

}

// Part 1
// func main() {
// 	for _, round := range rounds {
// 		s := strings.Split(round, " ")
// 		opp := s[0]
// 		resp := s[1]

// 		switch resp {
// 		case "X": // Rock
// 			score += 1 + matchScore(resp, opp)
// 		case "Y": // Paper
// 			score += 2 + matchScore(resp, opp)
// 		case "Z": // Scissors
// 			score += 3 + matchScore(resp, opp)
// 		}
// 	}

// 	fmt.Println(score)
// }
//
// func matchScore(x, y string) int {
// 	switch x {
// 	case "X": // Rock
// 		x = "A"
// 		if y == "C" {
// 			return 6
// 		}
// 	case "Y": // Paper
// 		x = "B"
// 		if y == "A" {
// 			return 6
// 		}
// 	case "Z": // Scissors
// 		x = "C"
// 		if y == "B" {
// 			return 6
// 		}
// 	}
// 	if x == y {
// 		return 3
// 	}

// 	return 0
// }

// Part 2
func main() {
	for _, round := range rounds {
		s := strings.Split(round, " ")
		score += matchScore(s[0], s[1])
	}

	fmt.Println(score)
}

func matchScore(x, y string) int {
	switch y {
	case "X": // Lose
		resp := map[string]string{
			"A": "C",
			"B": "A",
			"C": "B",
		}
		return scores[resp[x]]
	case "Y": // Draw
		return 3 + scores[x]
	case "Z": // Win
		resp := map[string]string{
			"A": "B",
			"B": "C",
			"C": "A",
		}
		return 6 + scores[resp[x]]
	}
	return 0
}
