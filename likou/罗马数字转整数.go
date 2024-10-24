package main

import (
	"fmt"
)

//var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
//
//func romanToInt(s string) (ans int) {
//	//核心就是用左边比右边的数据小 那么就减去其他情况就直接相加就ok了
//	n := len(s)
//	for i := range s {
//		value := symbolValues[s[i]]
//		if i < n-1 && value < symbolValues[s[i+1]] {
//			ans -= value
//		} else {
//			ans += value
//		}
//	}
//	return
//}

var A = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	num := 0
	for i := 0; i <= len(s)-1; i++ {
		k := A[string(s[i])]
		if i < len(s)-1 && k < A[string(s[i+1])] {
			num -= k
		} else {
			num += k
		}
	}
	return num
}
func main() {
	fmt.Println(romanToInt("IV"))
}
