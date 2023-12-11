package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	steps int
	route string
	currentNodes []*node
	letters = regexp.MustCompile("[[:alpha:]]+")
	firstZs = make(map[int]int)
	nodes = make(map[string]*node)
)

func main() {
	parseFile()

	iterateAllInstructions()

	var ns []int
	for _, c := range firstZs {
		ns = append(ns, c)
	}

	switch len(ns) {
	case 1:
		fmt.Println(ns[0])
	case 2:
		fmt.Println(LCM(ns[0], ns[1]))
	default:
		fmt.Println(LCM(ns[0], ns[1], ns[2:]...))
	}
}

func parseFile() {
	i, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	for _, l := range strings.Split(string(i), "\n") {
		if len(l) == 0 {
			continue
		}

		lts := letters.FindAllString(l, 3)
		if len(lts) == 1 {
			route = lts[0]
			continue
		}

		key, vals  := lts[0], lts[1:3]
		var (
			ok bool
			n, l, r *node
		)

		if n, ok = nodes[key]; !ok {
			n = &node{val: key}
			nodes[key] = n
		}

		if l, ok = nodes[vals[0]]; !ok {
			l = &node{val: vals[0]}
			nodes[vals[0]] = l
		}
		n.left = l

		if r, ok = nodes[vals[1]]; !ok {
			r = &node{val: vals[1]}
			nodes[vals[1]] = r
		}
		n.right = r

		if strings.HasSuffix(key, "A") {
			currentNodes = append(currentNodes, n)
		}
	}
}

type node struct {
	val string
	left *node
	right *node
}

// Copied and pasted this from the internet because my implementation
// sucked and wasn't working
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

// func lcm(ns... int) int {
// 	m := 1
// 	for i := range ns {
// 		m *= ns[i]
// 	}
// 	fmt.Println("lcm", m)
// 	return m / gcd(ns...)
// }

// func gcd(ns... int) int {
// 	if len(ns) <= 2 {
// 		if len(ns) == 1 {
// 			return 1
// 		}

// 		a, b := ns[0], ns[1]
// 		for b != 0 {
// 			t := a
// 			b = a % b
// 			a = t
// 		}

// 		return a
// 	}
// 	return gcd(ns[0], gcd(ns[1:]...))
// }

func iterateAllInstructions() {
	for i := 0; i < len(currentNodes); i++ {
		iterate(i)
	}
}

func iterate(i int) int {
	var n int

	for {
		for _, d := range route {
			switch string(d) {
			case "R":
				currentNodes[i] = currentNodes[i].right
			case "L":
				currentNodes[i] = currentNodes[i].left
			}
			n++

			if strings.HasSuffix(currentNodes[i].val, "Z") {
				if _, ok := firstZs[i]; !ok {
					firstZs[i] = n
					return n
				}
			}
		}
	}
}
