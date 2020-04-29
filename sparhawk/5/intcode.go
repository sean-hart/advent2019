package main

type program struct {
	code []int
	cur int // Execution cursor location
	input int
	output int
	mode int
}

func add(p program) program {
	param1, param2, loc := getParams(p)
	p.code[loc] = param1 + param2
	p.cur += 4
	return p
}

func multiply(p program) program {
	param1, param2, loc := getParams(p)
	p.code[loc] = param1 * param2
	p.cur += 4
	return p
}

func getParams(p program) (int, int, int) {

	param1 := p.code[p.cur+1]
	if (p.code[p.cur] / 100 % 10) == 0 {
		param1 = p.code[param1]
	}

	param2 := p.code[p.cur+2]
	if (p.code[p.cur] / 1000 % 10) == 0 {
		param2 = p.code[param2]
	}

	loc := p.code[p.cur+3]
	if (p.code[p.cur] / 10000 % 10) == 1 {
		loc = p.code[loc]
	}
	
	return param1, param2, loc
}

func write(p program) program {
	p.code[p.code[p.cur+1]] = p.input
	p.cur += 2
	return p
}

func read(p program) program {
	p.output = p.code[p.code[p.cur+1]]
	p.cur += 2
	return p
}

func intcode(p program) program {
	op := p.code[p.cur] % 100
	switch op {
	case 1:
		return add(p)
	case 2:
		return multiply(p)
	case 3:
		return write(p)
	case 4:
		return read(p)
	case 99:
		return p
	default:
		return p
	}
}

func runProgram(p program) program {
	for p.cur < len(p.code) {
		if p.code[p.cur] != 99 {
			p = intcode(p)
		} else {
			return p
		}
	}
	return p
}

func getDigit(num, pos int) int {
	s := intToSlice(num)
	if pos > len(s) {
		return 0
	}
	return s[pos]
}

func intToSlice(num int) []int {
	res := []int{}
	for num / 10 > 0 {
		res = append([]int{num % 10}, res...)
		num /= 10
	}
	res = append([]int{num % 10}, res...)
	return res
}