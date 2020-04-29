package libs

// import "fmt"

// Parameters hold the instructions for the computer.
type Parameters struct {
	leftReg, rightReg, storeReg int
}

// RunComputer will run opcode till termination
func RunComputer(inputInt int, memory []int, instructionPointer int) (outputInt int, outputMem []int, nextPointer int,) {
	// opcode := memory[instructionPointer]
	digits := GetDigits(memory[instructionPointer])
	opcode := digits[0] + (digits[1] * 10)

	// parameterModes := []int{digits[4], digits[3], digits[2]}
	parameterModes := []int{digits[2], digits[3], digits[4]}

	switch {
	case opcode == 99:
		return outputInt, memory, 99
	case opcode == 1:
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		memory[rawParams[2]]= parsedParams[0] + parsedParams[1]
		nextPointer = instructionPointer + 4
	case opcode == 2:
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		memory[rawParams[2]]= parsedParams[0] * parsedParams[1]
		nextPointer = instructionPointer + 4
	case opcode == 3:
		rawParams := getParams(1, memory, instructionPointer)
		memory[rawParams[0]] = inputInt
		nextPointer = instructionPointer + 2
	case opcode == 4:
		rawParams := getParams(1, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		outputInt = parsedParams[0]
		nextPointer = instructionPointer + 2
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

// ParseParams determines the actual value based on mode.
func ParseParams(params, paramModes, memory []int) (parsedParams []int) {
	for i, param := range params {
		if paramModes[i] == 1 {
			parsedParams = append(parsedParams, param)
		} else {
			parsedParams = append(parsedParams, memory[param])
		}
	}
	return parsedParams
}
