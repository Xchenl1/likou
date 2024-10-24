package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//1.双重for循环显然不行 2.对撞双指针
	//s1 := make([]int, 0)
	//for i1 := 0; i1 < len(nums)-1; i1++ {
	//	k := nums[i1]
	//	for i := i1 + 1; i < len(nums); i++ {
	//		if k+nums[i] == target {
	//			s1 = append(s1, k)
	//			s1 = append(s1, nums[i])
	//			return s1
	//		}
	//	}
	//}
	//return nil
	left, right := 0, len(nums)-1
	for left < right {
		if left >= len(nums) || right < 0 {
			break
		}
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] == target {
			s1 := make([]int, 0)
			s1 = append(s1, nums[left])
			s1 = append(s1, nums[right])
			return s1
		} else {
			right++
		}
	}
	return nil
}

func main() {
	s := []int{1, 2, 3, 5}
	fmt.Println(twoSum(s, 6))
}
