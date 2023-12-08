package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strconv"
)

var dot = byte(46)

func main() {
	f, err := os.Open("inputs.txt")

	if err != nil {
		panic(err)
	}
	g := &grid{}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadBytes('\n')

		if err == io.EOF || len(line) == 0 {
			break
		}

		g.addLine(line[:len(line)-1])
	}

	var sum int

	for i, l := range g.lines {
		for j, it := range l.contents {
			n, ok := it.(number)
			if !ok {
				continue
			}

			coords := n.Coordinates()

			// Check previous item in line
			if j > 0 {
				prev := l.contents[j-1]

				if _, ok := prev.(symbol); ok {
					if coords[0] == prev.Coordinates()[1]+1 {
						sum += n.val
						continue
					}
				}

			}

			// Check next symbol in line
			if j < len(l.contents)-1 {
				next := l.contents[j+1]

				if _, ok := next.(symbol); ok {
					if coords[1]+1 == next.Coordinates()[0] {
						// BOOM
						sum += n.val
						continue
					}
				}

			}

			// Check previous line for symbol
			if i > 0 {
				prevLine := g.lines[i-1]

				for _, plit := range prevLine.contents {
					if _, ok := plit.(symbol); ok {
						point := plit.Coordinates()[0]

						if point >= coords[0]-1 && point <= coords[1]+1 {
							// BOOM
							sum += n.val
							continue
						}
					}
				}
			}

			// Check next line for symbol
			if i < len(g.lines) - 1{
				nextLine := g.lines[i+1]

				for _, nlit := range nextLine.contents {
					if _, ok := nlit.(symbol); ok {
						point := nlit.Coordinates()[0]

						if point >= coords[0]-1 && point <= coords[1]+1 {
							// Boom
							sum += n.val
							continue
						}
					}

				}
			}
		}
	}

	fmt.Println("Sum:", sum)
}

type grid struct {
	lines []gridLine
}

func (g *grid) addLine(bs []byte) {
	l := gridLine{}
	var i int

	for i < len(bs) {
		if bs[i] == dot {
			i++
			continue
		}

		var gi gridItem
		if isNumber(bs[i]) {
			j := i
			for j < len(bs) && isNumber(bs[j]) {
				j++
			}

			n, err := strconv.Atoi(string(bs[i:j]))
			if err != nil {
				panic(err)
			}

			gi = number{val: n, start: i, end: j - 1}
			i = j
		} else {
			gi = symbol{val: string(bs[i]), point: i}
			i++
		}

		l.addItem(gi)
	}

	g.lines = append(g.lines, l)
}

type gridLine struct {
	idx      int
	contents []gridItem
}

func (gl *gridLine) addItem(gi gridItem) {
	gl.contents = append(gl.contents, gi)
}

type gridItem interface {
	Coordinates() [2]int
	Value() string
}

type number struct {
	start, end int
	val        int
}

func (n number) Coordinates() [2]int {
	return [2]int{n.start, n.end}
}

func (n number) Value() string {
	return fmt.Sprint(n.val)
}

type symbol struct {
	point int
	val   string
}

func (s symbol) Coordinates() [2]int {
	return [2]int{s.point, s.point}
}

func (s symbol) Value() string {
	return s.val
}

func isNumber(b byte) bool {
	return b >= 48 && b <= 57
}
