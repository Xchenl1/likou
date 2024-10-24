package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	//用map
	var i = 0
	a := make(map[int]int)
	for i = 0; i < len(nums); i++ {
		a[nums[i]]++
		count, _ := a[nums[i]]
		if count > len(nums)/2 {
			break
		}
	}
	return nums[i]
}

func main() {
	a := []int{3, 3, 4}
	fmt.Println(majorityElement(a))
}
