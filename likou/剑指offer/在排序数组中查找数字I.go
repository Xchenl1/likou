package main

import "fmt"

func search(nums []int, target int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			continue
		} else if nums[i] == target {
			j++
		} else if nums[i] > target {
			break
		}
	}
	return j
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(search(nums, target))
}
