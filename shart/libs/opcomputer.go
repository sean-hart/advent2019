package libs

// Parameters hold the instructions for the computer.
type Parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(inputInt int, memory []int, instructionPointer int) (outputInt int, outputMem []int, nextPointer int) {
	if memory[instructionPointer] == 99 {
		return outputInt, memory, 99
	}
	opcode := memory[instructionPointer]
	rawParams := memory[instructionPointer+1 : instructionPointer+4]
	inst := Parameters{rawParams[0], rawParams[1], rawParams[2]}

	switch opcode {
	case 99:
		return outputInt, memory, 99
	case 1:
		memory[inst.storeReg] = memory[inst.leftReg] + memory[inst.rightReg]
	case 2:
		memory[inst.storeReg] = memory[inst.leftReg] * memory[inst.rightReg]
	}
	if memory[instructionPointer+4] != 99 {
		RunComputer(inputInt, memory, instructionPointer+4)
	}
	return outputInt, memory, 99
}

//ResetMem will reset memory to a known state.
func ResetMem(memory []int) (newMem []int) {
	newMem = make([]int, len(memory))
	copy(newMem, memory)
	return
}
