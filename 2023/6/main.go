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
		time int
		dist int
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
			var bs []byte
			for i := 1; i < len(p); i++ {
				bs = append(bs,[]byte(p[i])...)
			}
			time, _ = strconv.Atoi(string(bs))
		} else {
			var bs []byte
			for i := 1; i < len(p); i++ {
				bs = append(bs,[]byte(p[i])...)
			}
			dist, _ = strconv.Atoi(string(bs))
		}
	}

	var c int
	for j := 1; j < time; j++ {
		r := time - j

		if (r * j) > dist {
			c++
		}
	}

	fmt.Println("result", c)
}

