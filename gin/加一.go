package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	//找9的个数
	var b = len(digits)
	i := 0
	for b > 0 {
		if digits[b-1] == 9 {
			i++
		}
		b--
	}
	if i == len(digits) {
		digits = append(digits, 1)
		for i1 := 0; i1 < len(digits); i1++ {
			if i1 == 0 {
				digits[i1] = 1
			} else {
				digits[i1] = 0
			}

		}
	} else {
		var b2 = len(digits)
		for b2 > 0 {
			if digits[b2-1] == 9 {
				digits[b2-1] = 0
			} else if digits[b2-1] != 9 {
				digits[b2-1]++
				break
			}
			b2--
		}
	}
	fmt.Println(digits)
	return digits
}
func main() {
	var a []int = []int{9, 8, 9}
	plusOne(a)
}
