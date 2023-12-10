package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	i, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	letters := regexp.MustCompile("[[:alpha:]]+")
	var instr string
	mappings := make(map[string][2]string)
	for _, l := range strings.Split(string(i), "\n") {
		if len(l) == 0 {
			continue
		}

		s := strings.Split(l, "=")
		if len(s) == 1 {
			instr = s[0]
			continue
		}

		x := letters.FindAllString(l, 3)
		mappings[x[0]] = [2]string{x[1], x[2]}
	}

	cur := "AAA"

	var steps int
	for cur != "ZZZ" {
		for _, d := range instr {
			if string(d) == "L" {
				cur = mappings[cur][0]
			} else if string(d) == "R" {
				cur = mappings[cur][1]
			} else {
				panic("sdfs")
			}
			steps++
		}
	}

	fmt.Println("Steps:", steps)
}
