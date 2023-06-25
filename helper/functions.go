package helper

import "strconv"

func ConvertToInt(str string) int {
	// Convert string to int, return 0 if conversion fails
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}
