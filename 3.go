package main

import "fmt"
import "math"

func main() {
	num := 600851475143
	limit := int(math.Floor(math.Sqrt(float64(num))))
	primes := make([]bool, limit)

	for i := 0; i < limit; i++ {
		primes[i] = (num%(i+2) == 0)
	}

	for i := 0; i < limit; i++ {
		if !primes[i] {
			continue
		}

		step := i + 2
		for j := i + step; j < limit; j += step {
			primes[j] = false
		}
	}

	for i := limit - 1; i >= 0; i-- {
		if primes[i] {
			fmt.Println(i + 2)
			return
		}
	}
}
