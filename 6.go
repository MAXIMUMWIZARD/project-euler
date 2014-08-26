package main

import (
	"fmt"
)

func main() {
	n := 100
	sum_of_squares := 0
	for i := 1; i <= n; i++ {
		sum_of_squares += i * i
	}

	square_of_sums := (n + 1) * (n + 1) * n * n / 4

	fmt.Println(sum_of_squares, square_of_sums, square_of_sums - sum_of_squares)
}
