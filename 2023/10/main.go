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
	startingNode         *node
	grid                 [][]rune
	dists                [][]int
	visited              [][]bool
	lvisited             [][]*node
)

var (
	minPathX, maxPathX = math.MaxInt, 0
	minPathY, maxPathY = math.MaxInt, 0

	goesUp    = []rune{'|', 'L', 'J'}
	goesDown  = []rune{'|', '7', 'F'}
	goesRight = []rune{'-', 'L', 'F'}
	goesLeft  = []rune{'-', '7', 'J'}
)

func main() {
	parseInput()
	initGrids()
	findLoop()
	buildLoop()
}

func initGrids() {
	lvisited = make([][]*node, 0)
	dists = make([][]int, 0)
	visited = make([][]bool, 0)
	for range grid {
		row := make([]int, 0)
		rowb := make([]bool, 0)
		lvisited = append(lvisited, make([]*node, len(grid[0])))
		for range grid[0] {
			row = append(row, math.MaxInt)
			rowb = append(rowb, false)
		}
		dists = append(dists, row)
		visited = append(visited, rowb)
	}
	dists[startingX][startingY] = 0
}

func findLoop() {
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
		if minPathX > crds.x {
			minPathX = crds.x
		}
		if minPathY > crds.y {
			minPathY = crds.y
		}
		if maxPathX < crds.x {
			maxPathX = crds.x
		}
		if maxPathY < crds.y {
			maxPathY = crds.y
		}

		t := tile(crds.x, crds.y)
		if t == 'S' {
			options := make([][]rune, 0)
			if crds.x > 0 && Contains(tile(crds.x-1, crds.y), goesDown) {
				stk = append(stk, newCoord(crds.x-1, crds.y))
				dists[crds.x-1][crds.y] = 1
				options = append(options, goesUp)
			}
			if Contains(tile(crds.x+1, crds.y), goesUp) {
				stk = append(stk, newCoord(crds.x+1, crds.y))
				dists[crds.x+1][crds.y] = 1
				options = append(options, goesDown)
			}
			if Contains(tile(crds.x, crds.y+1), goesLeft) {
				stk = append(stk, newCoord(crds.x, crds.y+1))
				dists[crds.x][crds.y+1] = 1
				options = append(options, goesRight)
			}
			if crds.y > 0 && Contains(tile(crds.x, crds.y-1), goesRight) {
				stk = append(stk, newCoord(crds.x, crds.y-1))
				dists[crds.x][crds.y-1] = 1
				options = append(options, goesLeft)
			}

			for i := range options[0] {
				for j := range options[1] {
					if options[0][i] == options[1][j] {
						grid[crds.x][crds.y] = options[0][i]
					}
				}
			}
			visited[crds.x][crds.y] = true
			continue
		}

		if Contains(t, goesUp) && crds.x > 0 {
			if !visited[crds.x-1][crds.y] {
				stk = append(stk, newCoord(crds.x-1, crds.y))
			}
			if dists[crds.x-1][crds.y] > d+1 {
				dists[crds.x-1][crds.y] = d + 1
			}
		}
		if Contains(t, goesDown) && crds.x+1 < len(grid) {
			if !visited[crds.x+1][crds.y] {
				stk = append(stk, newCoord(crds.x+1, crds.y))
			}
			if dists[crds.x+1][crds.y] > d+1 {
				dists[crds.x+1][crds.y] = d + 1
			}
		}
		if Contains(t, goesRight) && crds.y+1 < len(grid[0]) {
			if !visited[crds.x][crds.y+1] {
				stk = append(stk, newCoord(crds.x, crds.y+1))
			}
			if dists[crds.x][crds.y+1] > d+1 {
				dists[crds.x][crds.y+1] = d + 1
			}
		}
		if Contains(t, goesLeft) && crds.y > 0 {
			if !visited[crds.x][crds.y-1] {
				stk = append(stk, newCoord(crds.x, crds.y-1))
			}
			if dists[crds.x][crds.y-1] > d+1 {
				dists[crds.x][crds.y-1] = d + 1
			}
		}
		visited[crds.x][crds.y] = true
	}
}

type node struct {
	x, y int
	next, prev *node
}

func (n *node) AddNext(other *node) {
	n.next = other
	other.prev = n
}

func buildLoop() {
	var (
		cur *node
		x,y int
	)
	startingNode = &node{x: startingX, y: startingY}
	cur = startingNode

	for {
		x, y = cur.x, cur.y
		if lvisited[x][y] != nil {
			break
		}

		lvisited[x][y] = cur
		t := tile(x, y)

		if Contains(t, goesUp) && x > 0 && lvisited[x-1][y] == nil {
			cur.AddNext(&node{x: x - 1, y: y})
		} else if Contains(t, goesDown) && x+1 < len(grid) && lvisited[x+1][y] == nil {
			cur.AddNext(&node{x: x + 1, y: y})
		} else if Contains(t, goesRight) && y+1 < len(grid[0]) && lvisited[x][y+1] == nil {
			cur.AddNext(&node{x: x, y: y + 1})
		} else if Contains(t, goesLeft) && y > 0 && lvisited[x][y-1] == nil {
			cur.AddNext(&node{x: x, y: y - 1})
		}
		if cur.next != nil {
			cur = cur.next
		}
	}
	cur.AddNext(startingNode)

	// We want to count ground tiles within the loop for each row
	// we start each row outside the loop and if we come across part of
	// the pipe we see what direction its going to keep track of whether
	// we are inside the loop or not
	totals := make(map[int]int)
	deltas := make([][]int, 0)
	for i := range grid {
		var (
			x int
			seen bool
		)
		deltas = append(deltas, make([]int, 0))
		for j := range grid[i] {
			n := lvisited[i][j]

			if n != nil {
				nxt, prev := n.next, n.prev
				// -2 to +2
				delta := nxt.x - prev.x

				if seen {
					x -= delta
				} else {
					seen = true
					x = delta * -1
				}
			} else {
				totals[x]++
			}
		}
	}
	fmt.Println("Total", totals)
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
