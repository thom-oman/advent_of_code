package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed inputs.txt
var input string

func main() {
	for i := 13; i < len(input)-1; i++ {
		prev_4 := input[i-13 : i+1]
		seen := make(map[rune]bool)
		duplicate := false
		for _, c := range prev_4 {
			_, ok := seen[c]
			if ok {
				duplicate = true
			}
			seen[c] = true
		}

		if !duplicate {
			fmt.Println(i + 1)
			os.Exit(1)
		}
	}
}
