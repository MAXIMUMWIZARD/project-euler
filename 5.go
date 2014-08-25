package main

import "fmt"

func factorable(n int) bool {
	for i := 1; i <= 20; i++ {
		if n % i != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 20; ; i += 20 {
		if factorable(i) {
			fmt.Println(i)
			return
		}
	}
}
