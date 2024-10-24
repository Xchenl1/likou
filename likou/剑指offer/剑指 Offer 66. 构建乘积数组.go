package main

import (
	"fmt"
)

// func constructArr(a []int) []int {
// i := 0
// var sum int = 1
// k := 0
// b := make([]int, 0)
//
//	for i < len(a) {
//		if i != k {
//			sum *= int(a[i])
//		}
//		i++
//		if i == len(a) {
//			b = append(b, sum)
//			sum = 1
//			i = 0
//			k++
//		}
//		if k == len(a) {
//			break
//		}
//	}
//
// return b
func productExceptSelf(a []int) []int {
	//答案上的写法是先将a[i]左边的数乘积放在一个数组中，之后再将右边数与其相乘放在相应位置
	k := make([]int, len(a))
	k1 := make([]int, len(a))
	for i1 := 0; i1 < len(a); i1++ {
		if i1 == 0 {
			k[i1] = 1
			k1[len(a)-i1-1] = 1
		} else {
			k[i1] = k[i1-1] * a[i1-1]
			k1[len(a)-i1-1] = a[len(a)-i1] * k1[len(a)-i1]
		}
	}
	for i, _ := range k {
		k[i] *= k1[i]
	}
	return k
}

func main() {
	a := []int{7, 2, 2, 4, 2, 1, 8, 8, 9, 6, 8, 9, 6, 3, 2, 1}
	fmt.Println(productExceptSelf(a))
}
