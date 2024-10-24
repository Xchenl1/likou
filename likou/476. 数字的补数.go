package main

import (
	"fmt"
	"math"
)

func findComplement(num int) int {
	k := make([]int, 0)
	for num != 0 {
		k = append(k, num%2)
		num /= 2
	}
	for i, v := range k {
		if v == 0 {
			k[i] = 1
		} else if v == 1 {
			k[i] = 0
		}
	}
	for k1 := len(k) - 1; k1 >= 0; k1-- {
		if k[k1] == 0 {
			k = k[:k1]
		} else {
			break
		}
	}
	var sum float64 = 0
	for l, v := range k {
		sum += float64(v) * math.Pow(2, float64(l))
	}
	return int(sum)
}
func main() {
	num := 4
	fmt.Println(findComplement(num))
}
