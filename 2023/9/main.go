package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sum int
	ns = make([][]int,0)
)

func main() {
	parseFile()

	for i := range ns {
		sum += predict(ns[i])
	}
	fmt.Println("Sum:", sum)
}

func predict(bls []int) int {
	diffs := buildDiffs(bls)

	var j int
	for i := len(diffs) - 1; i >= 0; i-- {
		d := diffs[i]
		j += diffs[i][len(d)-1]
	}
	return j
}

func buildDiffs(bls []int) [][]int {
	var (
		diffs [][]int
		allZeroes bool
	)
	diffs = append(diffs, bls)
	for !allZeroes {
		idx := len(diffs) - 1
		var dfs []int
		for i := range diffs[idx] {
			if i == len(diffs[idx]) - 1 {
				continue
			}
			dfs = append(dfs, diffs[idx][i+1] - diffs[idx][i])
		}
		allZeroes = true
		for i := range dfs {
			if dfs[i] != 0 {
				allZeroes = false
			}
		}

		diffs = append(diffs, dfs)
	}
	return diffs
}

func parseFile() {
	b, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	for _, l := range strings.Split(string(b), "\n") {
		if len(l) == 0 {
			continue
		}
		idx := len(ns)
		ns = append(ns, make([]int, 0))

		for _, f := range strings.Fields(l) {
			n, _ := strconv.Atoi(f)
			ns[idx] = append(ns[idx], n)
		}
	}
	fmt.Println(ns)
}
