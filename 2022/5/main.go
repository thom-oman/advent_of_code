package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	Crates []byte
}

func NewStack(crates []byte) *stack {
	return &stack{Crates: crates}
}

func (s *stack) Add(crate byte) {
	s.Crates = append(s.Crates, crate)
}

func (s *stack) AddMultiple(crates []byte) {
	s.Crates = append(s.Crates, crates...)
}

func (s stack) Peak() byte {
	return s.Crates[len(s.Crates)-1]
}

func (s stack) Size() int {
	return len(s.Crates)
}

func (s *stack) Pop() byte {
	var res byte
	if len(s.Crates) == 1 {
		res = s.Crates[0]
		s.Crates = make([]byte, 0)
		return res
	}
	if len(s.Crates) == 0 {
		log.Fatal("EMPTY CrATES")
	}
	res = s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[:len(s.Crates)-1]
	return res
}

func (s *stack) PopN(n int, to *stack) {
	// Part 1
	// for c := 0; c < n; c++ {
	// 	to.Add(s.Pop())
	// }

	// Part 2
	if len(s.Crates) < n {
		log.Fatal("len(s.Crates) < n")
	}

	tmp := make([]byte, n)
	for i := 0; i < n; i++ {
		// tmp = append(tmp, s.Pop())
		tmp[n-1-i] = s.Pop()
	}
	to.AddMultiple(tmp)
}

type instruction struct {
	N, From, To int64
}

func NewInstruction(n, from, to int64) *instruction {
	return &instruction{
		N:    n,
		From: from,
		To:   to,
	}
}

var stacks []*stack
var instructions []*instruction

func main() {
	parseInput()
	processInstructions()
	for _, stack := range stacks {
		fmt.Print(string(stack.Peak()))
	}
}

func processInstructions() {
	for _, ins := range instructions {
		from := stacks[ins.From-1]
		to := stacks[ins.To-1]
		from.PopN(int(ins.N), to)
	}
}

func parseInput() {
	readFile, err := os.Open("inputs.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	stacksDefinition := make([]string, 0)
	instructionsDefinition := make([]string, 0)
	finishedStacked := false

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			finishedStacked = true
			continue
		}
		if finishedStacked {
			instructionsDefinition = append(instructionsDefinition, line)
		} else {
			stacksDefinition = append(stacksDefinition, line)
		}
	}

	buildStacks(stacksDefinition)
	buildInstructions(instructionsDefinition)
}

func buildStacks(definition []string) {
	n_stacks := (len(definition[0]) + 1) / 4
	stacks_map := make([][]byte, n_stacks)
	stacks = make([]*stack, n_stacks)

	for x := len(definition) - 2; x >= 0; x-- {
		l := definition[x]
		stack_i := 0
		for i := 0; i+3 <= len(l); i += 4 {
			crate := l[i+1]
			if crate != 32 {
				stacks_map[stack_i] = append(stacks_map[stack_i], crate)
			}
			stack_i += 1
		}
	}

	for idx := range stacks_map {
		stacks[idx] = NewStack(stacks_map[idx])
	}
}

func buildInstructions(definition []string) {
	instructions = make([]*instruction, len(definition))
	for idx := range definition {
		s := strings.Split(definition[idx], " ")
		n, _ := strconv.ParseInt(s[1], 10, 0)
		from, _ := strconv.ParseInt(s[3], 10, 0)
		to, _ := strconv.ParseInt(s[5], 10, 0)

		instructions[idx] = NewInstruction(n, from, to)
	}
}
