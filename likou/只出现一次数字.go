package main

import "fmt"

func singleNumber(nums []int) int {
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a = a ^ nums[i]
	}
	return a
}

func main() {
	nums := []int{1, 2, 2, 1, 3}
	fmt.Println(singleNumber(nums))
}
