package libs

// import "fmt"

// Parameters hold the instructions for the computer.
type Parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(inputInt int, memory []int, instructionPointer int) (outputInt int, outputMem []int, nextPointer int,) {
	// if memory[instructionPointer] == 99 {
	// 	return outputInt, memory, 99
	// }
	opcode := memory[instructionPointer]
	// rawParams := memory[instructionPointer+1 : instructionPointer+4]
	// inst := Parameters{rawParams[0], rawParams[1], rawParams[2]}
	// nextPointer := instructionPointer

	switch {
	case opcode == 99:
		return outputInt, memory, 99
	case opcode == 1:
		rawParams := getParams(3, memory, instructionPointer)
		inst := Parameters{rawParams[0], rawParams[1], rawParams[2]}
		memory[inst.storeReg] = memory[inst.leftReg] + memory[inst.rightReg]
		nextPointer = instructionPointer + 4
	case opcode == 2:
		rawParams := getParams(3, memory, instructionPointer)
		inst := Parameters{rawParams[0], rawParams[1], rawParams[2]}
		memory[inst.storeReg] = memory[inst.leftReg] * memory[inst.rightReg]
		nextPointer = instructionPointer + 4
	case opcode == 3:
		rawParams := getParams(1, memory, instructionPointer)
		memory[rawParams[0]] = inputInt
		nextPointer = instructionPointer + 2
	case opcode == 4:
		rawParams := getParams(1, memory, instructionPointer)
		outputInt = memory[rawParams[0]]
		nextPointer = instructionPointer + 2
	case opcode >= 100:
		switch {
			
		}
		return 0, memory, 99
	}
	if memory[nextPointer] != 99 {
		outputInt, memory, nextPointer = RunComputer(inputInt, memory, nextPointer)
	}
	return outputInt, memory, 99
}

//ResetMem will reset memory to a known state.
func ResetMem(memory []int) (newMem []int) {
	newMem = make([]int, len(memory))
	copy(newMem, memory)
	return
}

func getParams(numParams int, memory []int, pointer int) (params []int) {
	rawParams := memory[pointer+1 : pointer+numParams+1]
	return rawParams
}

// GetDigits will return a slice of digits as [ones, tens, hundreds, ...].
func GetDigits(number int) (digits []int){
	for number >= 1 {
		digits = append(digits, number % 10)
		number = number / 10
	}
	
	for len(digits) < 5 {
		digits = append(digits, 0)
	}

	return digits
}
