package main

import (
	"github.com/sean-hart/advent2019/shart/libs"
)

type parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(memory []int, instructionPointer int) (output []int, nextPointer int) {
	output, nextPointer = libs.RunComputer(memory, instructionPointer)
	return output, nextPointer
}
