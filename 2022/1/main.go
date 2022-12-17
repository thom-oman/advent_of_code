package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Part 1
// func main() {
// 	readFile, err := os.Open("inputs.txt")
// 	defer readFile.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)

// 	var highestSeen int64
// 	var cur int64

// 	for fileScanner.Scan() {
// 		x := fileScanner.Text()
// 		if x == "" {
// 			if cur > highestSeen {
// 				highestSeen = cur
// 			}
// 			cur = 0
// 			continue
// 		}
// 		n, err := strconv.ParseInt(x, 10, 32)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		cur += n
// 	}

// 	fmt.Println(highestSeen)
// }

// Part 2
func main() {
	readFile, err := os.Open("inputs.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var allSeen []int64
	var cur int64

	for fileScanner.Scan() {
		x := fileScanner.Text()
		if x == "" {
			allSeen = append(allSeen, cur)
			cur = 0
			continue
		} else {
			n, err := strconv.ParseInt(x, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			cur += n
		}
	}

	sort.Slice(allSeen, func(i, j int) bool {
		return allSeen[i] < allSeen[j]
	})

	var r int64

	for _, i := range allSeen[len(allSeen)-3:] {
		r += i
	}

	fmt.Println(r)
}
