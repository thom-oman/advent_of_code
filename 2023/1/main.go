package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var input string

// Part 1
// func main() {
// 	var (
// 		b    byte
// 		i, s int
// 	)
// 	for i < len(input) {
// 		bs := []byte{}
// 		b = input[i]
// 		if b == 10 {
// 			i++
// 			continue
// 		}
// 		for b != 10 {
// 			if b >= 48 && b <= 57 {
// 				bs = append(bs, b)
// 			}

// 			i++
// 			b = input[i]
// 		}
// 		bs2 := []byte{bs[0], bs[len(bs)-1]}
// 		x, err := strconv.Atoi(string(bs2))
// 		if err != nil {
// 			panic(err)
// 		}
// 		s += x
// 		i++
// 	}
// 	fmt.Println("SUM:", s)
// }

// Part 2

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{root: newNode(byte(0))}
}

func (t *trie) addBytes(bs []byte) {
	cur := t.root

	for i := range bs {
		cur = cur.ChildFor(bs[i])
	}
}

func (t *trie) Candidate(s string) (bool, int) {
	bs := []byte(s)
	c := t.root
	var count int

	for i := range bs {
		found := false
		for j := range c.children {
			if found {
				continue
			}

			if c.children[j].val == bs[i] {
				found = true
				c = c.children[j]
				count++
			}
		}
		if !found {
			return false, count
		}
	}

	return true, count
}

type node struct {
	val      byte
	children []*node
}

func (n *node) ChildFor(b byte) *node {
	for i := range n.children {
		if n.children[i].val == b {
			return n.children[i]
		}
	}

	nn := newNode(b)
	n.children = append(n.children, nn)
	return nn
}

func newNode(b byte) *node {
	return &node{
		val:      b,
		children: make([]*node, 0),
	}
}

var stringsTrie = newTrie()
var rstringsTrie = newTrie()
var strToI = make(map[string]int)
var rstrToI = make(map[string]int)

func init() {
	numberStrings := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i, str := range numberStrings {
		strToI[str] = i + 1
		bs := []byte(str)
		rbs := make([]byte, len(bs))
		for j := range bs {
			rbs[len(bs)-1-j] = bs[j]
		}
		rstrToI[string(rbs)] = i + 1
		stringsTrie.addBytes(bs)
		rstringsTrie.addBytes(rbs)
	}
}

func main() {
	var s int

	for _, str := range strings.Split(input, "\n") {
		if len(str) == 0 {
			continue
		}

		s += 10*firstDigit(str) + lastDigit(str)
	}

	fmt.Println("Sum:", s)
}

func firstDigit(str string) int {
	for i := range str {
		if isNumber(str[i]) {
			n, _ := strconv.Atoi(string(str[i]))
			return n
		}

		if can, _ := stringsTrie.Candidate(string(str[i])); can {
			j := i + 1
			for {
				if isNumber(str[j]) {
					if n, ok := strToI[str[i:j-1]]; ok {
						return n
					}
					break
				}

				if can, _ = stringsTrie.Candidate(str[i:j]); !can {
					j--
					break
				}

				j++
			}

			if n, ok := strToI[str[i:j]]; ok {
				return n
			}
		}
	}

	panic("First digit not found")
}

func lastDigit(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if isNumber(str[i]) {
			n, _ := strconv.Atoi(string(str[i]))
			return n
		}

		if can, _ := rstringsTrie.Candidate(string(str[i])); can {
			j := i - 1
			for {
				if isNumber(str[j]) {
					j++
					break
				}

				trev := str[j : i+1]
				rev := make([]byte, len(trev))
				for k := range trev {
					rev[len(trev)-1-k] = trev[k]
				}
				if can, _ = rstringsTrie.Candidate(string(rev)); !can {
					j++
					break
				}

				j--
			}

			if n, ok := strToI[str[j:i+1]]; ok {
				return n
			}
		}
	}

	panic("Last digit not found")
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}
