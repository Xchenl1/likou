package main

import "fmt"

func missingNumber(nums []int) int {
	for i, v := range nums {
		if v != i {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0}
	fmt.Println(missingNumber(nums))
}
