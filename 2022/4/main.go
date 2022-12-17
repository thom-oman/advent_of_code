package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var input string
var lines []string
var c int

func init() {
	input = strings.TrimRight(input, "\n")
	lines = strings.Split(input, "\n")
}

type section struct {
	Min, Max int64
}

// Part 1
// func (s section) HasOverlap(other section) bool {
// 	return (s.Min >= other.Min && s.Max <= other.Max) || (s.Min <= other.Min && s.Max >= other.Max)
// }

// Part 2
func (s section) HasOverlap(other section) bool {
	if s.Max < other.Min || other.Max < s.Min {
		return false
	}

	return true
}

func main() {
	for _, l := range lines {
		sections := make([]section, 2)
		for idx, a := range strings.Split(l, ",") {
			q := strings.Split(a, "-")
			min, _ := strconv.ParseInt(q[0], 10, 32)
			max, _ := strconv.ParseInt(q[1], 10, 32)
			s := section{Min: min, Max: max}
			sections[idx] = s
		}
		if sections[0].HasOverlap(sections[1]) {
			c++
		}
	}
	fmt.Println(c)
}
