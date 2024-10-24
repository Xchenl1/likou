package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
}
func main() {
	fmt.Println(checkPerfectNumber(28))
}
