package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	maxDist              int
	startingX, startingY int
	grid                 [][]rune
	dists                [][]int
	visited              [][]bool
)

var (
	up, down, right, left rune
	goesUp                = []rune{'|', 'L', 'J'}
	goesDown              = []rune{'|', '7', 'F'}
	goesRight             = []rune{'-', 'L', 'F'}
	goesLeft              = []rune{'-', '7', 'J'}
)

func main() {
	parseInput()

	dists = make([][]int, 0)
	visited = make([][]bool, 0)
	for range grid {
		row := make([]int, 0)
		rowb := make([]bool, 0)
		for range grid[0] {
			row = append(row, math.MaxInt)
			rowb = append(rowb, false)
		}
		dists = append(dists, row)
		visited = append(visited, rowb)
	}
	dists[startingX][startingY] = 0

	var (
		crds coords
		stk  []coords
	)

	stk = append(stk, newCoord(startingX, startingY))
	for len(stk) > 0 {
		crds, stk = stk[0], stk[1:]
		d := dists[crds.x][crds.y]

		if maxDist < d {
			maxDist = d
		}

		t := tile(crds.x, crds.y)
		if t == 'S' {
			if Contains(tile(crds.x-1, crds.y), goesDown) {
				stk = append(stk, newCoord(crds.x-1, crds.y))
				dists[crds.x-1][crds.y] = 1
			}
			if Contains(tile(crds.x+1, crds.y), goesUp) {
				stk = append(stk, newCoord(crds.x+1, crds.y))
				dists[crds.x+1][crds.y] = 1
			}
			if Contains(tile(crds.x, crds.y+1), goesLeft) {
				stk = append(stk, newCoord(crds.x, crds.y+1))
				dists[crds.x][crds.y+1] = 1
			}
			if Contains(tile(crds.x, crds.y-1), goesRight) {
				stk = append(stk, newCoord(crds.x, crds.y-1))
				dists[crds.x][crds.y-1] = 1
			}
			visited[crds.x][crds.y] = true
			continue
		}

		if Contains(t, goesUp) {
			if crds.x > 0 {
				if !visited[crds.x - 1][crds.y] {
					stk = append(stk, newCoord(crds.x-1, crds.y))
				}

				if dists[crds.x-1][crds.y] > d+1 {
					dists[crds.x-1][crds.y] = d + 1
				}
			}
		}

		if Contains(t, goesDown) {
			if crds.x+1 < len(grid) {
				if !visited[crds.x + 1][crds.y] {
					stk = append(stk, newCoord(crds.x+1, crds.y))
				}

				if dists[crds.x+1][crds.y] > d+1 {
					dists[crds.x+1][crds.y] = d + 1
				}
			}
		}
		if Contains(t, goesRight) {
			if crds.y+1 < len(grid[0]) {
				if !visited[crds.x][crds.y+1] {
					stk = append(stk, newCoord(crds.x, crds.y+1))
				}
				if dists[crds.x][crds.y+1] > d+1 {
					dists[crds.x][crds.y+1] = d + 1
				}
			}
		}
		if Contains(t, goesLeft) {
			if crds.y > 0 {
				if !visited[crds.x][crds.y-1] {
					stk = append(stk, newCoord(crds.x, crds.y-1))
				}
				if dists[crds.x][crds.y-1] > d+1 {
					dists[crds.x][crds.y-1] = d + 1
				}
			}
		}
		visited[crds.x][crds.y] = true
	}

	fmt.Println("Max dist:", maxDist)
}

func Contains(r rune, arr []rune) bool {
	for i := range arr {
		if r == arr[i] {
			return true
		}
	}
	return false
}

func tile(i, j int) rune {
	return grid[i][j]
}

func parseInput() {
	b, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(b), "\n")

	grid = make([][]rune, 0)

	for i, l := range split {
		if len(l) == 0 {
			continue
		}

		row := make([]rune, 0)

		for j := range l {
			r := rune(l[j])
			row = append(row, r)

			if r == 'S' {
				startingX, startingY = i, j
			}
		}

		grid = append(grid, row)
	}
}

type coords struct {
	x, y int
}

func newCoord(i, j int) coords {
	return coords{i, j}
}
