package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	grid [][]rune
)

func main() {
	parseInput()
	rollNorth()
	determineLoad()
}

func determineLoad() {
	var (
		tot int
		maxLoad = len(grid)
	)
	for i := range grid {
		rowLoad := maxLoad - i
		for j := range grid[i] {
			if grid[i][j] == 'O' {
				tot += rowLoad
			}
		}
	}

	fmt.Println("total load:", tot)
}

func rollNorth() {
	for rowIdx := range grid {
		if rowIdx == 0 {
			continue
		}
	
		for colIdx := range grid[rowIdx] {
			if grid[rowIdx][colIdx] == 'O' {
				newIdx := rowIdx

				for newIdx > 0 && grid[newIdx-1][colIdx] == '.' {
					newIdx--
				}
				if newIdx != rowIdx {
					grid[rowIdx][colIdx] = '.'
					grid[newIdx][colIdx] = 'O'
				}
			}
		}
	}
}

func parseInput() {
	bs, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(bs), "\n") {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}
}
