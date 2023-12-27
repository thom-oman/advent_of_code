package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	grid [][]rune
	visited [][]bool
)

func main() {
	parseInputs()
	walkGrid()
	var tot int
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				tot++
			}
		}
	}
	fmt.Println(tot)
}

type direction int

const (
	Up direction = iota
	Right
	Down
	Left
)

type item struct {
	x, y int
	dir  direction
}

func newItem(x, y int, dir direction) item {
	return item{x, y, dir}
}

func walkGrid() {
	var (
		stk = []item{newItem(0, 0, Right)}
		cur item
		processed = make(map[item]bool)
	)

	for len(stk) > 0 {
		cur, stk = stk[0], stk[1:]
		if cur.x < 0 || cur.y < 0 || cur.x >= len(grid) || cur.y >= len(grid[0]) {
			continue
		} else if _, ok := processed[cur]; ok {
			continue
		}
		processed[cur] = true
		visited[cur.x][cur.y] = true

		newDir := cur.dir
		switch grid[cur.x][cur.y] {
		case '.':
			x,y := nextCoords(cur.x, cur.y, cur.dir)
			stk = append(stk, newItem(x,y, newDir))
		case '\\':
			switch cur.dir {
			case Up:
				newDir = Left
			case Down:
				newDir = Right
			case Left:
				newDir = Up
			case Right:
				newDir = Down
			}
			x,y := nextCoords(cur.x, cur.y, newDir)
			stk = append(stk, newItem(x,y, newDir))
		case '/':
			switch cur.dir {
			case Up:
				newDir = Right
			case Down:
				newDir = Left
			case Left:
				newDir = Down
			case Right:
				newDir = Up
			}
			x,y := nextCoords(cur.x, cur.y, newDir)
			stk = append(stk, newItem(x,y, newDir))
		case '-':
			switch cur.dir {
			case Up, Down:
				x,y := nextCoords(cur.x, cur.y, Right)
				stk = append(stk, newItem(x,y, Right))
				x,y = nextCoords(cur.x, cur.y, Left)
				stk = append(stk, newItem(x,y, Left))
			default:
				x,y := nextCoords(cur.x, cur.y, newDir)
				stk = append(stk, newItem(x,y, newDir))
			}
		case '|':
			switch cur.dir {
			case Left, Right:
				x,y := nextCoords(cur.x, cur.y, Up)
				stk = append(stk, newItem(x,y, Up))
				x,y = nextCoords(cur.x, cur.y, Down)
				stk = append(stk, newItem(x,y, Down))
			default:
				x,y := nextCoords(cur.x, cur.y, newDir)
				stk = append(stk, newItem(x,y, newDir))
			}
		default:
			panic("sdfdfg")
		}
	}
}

func nextCoords(i,j int, dir direction) (int, int) {
	switch dir {
	case Up:
		return i - 1, j
	case Down:
		return i + 1, j
	case Right:
		return i, j + 1
	case Left:
		return i, j - 1
	default:
		panic("invalid direction")
	}
}

func parseInputs() {
	bs, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(bs), "\n") {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
		visited = append(visited, make([]bool, len(line)))
	}
}
