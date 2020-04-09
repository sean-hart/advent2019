package main

type instructionSet struct {
	verb, leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(computerInput []int, startPosition int) (output []int, nextstart int) {
	if computerInput[startPosition] == 99 {
		return computerInput, 99
	}
	quartet := computerInput[startPosition:]
	inst := instructionSet{quartet[0], quartet[1], quartet[2], quartet[3]}

	switch inst.verb {
	case 99:
		return computerInput, 99
	case 1:
		computerInput[inst.storeReg] = computerInput[inst.leftReg] + computerInput[inst.rightReg]
	case 2:
		computerInput[inst.storeReg] = computerInput[inst.leftReg] * computerInput[inst.rightReg]
	}
	if computerInput[startPosition+4] != 99 {
		RunComputer(computerInput, startPosition+4)
	}
	return computerInput, 99
}
