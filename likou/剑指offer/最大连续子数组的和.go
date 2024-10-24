package main

import "fmt"

// 动态规划解决方案
// 寻找每个元素结尾的最大子数组之和  可以先找一个数与他前面的数的和看看是否大于该数
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
func main() {
	nums := []int{1, 2}
	fmt.Println(maxSubArray(nums))
}
