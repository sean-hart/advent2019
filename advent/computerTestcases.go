package main

var testCasesResetComputer = []struct {
	description string
	input       []int
	expected    []int
}{
	{
		description: "1+1=2",
		input:       []int{1, 0, 0, 0, 99},
		expected:    []int{2, 0, 0, 0, 99},
	},
	{
		description: "3*2=6",
		input:       []int{2, 3, 0, 3, 99},
		expected:    []int{2, 3, 0, 6, 99},
	},
	{
		description: "99*99 = 9801",
		input:       []int{2, 4, 4, 5, 99, 0},
		expected:    []int{2, 4, 4, 5, 99, 9801},
	},
	{
		description: "more complex",
		input:       []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		expected:    []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
}
