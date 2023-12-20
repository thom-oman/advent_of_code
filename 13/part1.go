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
		// fmt.Println(strings.Repeat("#", 10), "Pattern", i+1, strings.Repeat("#", 10))
		findReflectionPoint(patterns[i])
	}
}

func findReflectionPoint(ps []string) {
	var transposed []string

	for i := range ps[0] {
		var s string

		for j := range ps {
			s += string(ps[j][i])
		}
		transposed = append(transposed, s)
	}

	xx, xy := reflectionAxis(ps)
	if xx != -1 {
		sum += 100*xy
	}
	yx, yy := reflectionAxis(transposed)
	if yx != -1 {
		sum += yy
	}
}

func reflectionAxis(ls []string) (int, int) {
	for i := range ls {
		var (
			d int
			match = true
		)

		for {
			x, y := i-d, i+d+1

			if x < 0 || y >= len(ls) {
				if d < 1 {
					match = false
				}
				break
			}

			if ls[x] != ls[y] {
				match = false
				break
			}

			d++
		}

		if match {
			return i, i+1
		}
	}

	return -1, -1
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
