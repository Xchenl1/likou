package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	for i := 1; i*i < num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(isPerfectSquare(2147483647))
}
