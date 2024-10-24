package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
}

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	for _, j := range nums {
		fmt.Print(j, " ")
	}
}
