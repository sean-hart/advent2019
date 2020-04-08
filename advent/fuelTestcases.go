package main

var testCasesGetModuleFuelReqirement = []struct {
	description string
	moduleMass  int
	expected    int
}{
	{
		description: "Fuel required for a 12 mass module",
		moduleMass:  12,
		expected:    2,
	},
	{
		description: "Fuel required for a 14 mass module",
		moduleMass:  14,
		expected:    2,
	},
	{
		description: "Fuel required for a 1969 mass module",
		moduleMass:  1969,
		expected:    966,
	},
	{
		description: "Fuel required for a 100756 mass module",
		moduleMass:  100756,
		expected:    50346,
	},
	{
		description: "Fuel required for a 2 mass module",
		moduleMass:  2,
		expected:    0,
	},
}

var testCasesGetFuelTotal = []struct {
	description  string
	modulemasses []int
	expected     int
}{
	{
		description:  "Simple 1 item test",
		modulemasses: []int{12},
		expected:     2,
	},
	{
		description:  "Two item test",
		modulemasses: []int{12, 12},
		expected:     4,
	},
	{
		description: "more complicated test",
		modulemasses: []int{12, 12, 1969},
		expected: 970,
	},
}
