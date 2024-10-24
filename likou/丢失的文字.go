package main

import "fmt"

func missingNumber(nums []int) int {
	j := 0
	for i := 0; i <= len(nums); i++ {
		//如果在数组中找到与其相同的值就跳出循环
		for j = 0; j < len(nums); j++ {
			if i == nums[j] {
				break
			}
		}
		if j == len(nums) {
			return i
		}
	}
	return 0
}
func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//missingNumber(nums)
	fmt.Println(missingNumber(nums))
}
