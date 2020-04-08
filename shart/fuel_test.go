package main

import (
	"math"
	"testing"
)

func TestGetModuleFuelRequirement(t *testing.T) {
	for _, tc := range testCasesGetModuleFuelReqirement {
		actual := GetModuleFuelRequirement(tc.moduleMass)
		if math.IsNaN(float64(actual)) || math.Abs(float64(actual-tc.expected)) > 0 {
			t.Fatalf("FAIL: %v\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkGetModuleFuelRequirement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesGetModuleFuelReqirement {
			GetModuleFuelRequirement(tc.moduleMass)
		}
	}
}

func TestGetFuelTotal(t *testing.T) {
	for _, tc := range testCasesGetFuelTotal {
		actual := GetFuelTotal(tc.modulemasses)
		if math.IsNaN(float64(actual)) || math.Abs(float64(actual-tc.expected)) > 0 {
			t.Fatalf("FAIL: %v\nExpected: %#v\nActual%#v", tc.description, tc.expected, actual)
		}
	}
}

func BenchmarkGetFuelTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesGetFuelTotal {
			GetFuelTotal(tc.modulemasses)
		}
	}
}
