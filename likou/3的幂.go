package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for i := 3; i <= n; i *= 3 {
		if n == i {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(isPowerOfThree(9))
}
