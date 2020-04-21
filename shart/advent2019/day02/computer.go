package main

type parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(memory []int, instructionPointer int) (output []int, nextPointer int) {
	if memory[instructionPointer] == 99 {
		return memory, 99
	}
	opcode := memory[instructionPointer]
	rawParams := memory[instructionPointer+1 : instructionPointer+4]
	inst := parameters{rawParams[0], rawParams[1], rawParams[2]}

	switch opcode {
	case 99:
		return memory, 99
	case 1:
		memory[inst.storeReg] = memory[inst.leftReg] + memory[inst.rightReg]
	case 2:
		memory[inst.storeReg] = memory[inst.leftReg] * memory[inst.rightReg]
	}
	if memory[instructionPointer+4] != 99 {
		RunComputer(memory, instructionPointer+4)
	}
	return memory, 99
}

func resetMem(memory []int) (newMem []int) {
	newMem = make([]int, len(memory))
	copy(newMem, memory)
	return
}

func findPair(mem, available []int, desired int) (left, right int) {
	for _, left := range available {
		for _, right := range available {
			newMem := resetMem(mem)
			newMem[1] = left
			newMem[2] = right
			newMem, _ = RunComputer(newMem, 0)
			if newMem[0] == desired {
				return left, right
			}
		}
	}
	return
}
