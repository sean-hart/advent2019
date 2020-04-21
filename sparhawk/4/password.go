package password

// import (
// 	"fmt"
// )

func validPassword(password int) bool {
	numSlice := intToSlice(password)
	if !checkLength(numSlice) {
		return false
	}
	if !checkAdjacentPair(numSlice) {
		return false
	}
	if !checkNotDecrementing(numSlice) {
		return false
	}
	
	return true
}

func checkLength(password []int) bool {
	return len(password) == 6
}

func checkNotDecrementing(password []int) bool {
	min := password[0]
	for _, digit := range password {
		if digit < min {
			return false
		}
		min = digit
	} 
	return true
}

func checkAdjacentPair(password []int) bool {
	var lastDigit int
	for _, digit := range password {
		if digit == lastDigit {
			return true
		}
		lastDigit = digit
	}
	return false
}

func getKeys(min int, max int) []int {
	keys := []int{}
	
	for i := min; i <= max; i++ {
		if validPassword(i) {
			keys = append(keys, i)
		}
	}
	return keys
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