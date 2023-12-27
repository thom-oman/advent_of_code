package main

import "testing"

func TestPossibleToMatch(t *testing.T) {
	tests := []struct {
		val  string
		sizes []int
		exp bool
	}{
		{".??.###", []int{1,3}, true},
		{".??.###", []int{2,3}, true},
		{".??.###", []int{3,2}, false},
		{"?.?.###", []int{1,3}, true},
		{"?.?.###", []int{1,1,3}, true},
		{"???.###", []int{1,1,3}, true},
		{"#??.###", []int{1,1,3}, true},
		{"##?.###", []int{1,1,3}, false},
		{"???????", []int{1,1,3}, true},
		{"????????", []int{1,1,3}, true},
		{"?..??.###", []int{1,3}, true},
		{"????.######..#####.", []int{1,6,5}, true},
		{"????..#####..#####.", []int{1,6,5}, false},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			if possibleToMatch(tt.val, tt.sizes) != tt.exp {
				t.Fatalf("Expected %v for %v and %v", tt.exp, tt.val, tt.sizes)
			}
		})
	}
}

func TestMatchSizes(t *testing.T) {
	tests := []struct {
		val  string
		sizes []int
		strIdx int 
		matched int 
	}{
		{".??.###", []int{1,3}, 1, 0},
		{".??.###", []int{3,2}, 1, 0},
		{"#.?.###", []int{1,3}, 2, 1},
		{"#...###", []int{1,3}, 6, 2},
		{"#??.###", []int{1,1,3}, 1, 1},
		{"##?.###", []int{1,1,3}, 0, -1},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			if strIdx, match := matchSizes(tt.val, tt.sizes); strIdx != tt.strIdx || match != tt.matched {
				t.Fatalf("Expected (%v,%v)for %v and %v", tt.strIdx, tt.matched, tt.val, tt.sizes)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		val  string
		sizes []int
		exp bool
	}{
		{".#..###", []int{1,3}, true},
		{".##.###", []int{2,3}, true},
		{".##.###", []int{1,3}, false},
		{"#...###", []int{1,3}, true},
		{"####.######..#####.", []int{1,6,5}, false},
		{"#...#####..#####.", []int{1,6,5}, false},
		{"#...#####..#####.", []int{1,6,5}, false},
		{"?...#####..#####.", []int{1,6,5}, false},
		{"????..#####..#####.", []int{1,6,5}, false},
	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			if isValid(tt.val, tt.sizes) != tt.exp {
				t.Fatalf("Expected %v for %v and %v", tt.exp, tt.val, tt.sizes)
			}
		})
	}
}
