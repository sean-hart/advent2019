package main

import (
	"testing"
	"github.com/sean-hart/advent2019/shart/libs"
	"github.com/stretchr/testify/assert"
)

func TestRunComputer(t *testing.T) {
	for _, tc := range testCasesRunComputer {
		newInput := make([]int, len(tc.input))
		copy(newInput, tc.input)
		_, actual, _ := RunComputer(0, newInput, 0)
		assert.Equal(t, tc.expected, actual)
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRunComputer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesRunComputer {
			newInput := make([]int, len(tc.input))
			copy(newInput, tc.input)
			RunComputer(0, newInput, 0)
		}
	}
}

func TestFindPair(t *testing.T) {
	goal := int(19690720)
	left, right := findPair(day2Input, makeRange(1, 100), goal)
	assert.Equal(t, 5121, (100*left)+right)
}

func TestFindPair2(t *testing.T) {
	goal := int(4714701)
	left, right := findPair(day2Input2, makeRange(1, 99), goal)
	assert.Equal(t, 1202, (100*left)+right)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func findPair(mem, available []int, desired int) (left, right int) {
	for _, left := range available {
		for _, right := range available {
			newMem := libs.ResetMem(mem)
			newMem[1] = left
			newMem[2] = right
			_, newMem, _ = RunComputer(0, newMem, 0)
			if newMem[0] == desired {
				return left, right
			}
		}
	}
	return
}
