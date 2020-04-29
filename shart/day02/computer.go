package main

import (
	"github.com/sean-hart/advent2019/shart/libs"
)

type parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(input int, memory []int, instructionPointer int) ( output int, outMem []int, nextPointer int) {
	output, outMem, nextPointer = libs.RunComputer(input, memory, instructionPointer, 0)
	return 0, outMem, nextPointer
}
