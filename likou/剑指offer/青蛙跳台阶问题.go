package main

import "fmt"

func numWays(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if n == 0 {
		return 1
	}
	i := 3
	nums := make([]int, 0)
	nums = append(nums, 0, 1, 2)
	for i <= n {
		num := nums[i-2] + nums[i-1]
		nums = append(nums, num%1000000007)
		i++
	}
	return nums[len(nums)-1]
}

func main() {
	fmt.Println(numWays(7))
}
