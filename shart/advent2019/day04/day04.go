package day04

import (
	"strconv"
	"strings"
)

// GetAllPasswords will get all valid passwords and return a slice.
func GetAllPasswords(inputString string) (validPasswords []int) {
	inputMinMax := strings.Split(inputString, "-")
	min, _ := strconv.Atoi(inputMinMax[0])
	max, _ := strconv.Atoi(inputMinMax[1])
	for i := min; i <= max; i++ {
		if ValidatePassword(i) {
			validPasswords = append(validPasswords, i)
		}
	}
	return validPasswords
}

// ValidatePassword takes a 6 digit number and checks to see if it is a
// valid password.
func ValidatePassword(inputInteger int) bool {
	inputSlice := strings.Split(strconv.Itoa(inputInteger), "")
	doubles := false
	if len(inputSlice) != 6 {
		return false
	}
	for i := 1; i < 6; i++ {
		current, _ := strconv.Atoi(inputSlice[i])
		previous, _ := strconv.Atoi(inputSlice[i-1])
		if current == previous {
			doubles = true
		}
		if current < previous {
			return false
		}
	}
	if doubles {
		return true
	}
	return false
}
