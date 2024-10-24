package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	k := 1
	for k <= n {
		k = k * 2
		if k == n {
			return true
		}
	}
	return false
}

func main() {
	k := 16
	fmt.Print(isPowerOfTwo(k))
}
