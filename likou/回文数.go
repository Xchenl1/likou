package main

import (
	"fmt"
	"strconv"
)

func isPalindrome1(x int) bool {
	x1 := strconv.Itoa(x)
	for i := 0; i < len(x1)/2; i++ {
		if x1[i] != x1[len(x1)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome1(10))
}
