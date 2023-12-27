package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	parseInput()
}

func hash(s string) float64 {
	var cur float64
	for i := range s {
		cur += float64(s[i])
		cur *= 17
		cur = math.Mod(cur, 256)
	}
	fmt.Println(s, []byte(s), cur)
	return cur
}

func parseInput() {
	var sum float64 
	bs, _ := os.ReadFile("inputs.txt")
	var line string
	for _, l := range strings.Split(string(bs), "\n") {
		if len(l) > 0 {
			line = l
		}
	}
	for _, in := range strings.Split(line, ",") {
		sum += hash(in)
	}
	fmt.Println(sum)
}
