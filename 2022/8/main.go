package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var input string

func main() {
	tree := strings.Split(input, "\n")
	dim_x := len(tree)
	dim_y := len(tree[0])

	visible := make([][]bool, dim_x)
	// Set everything as not visible
	for row_idx := range visible {
		visible[row_idx] = make([]bool, dim_y)

		for col_idx := range visible[row_idx] {
			visible[row_idx][col_idx] = false
		}
	}

	// Track largest values we see from top and left
	var biggestHorizontally int64
	biggestVertically := make([]int64, dim_y)
	for col_idx := range biggestVertically {
		biggestVertically[col_idx] = -1
	}

	for row_idx := 0; row_idx <= len(tree)-1; row_idx++ {
		biggestHorizontally = -1

		for col_idx := range tree[row_idx] {
			n, _ := strconv.ParseInt(string(tree[row_idx][col_idx]), 10, 0)

			if n > biggestVertically[col_idx] || n > biggestHorizontally {
				visible[row_idx][col_idx] = true
			}
			if n > biggestVertically[col_idx] {
				biggestVertically[col_idx] = n
			}
			if n > biggestHorizontally {
				biggestHorizontally = n
			}
		}
	}

	// ...then from the bottom and right
	for col_idx := range biggestVertically {
		biggestVertically[col_idx] = -1
	}
	for row_idx := len(tree) - 1; row_idx >= 0; row_idx-- {
		biggestHorizontally = -1

		for col_idx := len(tree[row_idx]) - 1; col_idx >= 0; col_idx-- {
			n, _ := strconv.ParseInt(string(tree[row_idx][col_idx]), 10, 0)
			if n > biggestVertically[col_idx] || n > biggestHorizontally {
				visible[row_idx][col_idx] = true
			}
			if n > biggestVertically[col_idx] {
				biggestVertically[col_idx] = n
			}
			if n > biggestHorizontally {
				biggestHorizontally = n
			}
		}
	}

	c := 0
	for row_idx := range tree {
		for col_idx := range tree[row_idx] {
			if visible[row_idx][col_idx] {
				c++
			}
		}
	}
	fmt.Println(c)
}
