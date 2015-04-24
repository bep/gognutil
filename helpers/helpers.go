package helpers

import (
	"math"
)

// MaxInt returns the larger of x or y.
func MaxInt(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// NumDigits returns the numbers of digits in number.
// 0 is 1 digit, 1000 is 4.
func NumDigits(number int) int {
	var digits int
	var i = number
	for {
		i /= 10
		digits++
		if i == 0 {
			break
		}
	}
	return digits
}
