package day04

import (
	"strconv"
	"strings"
)

// GetAllPasswords will get all valid passwords and return a slice.
// O = n
func GetAllPasswords(inputString string) (validPasswords []int) {
	inputMinMax := strings.Split(inputString, "-")
	min, _ := strconv.Atoi(inputMinMax[0])
	max, _ := strconv.Atoi(inputMinMax[1])
	// TODO: How can we filter here?
	for i := min; i <= max; i++ {
		if ValidatePassword(i) {
			validPasswords = append(validPasswords, i)
		}
	}
	return validPasswords
}

// ValidatePassword takes a 6 digit number and checks to see if it is a
// valid password.
// O = 2*6
func ValidatePassword(inputInteger int) bool {
	inputSlice := strings.Split(strconv.Itoa(inputInteger), "")

	if len(inputSlice) != 6 {
		return false
	}

	var groups = make(map[int]int)
	firstDigit, _ := strconv.Atoi(inputSlice[0])
	groups[firstDigit] = 1

	for i := 1; i < 6; i++ {
		current, _ := strconv.Atoi(inputSlice[i])
		previous, _ := strconv.Atoi(inputSlice[i-1])
		groups[current] = groups[current] + 1
		if current < previous {
			return false
		}
	}
	for _, v := range groups {
		if v == 2 {
			return true
		}
	}
	return false
}
