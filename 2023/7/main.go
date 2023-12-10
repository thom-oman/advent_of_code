package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	var hds []hand
	for _, l := range strings.Split(string(f), "\n") {
		if len(l) == 0 {
			continue
		}

		f := strings.Fields(l)
		b, _ := strconv.Atoi(f[1])

		h := newHand(f[0], b)
		hds = append(hds, h)
	}
	
	sort.Slice(hds, func(i,j int) bool {
		x,y := hds[i], hds[j]
		if x.ct == y.ct {
			return compareCards(x.cards, y.cards)
		}
		return x.ct < y.ct
	})

	var t int
	for i := range hds {
		t += (i+1) * hds[i].bid	
	}
	fmt.Println("Total", t)
}

func compareCards(a,b []byte) bool {
	for idx := range a {
		i,j := a[idx], b[idx]

		if i == j {
			continue
		}

		if isNumber(i) || isNumber(j) {
			// If both are numbers, simply compare the byte values as proxy
			if isNumber(i) && isNumber(j) {
				return i < j
			}

			// if only first is a number then other isn't, otherise the other is a number
			// whilst the first is not. The number always lose
			return isNumber(i)
		}

		return bytes.IndexByte(letterCardsSorted, i) < bytes.IndexByte(letterCardsSorted, j) 
	}

	// Identical, so choose one loser
	return false
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

// A = 65
// K = 75
// Q = 81
// J = 74
// T = 84
var (
	cardA byte = 65
	cardK byte = 75
	cardQ byte = 81
	cardJ byte = 74
	cardT byte = 84
)
var letterCardsSorted = []byte{cardA, cardK, cardQ, cardJ, cardT}

type hand struct {
	cards []byte
	bid int
	ct cardType
}

type cardType int

const (
	highCard cardType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func newHand(cards string, b int) hand {
	bs := []byte(cards)
	// var jids []int

	var h hand
	h.cards = bs
	h.bid = b

	cs := make(map[byte]int)
	for i := range bs {
		// if bs[i] == 74 {
		// 	jids = append(jids, i)
		// }
		cs[bs[i]]++
	}
	if len(cs) == len(bs) {
		h.ct = highCard
		return h
	} 

	if len(cs) == 1 {
		h.ct = fiveOfAKind
		return h
	}
	var ps, ts int
	for j := range cs {
		if cs[j] == 2 {
			ps++
		}
		if cs[j] == 3 {
			ts++
		}
		if cs[j] == 4 {
			h.ct = fourOfAKind
			return h
		}
	}

	if ts > 0 {
		if ps > 0 {
			h.ct = fullHouse
			return h
		}
		h.ct = threeOfAKind
		return h
	}

	if ps == 2 {
		h.ct = twoPair
		return h
	} else if ps == 1 {
		h.ct = onePair
		return h
	}

	return h
}
