package main

import "fmt"

func main() {
	fmt.Println(sumUnder(3, 1000) + sumUnder(5, 1000) - sumUnder(15, 1000))
}

func sumUnder(base int, limit int) int {
	count := (limit - 1) / base

	if count%2 == 0 {
		return (count / 2) * (count/2*base + (count/2+1)*base)
	} else {
		return count * (count/2 + 1) * base
	}
}
