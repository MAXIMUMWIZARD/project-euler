package main

import "fmt"

/* Even Fibonacci numbers

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms. */

func main() {
	a, b, even, sum := 1, 1, 2, 0

	for even < 4000000 {
		sum += even

		a = b + even
		b = even + a
		even = a + b
	}

	fmt.Println(sum)
}
