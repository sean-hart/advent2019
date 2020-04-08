package main

import "fmt"

// RunComputer will run opcode till termination
func RunComputer(computerInput []int) (output []int) {
	// Take the first slice
	working := computerInput
	startPosition := 0
	code := working[startPosition:]
	for code[0] != 99 {
		fmt.Println(code[0])
		if code[0] == 1 {
			working[code[3]] = working[code[1]] + working[code[2]]
		} else if code[0] == 2 {
			working[code[3]] = working[code[1]] * working[code[2]]
		} else {
			return working
		}
		startPosition = startPosition + 4
		code = working[startPosition:]
	}
	output = working
	return output
}
