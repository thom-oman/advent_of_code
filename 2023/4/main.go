package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
)

type stackitem struct {
	idx      int
	original bool
}

func main() {
	f, err := os.Open("inputs.txt")

	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	var games []game

	for {
		bs, err := r.ReadBytes('\n')
		if len(bs) == 0 {
			break
		}

		if err != nil && err != io.EOF {
			panic(err)
		}

		games = append(games, newGame(bs))
	}

	var total int
	copies := make(map[int]int)

	for i := 0; i < len(games); i++ {
		copies[i]++
		total += copies[i]

		m := games[i].Matches()
		fmt.Printf("game: %v\tmatches: %v\tcard count: %v\t\ttotal: %v\n", i, m, copies[i], total)

		if m == 0 {
			continue
		}

		for j := 1; j <= m; j++ {
			copies[i+j] += copies[i]
		}
	}

	fmt.Println("Total", total)
}

type game struct {
	number     string
	winning    []int
	candidates []int
	matches    int
}

func (g game) Matches() int {
	var matches int

	for i := range g.winning {
		if slices.Contains(g.candidates, g.winning[i]) {
			matches++
		}
	}

	return matches
}

func (g game) Score() int {
	m := g.Matches()
	if m == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(m-1)))
}

func newGame(bs []byte) game {
	var (
		numberFound, winningFound bool
		number                    string
		winning, candidates       []int
	)

	var i int
	for i < len(bs) {
		if isNumber(bs[i]) {
			if !numberFound {
				var ns []byte
				j := iterateWhile(
					bs,
					isNumber,
					func(b byte) { ns = append(ns, b) },
					i,
				)

				number = string(ns)
				i = j
				numberFound = true
			} else if !winningFound {
				var cur []byte

				j := iterateWhile(
					bs,
					func(b byte) bool { return b != 124 },
					func(b byte) {
						if b != 32 {
							cur = append(cur, b)
							return
						}

						if len(cur) > 0 {
							n, _ := strconv.Atoi(string(cur))
							winning = append(winning, n)
							cur = make([]byte, 0)
						}
					},
					i,
				)

				i = j
				winningFound = true
			} else {
				var cur []byte
				j := iterateWhile(
					bs,
					func(b byte) bool { return b != 10 },
					func(b byte) {
						if b != 32 {
							cur = append(cur, b)
							return
						}

						if len(cur) > 0 {
							n, _ := strconv.Atoi(string(cur))
							candidates = append(candidates, n)
							cur = make([]byte, 0)
						}
					},
					i,
				)
				if len(cur) > 0 {
					n, _ := strconv.Atoi(string(cur))
					candidates = append(candidates, n)
				}

				i = j

			}
		}

		i++
	}
	return game{
		number:     number,
		winning:    winning,
		candidates: candidates,
		matches:    -1,
	}
}

func iterateWhile(bs []byte, f func(byte) bool, x func(byte), i int) int {
	var j int
	for j = i; j < len(bs) && f(bs[j]); j++ {
		x(bs[j])
	}
	return j
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}
