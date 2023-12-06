package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var inputs string

func main() {
	digit := regexp.MustCompile("[[:digit:]]+")
	alpha := regexp.MustCompile("[[:alpha:]]+")

	var count int
	for _, l := range strings.Split(inputs, "\n") {
		if len(l) == 0 {
			continue
		}

		split := strings.Split(l, ":")
		// game, _ := strconv.Atoi(string(digit.Find([]byte(split[0]))))

		// possible := true
		mins := make(map[string]int)

		for _, set := range strings.Split(split[1], ";") {
			for _, pick := range strings.Split(set, ",") {
				bs := []byte(pick)
				col := string(alpha.Find(bs))
				c, _ := strconv.Atoi(string(digit.Find(bs)))

				// Part 1
				// switch col {
				// case "red":
				// 	if c > 12 {
				// 		possible = false
				// 	}
				// case "green":
				// 	if c > 13 {
				// 		possible = false
				// 	}
				// case "blue":
				// 	if c > 14 {
				// 		possible = false
				// 	}
				// }

				// Part 2
				col_c := mins[col]
				if c > col_c {
					mins[col] = c
				}
			}
		}

		n := 1
		for _, c := range mins {
			n *= c
		}

		count += n

		// if possible {
		// 	count += game
		// }
	}

	fmt.Println(count)
}
