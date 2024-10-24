package main

import "fmt"

func hammingDistance(x int, y int) int {
	k1 := make([]int, 0)
	for x != 0 {
		k1 = append(k1, x%2)
		x /= 2
	}
	k2 := make([]int, 0)
	for y != 0 {
		k2 = append(k2, y%2)
		y /= 2
	}
	num := 0
	if len(k1) > len(k2) {
		for p := len(k2); p < len(k1); p++ {
			k2 = append(k2, 0)
		}
		k22 := Reverse(k2)
		k11 := Reverse(k1)
		for i := 0; i < len(k11); i++ {
			if k11[i] != k22[i] {
				num++
			}
		}
	} else if len(k1) < len(k2) {
		for p := len(k1); p < len(k2); p++ {
			k1 = append(k1, 0)
		}
		k11 := Reverse(k1)
		k22 := Reverse(k2)
		for i := 0; i < len(k11); i++ {
			if k11[i] != k22[i] {
				num++
			}
		}
	} else if len(k1) == len(k2) {
		for i := 0; i < len(k1); i++ {
			if k1[i] != k2[i] {
				num++
			}
		}
	}
	return num
}

func Reverse(num []int) []int {
	for i := 0; i < len(num)/2; i++ {
		a := num[i]
		num[i] = num[len(num)-1-i]
		num[len(num)-1-i] = a
	}
	return num
}

func main() {
	fmt.Println(hammingDistance(93, 73))
}
