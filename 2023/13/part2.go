package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	sum      int
	patterns [][]string
)

func main() {
	parseInput()
	determineReflectionPoints()
	fmt.Println("Total:", sum)
}

func determineReflectionPoints() {
	for i := range patterns {
		findReflectionPoint(patterns[i])
	}
}

func findReflectionPoint(ps []string) {
	xx, xy := reflectionAxis(ps)
	yx, yy := reflectionAxis(transpose(ps))

	if xx != -1 {
		sum += 100 * xy
	} else if yx != -1 {
		sum += yy
	} else {
		panic("cannot find pattern")
	}
}

func transpose(arr []string) []string {
	var transposed []string

	for i := range arr[0] {
		var s string

		for j := range arr {
			s += string(arr[j][i])
		}
		transposed = append(transposed, s)
	}
	return transposed
}

func reflectionAxis(ls []string) (int, int) {
	for i := range ls[:len(ls)-1] {
		diff := differentChars(ls[i], ls[i+1])
		if diff > 1 {
			continue
		}

		// Find closest boundary
		steps := distanceToBoundary(ls, i)
		if d := distanceToBoundary(ls, i+1); d < steps {
			steps = d
		}

		// If we're at the end don't give a fuck no more
		if steps == 0 && diff == 1 {
			return i, i + 1
		}

		for j := 1; j <= steps; j++ {
			diff += differentChars(ls[i-j], ls[i+1+j])
		}

		if diff != 1 {
			continue
		}

		return i, i + 1
	}

	return -1, -1
}

func distanceToBoundary(ls []string, i int) int {
	if i < 0 || i+1 >= len(ls) {
		return 0
	}
	
	toStart, toEnd := i, len(ls) - i - 1

	if toStart > toEnd {
		return toEnd
	}

	return toStart
}


func differentChars(a, b string) int {
	var n int
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n
}

func parseInput() {
	bs, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	var cur []string

	for _, l := range strings.Split(string(bs), "\n") {
		if len(l) > 0 {
			cur = append(cur, l)
			continue
		}

		if len(cur) > 0 {
			patterns = append(patterns, cur)
			cur = make([]string, 0)
		}
	}
}
