package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattanDistance(t *testing.T) {
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

func TestFindIntersections(t *testing.T) {
	for _, tc := range testCasesFindIntersections {
		actual := FindIntersections(tc.input)
		assert.Truef(t, DeepEqual(tc.expected, actual), "Expected: %v, Actual: %v", tc.expected, actual)
	}
}

func BenchmarkFindIntersections(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesFindIntersections {
			FindIntersections(tc.input)
		}
	}
}

func TestExpandRoute(t *testing.T) {
	for _, tc := range testCasesExpandRoute {
		actual := ExpandRoute(tc.input)
		assert.True(t, DeepEqual(tc.expected, actual))
	}
}

func BenchmarkExpandRoute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesExpandRoute {
			ExpandRoute(tc.input)
		}
	}
}

func TestShortestWireSum(t *testing.T) {
	for _, tc := range testCasesShortestWireSum {
		actual := ShortestWireSum(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkShortestWireSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesShortestWireSum {
			ShortestWireSum(tc.input)
		}
	}
}

// DeepEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func DeepEqual(a, b []Coord) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
