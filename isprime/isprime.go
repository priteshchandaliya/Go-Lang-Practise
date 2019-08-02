package isprime

import (
	"math"
)

// IsPrime returns if a number is prime or not.
func IsPrime(value int) bool {

	if value <= 1 {
		return false
	}

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}

	return true
}
