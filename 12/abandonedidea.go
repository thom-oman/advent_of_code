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
	for i := range rows {
		fmt.Println(i, string(rows[i].springs), rows[i].sizes)
		processRow(rows[i])
	}
}

func processRow(r row) {
	var (
		blocks = splitIntoBlocks(r.springs)
		blockToSizeIdx = make(map[*block][]int)
		matchedSizes = make([]bool, len(r.sizes))
	)

	// If we have the same amount then they must match 1-1
	if len(blocks) == len(r.sizes) {
		for j := range blocks {
			matchedSizes[j] = true
			blockToSizeIdx[blocks[j]] = []int{j}
		}
	}

	// If any have no unknowns then we must be able to match it
	for i := range blocks {
		b := blocks[i]
		if _, ok := blockToSizeIdx[b]; ok {
			continue
		}

		if unknownsInCurrentBlock(b.contents) != 0 {
			continue
		}

		for j := range r.sizes {
			if len(b.contents) == r.sizes[j] && !matchedSizes[j] {
				matchedSizes[j] = true
				blockToSizeIdx[b] = []int{i}
			}
		}
	}

	maxIdx := 0
	b := blocks[0]

	// We want to go block by block and find what sizes we need to make
	// accounting for all the unknowns.
	for b != nil {
		if _, ok := blockToSizeIdx[b]; ok {
			if len(blockToSizeIdx[b]) != 1 {
				panic("WTF")
			}
			maxIdx = blockToSizeIdx[b][0]
			b = b.next
			continue
		}

		targetLen := len(b.contents)
		
		for i := range matchedSizes {
			if i < maxIdx {
				continue
			}
			if matchedSizes[i] {
				maxIdx = i
				continue
			}



		}

		var sizes []int
		for sumInts(sizes) < targetLen {
			sizes = append(sizes, r.sizes[i])
			fmt.Println("loop", i, "target", targetLen, sizes, r.sizes, "Sum:", sumInts(sizes))
			i++
			if i == len(r.sizes)-1 {
				break
			}
		}
		blockToSizeIdx[b] = sizes

		fmt.Println(sizes)
		b = b.next
	}

	fmt.Println(blockToSizeIdx)
	fmt.Println("")
}

func sumInts(ns []int) int {
	var c int
	for i := range ns {
		c += ns[i]
	}
	return c
}

func unknownsInCurrentBlock(b []rune) int {
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
	return &block{contents: rs, unknowns: unknownsInCurrentBlock(rs)}
}

type row struct {
	springs []rune
	sizes   []int
}

func parseInputFile() {
	// bs, err := os.ReadFile("inputs.txt")
	bs, err := os.ReadFile("testinputs.txt")
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
