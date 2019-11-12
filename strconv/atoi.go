package strconv

import (
	"errors"
	"math"
)

// AtoI string to int conversion using base 10 expected
func AtoI(s string) (int, error) {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]

		if i == 0 && char == '-' {
			result *= -1
		} else if char < '0' || char > '9' {
			return 0, errors.New("Not a valid number")
		} else {
			num := int(char) - '0'
			mod := int(math.Pow(10.0, float64(len(s)-i-1)))
			result += num * mod
		}
	}

	return result, nil
}
