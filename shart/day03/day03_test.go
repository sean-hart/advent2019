package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManhattanDistances(t *testing.T) {
	for _, tc := range testCasesManhattanDistances {
		actual := GetManhattanDistance(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkManhattanDistances(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesManhattanDistances {
			GetManhattanDistance(tc.input)
		}
	}
}
