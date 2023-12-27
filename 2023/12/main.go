package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	rows     []*row
	allPerms []map[string]bool
)

func main() {
	parseInputFile()
	processRows()
}

func processRows() {
	var tot int
	for _, r := range rows {
		allPerms = append(allPerms, make(map[string]bool))
		fmt.Println("[START: PROCESSROW]", string(r.springs), r.sizes)
		tot += permutations(r.springs, r.sizes)
		fmt.Println("[END: PROCESSROW]", tot)
	}
	fmt.Println("total:", tot)
}

func lastPerms() map[string]bool {
	return allPerms[len(allPerms)-1]
}

func permutations(s string, sizes []int) int {
	var (
		perms int
		ap = lastPerms()
		// Find where first unknown is
		idx = strings.Index(s, "?")
	)

	if idx == -1 {
		if isValid(s, sizes) {
			return 1
		} else {
			return 0 
		}
	}

	// Create a string for each variant
	p1 := s[:idx] + "." + s[idx+1:]
	p2 := s[:idx] + "#" + s[idx+1:]

	if _, ok := ap[p1]; !ok {
		if ap[p1] = possibleToMatch(p1, sizes); ap[p1] {
			strIdx, sizesMatch := matchSizes(p1, sizes)
		}
	}

	if _, ok := ap[p2]; !ok {
		if ap[p2] = possibleToMatch(p2, sizes); ap[p2] {
		}
	}

	return perms
}

type walker struct {
	s string
	cur, nextCur int
	remaining string
	n int
	unknowns, damaged int
}

func (w *walker) moveCursorToPossibleDamaged() {
	w.cur = nextPossibleDamaged(w.s, w.cur)
}

func (w *walker) determineRemaining() bool {
	w.remaining = w.s[w.cur:]
	w.nextCur = w.cur + strings.Index(w.remaining, ".")
	if w.nextCur == w.cur-1 && len(w.remaining) < w.n {
		return false
	}
	if w.nextCur != w.cur -1 {
		w.remaining = w.s[w.cur:w.nextCur]
	}
	return true
}

func (w *walker) countRemaining() {
	w.damaged = strings.Count(w.remaining, "#")
	w.unknowns = strings.Count(w.remaining, "?")
	// fmt.Printf(
	// 	"str=%v rem=%v size=(%v) dam=%v unk=%v\n", 
	// 	w.s, w.remaining, w.n, w.damaged, w.unknowns,
	// )
}

func possibleToMatch(s string, sizes []int) bool {
	// Start at first non . character
	// It's vital that this index is relative to original string

	// Handle special case of 100% unknowns
	if strings.Count(s, "?") == len(s) {
		var minPos int
		for i := range sizes {
			minPos += sizes[i] + 1
		}
		return minPos-1 <= len(s)
	}

	w := walker{s: s}
	w.moveCursorToPossibleDamaged()
	for i, n := range sizes {
		w.n = n

		// Find index where the next KNOWN .
		if !w.determineRemaining() {
			return false
		}

		w.countRemaining()

		// If there's definitely an island here and it matches our count
		// we can prove or disprove a match
		if w.damaged > 0 {
			if w.damaged+w.unknowns == w.n {
				w.moveCursorToPossibleDamaged()
				
				if w.cur == -1 {
					// No more damaged so either this is the last size to match and
					// we have successfully matched or there are more but no more damaged left
					return i == len(s)-1
				}
				continue
			}

			// We know there must be an island but the maximum size is less than we need
			if w.damaged+w.unknowns< n { return false }
			if w.damaged > n { return false }

			if strings.Count(w.remaining[w.n:], "#") == 0 {
				w.cur += n + 1
				continue
			}

			w.moveCursorToPossibleDamaged()
			if w.cur == -1 {
				// fmt.Println(w.cur, len(s)-1)
				return i == len(s)-1
			}
		} else { // No damaged
			if w.unknowns < n {
				// If we don't have enough remaining and this is the last
				// island, then we fail, otherwise we can assume all ? are .
				if w.nextCur == -1 {
					return false
				}

				// Assume unknowns are ., look to next island
				w.moveCursorToPossibleDamaged()
				if w.cur == -1 {
					return false
				}

				// TODO
				// remaining = s[cur:]
				// nextCur = cur + strings.Index(remaining, ".")
				// // No more . remaining, so we can check that we have enough
				// // characters left
				// if nextCur == cur-1 && len(remaining) < n {
				// 	return false
				// } 

				// // This excludes the next . if there is one
				// if nextCur != cur-1 {
				// 	remaining = s[cur:nextCur]
				// }
			}

			if w.unknowns > n {
				if strings.Count(w.remaining[n:], "#") == 0 {
					w.cur += n + 1
					continue
				}
			}

			if w.nextCur == -1 {
				return i == len(s)-1
			}

			// cur = nextPossibleDamaged(s, nextCur)
			w.moveCursorToPossibleDamaged()
			if w.cur == -1 {
				return i == len(s)-1
			}
		}
	}

	return true
}

func nextPossibleDamaged(s string, cur int) int {
	a, b := strings.Index(s[cur:], "#"), strings.Index(s[cur:], "?")
	if a == -1 && b == -1 {
		return -1
	} else if a == -1 || b == -1 {
		return cur + maxInt(a, b)
	} else {
		return cur + minInt(a, b)
	}
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func matchSizes(s string, sizes []int) (int, int) {
	var idx, matched int

	for {
		r := s[idx:]
		i := strings.Index(r, ".")
		j := strings.Index(r, "#")
		k := strings.Index(r, "?")
		if k < j && k < i {
			return idx, matched
		}
	}

	for _, l := range strings.Split(s, ".") {
		if len(l) == 0 {
			continue
		}

		if strings.ContainsAny(l, "?") {
			return matched
		}

		matched++
	}

	return matched
}

func isValid(s string, sizes []int) bool {
	if strings.Count(s, "?") != 0 {
		return false
	}
	
	var i int

	for _, b := range strings.Split(s, ".") {
		if len(b) > 0 {
			if len(b) != sizes[i] {
				return false
			}
			i++
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

func unknownsInBlock(b string) int {
	var c int
	for i := range b {
		if b[i] == '?' {
			c++
		}
	}
	return c
}

type block struct {
	contents []rune
	unknowns int
	next     *block
}

func (b block) String() string {
	return string(b.contents)
}

func newBlock(rs []rune) *block {
	return &block{contents: rs, unknowns: unknownsInBlock(string(rs))}
}

func splitIntoBlocks(s []rune) []*block {
	var (
		cur    = make([]rune, 0)
		blocks = make([]*block, 0)
	)

	for i := range s {
		if s[i] != '.' {
			cur = append(cur, s[i])
			continue
		}

		if len(cur) > 0 {
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

type row struct {
	springs string
	sizes   []int
	blocks  []*block
}

func newRow(ss []rune, sz []int) *row {
	return &row{
		springs: string(ss),
		sizes:   sz,
		blocks:  splitIntoBlocks(ss),
	}
}

func parseInputFile() {
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

		row := newRow(springs, sizes)
		// row := newRow(duplicateSprings(springs), duplicateSizes(sizes))
		rows = append(rows, row)
	}
}

func duplicateSprings(rs []rune) []rune {
	var duplicated []rune
	for i := 1; i <= 5; i++ {
		duplicated = append(duplicated, rs...)
		if i != 5 {
			duplicated = append(duplicated, '?')
		}
	}
	return duplicated
}

func duplicateSizes(rs []int) []int {
	var duplicated []int
	for i := 1; i <= 5; i++ {
		duplicated = append(duplicated, rs...)
	}
	return duplicated
}
