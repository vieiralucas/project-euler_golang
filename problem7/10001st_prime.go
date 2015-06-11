// What is the 10 001st prime number?

package main

import (
	"fmt"
	"math"
)

func main() {
	current := 3
	primes := []int{2}

	for len(primes) < 10001 {
		if tryPreviousPrimes(primes, current) {
			primes = append(primes, current)
		}

		current += 2
	}

	fmt.Println(primes[10000])
}

func tryPreviousPrimes(primes []int, current int) bool {
	sqrt := int(math.Sqrt(float64(current)))
	if sqrt == 1 {
		return true
	}

	for i := 0; i < len(primes); i++ {
		// fmt.Printf("sqrt: %d - trying: %d - current: %d - len: %d\n", sqrt, primes[i], current, len(primes))
		if primes[i] > sqrt {
			return true
		}

		if current%primes[i] == 0 {
			return false
		}
	}

	return true
}
