package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	rows []row
)

func main() {
	parseInputFile()
	processRows()
}

func processRows() {
	var tot int
	for i := range rows {
		fmt.Println(i, string(rows[i].springs), rows[i].sizes)
		tot += processRow(rows[i])
	}
	fmt.Println("total:", tot)
}

func processRow(r row) int {
	var (
		processed = make(map[string]bool)
		x []rune
		stk [][]rune
		perms int
	)
	stk = append(stk, r.springs)

	for len(stk) > 0 {
		x, stk = stk[0], stk[1:]

		if processed[string(x)] {
			fmt.Println("Already processed")
			continue
		}

		if valid(x, r.sizes) {
			processed[string(x)] = true
			perms++
			continue
		} else if unknownsInBlock(x) == 0 {
			continue
		}

		permuted := permuteFirstUnknown(x)
		stk = append(stk, permuted...)
	}
	// fmt.Println(processed)
	return perms
}

func permuteFirstUnknown(rs []rune) [][]rune {
	var (
		idx int
		found bool
	)

	for i := range rs {
		if rs[i] == '?' {
			idx = i
			found = true
			break
		}
	}

	if !found {
		panic("NOT FOUND")
	}

	var (
		perm1 = make([]rune, len(rs))
		perm2 = make([]rune, len(rs))
	)
	copy(perm1, rs)
	copy(perm2, rs)
	perm1[idx] = '.'
	perm2[idx] = '#'

	return [][]rune{perm1, perm2}
}

func valid(rs []rune, sizes []int) bool {
	if unknownsInBlock(rs) != 0 {
		return false
	}

	blks := splitIntoBlocks(rs)
	if len(blks) != len(sizes) {
		return false
	}

	for i := range blks {
		if len(blks[i].contents) != sizes[i] {
			return false
		}
	}

	return true
}

func sumInts(ns []int) int {
	var c int
	for i := range ns {
		c += ns[i]
	}
	return c
}

func unknownsInBlock(b []rune) int {
	var c int
	for i := range b {
		if b[i] == '?' {
			c++
		}
	}
	return c
}

func splitIntoBlocks(s []rune) []*block {
	var (
		cur    = make([]rune, 0)
		blocks = make([]*block, 0)
	)

	for i := range s {
		if s[i] != '.' {
			cur = append(cur, s[i])
		} else if len(cur) > 0 {
			b := newBlock(cur)
			if len(blocks) > 0 {
				blocks[len(blocks)-1].next = b
			}
			blocks = append(blocks, b)
			cur = make([]rune, 0)
		}
	}

	if len(cur) > 0 {
		b := newBlock(cur)
		if len(blocks) > 0 {
			blocks[len(blocks)-1].next = b
		}
		blocks = append(blocks, b)
	}

	return blocks
}

type block struct {
	contents []rune
	unknowns int
	next     *block
}

func newBlock(rs []rune) *block {
	return &block{contents: rs, unknowns: unknownsInBlock(rs)}
}

type row struct {
	springs []rune
	sizes   []int
}

func parseInputFile() {
	bs, err := os.ReadFile("inputs.txt")
	// bs, err := os.ReadFile("testinputs.txt")
	if err != nil {
		panic(err)
	}

	for _, l := range strings.Split(string(bs), "\n") {
		if len(l) == 0 {
			continue
		}

		prsd := strings.Fields(l)

		var springs []rune
		for i := range prsd[0] {
			springs = append(springs, rune(prsd[0][i]))
		}

		var sizes []int
		for _, n := range strings.Split(prsd[1], ",") {
			s, _ := strconv.Atoi(n)
			sizes = append(sizes, s)
		}

		rows = append(rows, row{springs: springs, sizes: sizes})
	}
}
