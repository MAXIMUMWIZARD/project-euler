package main

import (
	"fmt"
	"math"
)

func main() {
	// for n > 2 p(n) < 2^n
	// for n > 26 p(n) , 1.2^n
	// for n > 7022 p(n) < nlogn + n(loglogn - .9385)
	n := 10001
	bound := int(math.Floor(float64(n) * math.Log2(float64(n)) + float64(n) * (math.Log2(math.Log2(float64(n))) - 0.9385)))

	primes := make([]bool, bound + 1)
	for i := 0; i <= bound; i++ {
		primes[i] = true
	}

	for i := 2; i <= bound; i++ {
		if !primes[i] {
			continue
		}

		for j := i + i; j <= bound; j += i {
			primes[j] = false
		}
	}

	count := 0
	for i := 2; i <= bound; i++ {
		if primes[i] {
			count++
		}

		if count == n {
			fmt.Println(i)
			return
		}
	}
}
