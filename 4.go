package main

import (
	"fmt"
	"strconv"
)

func main() {
	greatest := 0

	for i := 999; i > 99; i-- {
		for j := 999; j >= i; j-- {
			n := i * j

			if n > greatest && isPalindromic(n) {
				greatest = n
			}
		}
	}

	fmt.Println(greatest)
}

func isPalindromic(n int) bool {
	num := strconv.Itoa(n)

	for i := 0; i < len(num)/2; i++ {
		if num[i:i+1] != num[len(num)-i-1:len(num)-i] {
			return false
		}
	}

	return true
}
