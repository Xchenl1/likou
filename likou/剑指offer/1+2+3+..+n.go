package main

import "fmt"

//	func sumNums(n int) int {
//		var k func(n int) bool
//		sum := 0
//		k = func(n int) bool {
//			sum += n
//			return n > 0 && k(n-1)
//		}
//		k(n)
//		return sum
//	}
func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return sumNums(n-1) + n
}

func main() {
	fmt.Println(sumNums(6))
}
