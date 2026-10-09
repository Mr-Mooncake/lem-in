package validation

import (
	"fmt"
	"strconv"
)

func ants(input []string) bool {
	ants, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		return false
	}

	if ants <= 0 {
		fmt.Println("ERROR: invalid data format, number of ants should be greater than or equal to 1")
		return false
	}
	return true
}
