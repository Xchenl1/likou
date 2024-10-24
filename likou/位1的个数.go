package main

import "fmt"

func hammingWeight(num uint32) int {
	var sum int
	for i := 0; i < 32; i++ {
		k := num & 1
		if k == 1 {
			sum += 1
		}
		num >>= 1
	}
	fmt.Println(sum)
	return sum
}

func main() {
	var a uint32
	a = 0b10111100101111111111111111111111
	hammingWeight(a)
}
