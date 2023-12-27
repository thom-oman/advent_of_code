package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	grid [][]rune
	galaxies []coordinates
	maxDist int
)

var (
	countx = make(map[int]int)
	county = make(map[int]int)
)

func main() {
	parseInputFile()
	determineDistances()
}

func determineDistances() {
	fmt.Println("No galaxies:", len(galaxies))
	fmt.Println("Countx:", countx)
	fmt.Println("County:", county)
	var total int
	for i := range galaxies {
		for j := range galaxies {
			if i < j {
				total += determineDistanceBetween(i, j)
			}
		}
	}
	fmt.Println("Total:", total)
}

func determineDistanceBetween(i,j int) int {
	var tot float64

	a, b := galaxies[i], galaxies[j]

	tot += math.Abs(a.x - b.x)
	tot += math.Abs(a.y - b.y)

	for ii := math.Min(a.x, b.x); ii < math.Max(a.x, b.x); ii++ {
		if countx[int(ii)] == 0 {
			tot += 999999
		}
	}

	for ii := math.Min(a.y, b.y); ii < math.Max(a.y, b.y); ii++ {
		if county[int(ii)] == 0 {
			tot += 999999 
		}
	}

	return int(tot)
}

type coordinates struct {
	x,y float64
}

func newCoordinates(x,y int) coordinates {
	return coordinates{x: float64(x), y: float64(y)}
}

func parseInputFile() {
	s, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(s), "\n") 

	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		rw := make([]rune, 0)

		for j := range lines[i] {
			r := rune(lines[i][j])
			rw = append(rw, r)

			if r == '#' {
				countx[i]++
				county[j]++
				galaxies = append(galaxies, newCoordinates(i, j))
			}
		}
		grid = append(grid, rw)
	}
}
