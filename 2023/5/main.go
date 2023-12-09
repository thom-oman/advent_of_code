package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	mapmatch = regexp.MustCompile("[map:]")
)

func main() {
	lines := parseInputFile()
	var seeds []int
	for _, s := range strings.Fields(strings.Split(lines[0], ":")[1]) {
		n, _ := strconv.Atoi(s)
		seeds = append(seeds, n)
	}

	var (
		m        *mapping
		mappings []*mapping
	)

	for _, l := range lines[1:] {
		if len(l) == 0 {
			continue
		}

		if mapmatch.Match([]byte(l)) {
			mappings = append(mappings, NewMapping(l))
			continue
		}
		m = mappings[len(mappings)-1]

		var rval []int
		for _, f := range strings.Fields(l) {
			r, _ := strconv.Atoi(f)
			rval = append(rval, r)
		}
		m.AddRange(rval[0], rval[1], rval[2])
	}

	blah := math.Inf(1)
	for si := range seeds {
		s := seeds[si]
		for i := range mappings {
			s = mappings[i].Map(s)
		}
		if float64(s) < blah {
			blah = float64(s)
		}
	}

	fmt.Println(int(blah))
}

type mapping struct {
	from, to    string
	ranges []nrange
}

type nrange struct {
	dest, source, length int
}

func NewMapping(l string) *mapping {
	maps := strings.Split(strings.Split(l, "map:")[0], "-")
	return &mapping{
		from:        maps[0],
		to:          maps[2],
	}
}

func (m *mapping) Map(val int) int {
	for i := range m.ranges {
		r := m.ranges[i]
		// fmt.Println("Using mapping:", r.source, r.dest, r.length)
		if r.source <= val && val < r.source + r.length {
			return val + (r.dest - r.source)
		}
	}
	return val
}

func (m *mapping) AddRange(dest, source, length int) {
	m.ranges = append(m.ranges, nrange{dest, source, length})
}

func parseInputFile() []string {
	f, err := os.Open("inputs.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	var lines []string
	for {
		l, err := r.ReadBytes('\n')
		if len(l) == 0 {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		lines = append(lines, string(l[:len(l)-1]))
	}

	return lines
}
