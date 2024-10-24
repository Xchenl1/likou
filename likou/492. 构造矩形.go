package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	s := make([]int, 2)
	cha := 999999999999999
	for i := area; i >= int(math.Sqrt(float64(area))); i-- {
		if area%i == 0 {
			chu := area / i
			if i-chu < cha && i > chu {
				s[0] = i
				s[1] = chu
			}
		}
	}
	return s
}
func main() {
	fmt.Println(constructRectangle(2))
}
