package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs.txt")

	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)

	var total int
	for {
		bs, err := r.ReadBytes('\n')
		if len(bs) == 0 {
			break
		}

		if err != nil && err != io.EOF {
			panic(err)
		}

		g := newGame(bs)
		total += g.Score()
	}
	fmt.Println("score", total)
}

type game struct {
	number int
	winning []int
	candidates []int
}

func (g game) Score() int {
	var matches int

	for i := range g.candidates {
		for j := range g.winning {
			if g.candidates[i] == g.winning[j] {
				matches++
			}
		}
	}
	if matches == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(matches - 1)))
}

func newGame(bs []byte) game {
	var (
		numberFound, winningFound bool
		number int
		winning, candidates []int
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
				
				number, _ = strconv.Atoi(string(ns))
				i = j
				numberFound = true
			} else if !winningFound {
				var cur []byte

				j := iterateWhile(
					bs,
					func(b byte)  bool { return b != 124 },
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
					func(b byte)  bool { return b != 10 },
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
		number: number,
		winning: winning,
		candidates: candidates,
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
