package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}

	var (
		times []int
		dists []int
	)

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadBytes('\n')
		if len(l) <= 1 {
			break
		}

		if err != nil && err != io.EOF {
			panic(err)
		}

		fmt.Println(l)
		p := strings.Fields(string(l))
		if p[0] == "Time:" {
			for i := 1; i < len(p); i++ {
				n, _ := strconv.Atoi(p[i])
				times = append(times,n) 
			}
		} else {
			for i := 1; i < len(p); i++ {
				n, _ := strconv.Atoi(p[i])
				dists = append(dists,n) 
			}
		}
	}

	res := 1
	for i := range times {
		t,d := times[i], dists[i]

		var c int
		for j := 1; j < t; j++ {
			r := t - j

			if (r * j) > d {
				c++
			}
		}
		res *= c
	}

	fmt.Println("result", res)
}

