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

	sort.SliceStable(hds, func(i, j int) bool {
		// This returns true if i should be ranked lower j
		x, y := hds[i], hds[j]

		if x.ct != y.ct {
			return x.ct < y.ct
		}
		return compareCards(x, y)
	})

	var t int
	for i, h := range hds {
		t += (i + 1) * h.bid
	}
	fmt.Println("Total", t)
}

func compareCards(a, b hand) bool {
	for idx := range a.cards {
		i, j := a.cards[idx], b.cards[idx]

		if i != j {
			return compareValue(i, j)
		}
	}

	panic("cant split")
}

func compareValue(i, j byte) bool {
	// If either card is J then we don't have to check further
	if i == cardJ || j == cardJ {
		return i == cardJ
	}

	// If both are numbers, compare the byte values as proxy
	if isNumber(i) && isNumber(j) {
		return i < j
	}

	// if only first is a number then other isn't or vice versa, the the number loses
	if isNumber(i) || isNumber(j) {
		return isNumber(i)
	}

	ii := bytes.IndexByte(letterCardsSorted, i)
	ij := bytes.IndexByte(letterCardsSorted, j)
	if ii < 0 || ij < 0 {
		panic(fmt.Sprintf("%v\t%v", ii, ij))
	}
	return ii < ij
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}

var (
	cardA byte = 65
	cardK byte = 75
	cardQ byte = 81
	cardJ byte = 74
	cardT byte = 84
)
var letterCardsSorted = []byte{cardT, cardQ, cardK, cardA}

type hand struct {
	cards []byte
	bid   int
	ct    cardType
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

	return hand{
		cards: bs,
		bid:   b,
		ct:    determineCardType(bs),
	}
}

func determineCardType(bs []byte) cardType {
	ccount := make(map[byte]int)

	for i := range bs {
		ccount[bs[i]]++
	}

	jc, _ := ccount[cardJ]
	startType := cardTypeForCounts(ccount)
	return improvedCardType(startType, jc)
}

func improvedCardType(startType cardType, jc int) cardType {
	// No jacks/jokers or hand cannot be improved
	if jc == 0 || startType == fiveOfAKind {
		return startType
	}

	switch startType {
	case fourOfAKind:
		// Change joker to match other 4 or change all
		// 4 jokers to match 5th card
		if jc == 1 || jc == 4 {
			return fiveOfAKind
		}
	case fullHouse:
		// Similarly, change minoriy to match majority
		if jc == 2 || jc == 3 {
			return fiveOfAKind
		}
	case threeOfAKind:
		// If the 3 are jokers, then no improvement can be made
		if jc == 3 {
			return fourOfAKind
		}
		// swap joker to match other 3
		if jc == 1 {
			return fourOfAKind
		}
	case twoPair:
		// One of the pairs is jokers so make them match other pair
		if jc == 2 {
			return fourOfAKind
		}

		// Turn one of the pairs into a three
		if jc == 1 {
			return fullHouse
		}
	case onePair:
		// If pair is jacks then no improvement can be made
		if jc == 2 {
			return threeOfAKind
		}
		// Otherwise, if one joker then we can turn pair into 3
		if jc == 1 {
			return threeOfAKind
		}
	case highCard:
		// Can only create a pair with one of the others
		if jc == 1 {
			return onePair
		}
	}
	return startType
}

func cardTypeForCounts(ccount map[byte]int) cardType {
	var t cardType
	var cs []int

	for _, c := range ccount {
		cs = append(cs, c)

		switch c {
		case 5:
			if t == highCard {
				t = fiveOfAKind
			}
		case 4:
			if t == highCard {
				t = fourOfAKind
			}
		case 3:
			if t == onePair {
				t = fullHouse
			} else if t == highCard {
				t = threeOfAKind
			}
		case 2:
			if t == threeOfAKind {
				t = fullHouse
			} else if t == onePair {
				t = twoPair
			} else if t == highCard {
				t = onePair
			}
		}
	}
	return t
}

