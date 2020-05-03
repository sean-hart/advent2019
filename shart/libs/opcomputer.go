package libs

// RunComputer will run opcode till halt.
func RunComputer(inputInt int, memory []int, instructionPointer int, currentOutput int) (outputInt int, outputMem []int, nextPointer int) {
	opcode := memory[instructionPointer] % 100
	parameterModes := GetDigits(memory[instructionPointer] / 100)

	switch {
	case opcode == 99:
		return outputInt, memory, 99

	case opcode == 1:
		// p3 = p1 + p2
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		memory[rawParams[2]] = parsedParams[0] + parsedParams[1]
		nextPointer = instructionPointer + 4

	case opcode == 2:
		// p3 = p1 * p2
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		memory[rawParams[2]] = parsedParams[0] * parsedParams[1]
		nextPointer = instructionPointer + 4

	case opcode == 3:
		//p1 = input
		rawParams := getParams(1, memory, instructionPointer)
		memory[rawParams[0]] = inputInt
		nextPointer = instructionPointer + 2

	case opcode == 4:
		//output = p1
		rawParams := getParams(1, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		outputInt = parsedParams[0]
		nextPointer = instructionPointer + 2

	case opcode == 5:
		//jumpto if p1 != 0
		rawParams := getParams(2, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		if parsedParams[0] != 0 {
			nextPointer = parsedParams[1]
		} else {
			nextPointer = instructionPointer + 3
		}

	case opcode == 6:
		//jumpto if p1 == 0
		rawParams := getParams(2, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		if parsedParams[0] == 0 {
			nextPointer = parsedParams[1]
		} else {
			nextPointer = instructionPointer + 3
		}

	case opcode == 7:
		//if p1 < p2 set p3 to 1 else 0
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		if parsedParams[0] < parsedParams[1] {
			memory[rawParams[2]] = 1
		} else {
			memory[rawParams[2]] = 0
		}
		nextPointer = instructionPointer + 4
	case opcode == 8:
		// if p1 == p2 set p3 to 1 else 0
		rawParams := getParams(3, memory, instructionPointer)
		parsedParams := ParseParams(rawParams, parameterModes, memory)
		// fmt.Printf("Parsed0: %v, Parsed1: %v", parsedParams[0], parsedParams[1])
		if parsedParams[0] == parsedParams[1] {
			memory[rawParams[2]] = 1
		} else {
			memory[rawParams[2]] = 0
		}
		nextPointer = instructionPointer + 4
	}

	if outputInt == 0 {
		outputInt = currentOutput
	}

	if memory[nextPointer] != 99 {
		outputInt, memory, nextPointer = RunComputer(inputInt, memory, nextPointer, outputInt)
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
func GetDigits(number int) (digits []int) {
	for number >= 1 {
		digits = append(digits, number%10)
		number = number / 10
	}

	for len(digits) < 3 {
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
