package main

import (
	"sort"
)

func maximumProduct(nums []int) int {
	// 排序
	sort.Ints(nums)
	// 总数为3 或者 有一个负数
	if len(nums) == 3 || (nums[0] < 0 && nums[1] >= 0) {
		return nums[0] * nums[1] * nums[2]
	}
	// 如果负数个数大于等于2 需要比较两个负数相乘x最大正数 与  三个最大整数积 比较
	if nums[0] < 0 && nums[1] < 0 {
		if nums[0]*nums[1]*nums[len(nums)-1] > nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3] {
			return nums[0] * nums[1] * nums[len(nums)-1]
		} else {
			return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
		}
	}
	return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
}
