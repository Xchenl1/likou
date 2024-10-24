package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	//我的想法是利用冒泡算法
	sort.Ints(nums)
	if len(nums) < 3 {
		return nums[len(nums)-1]
	}
	k := make(map[int]int, 0)
	l := 0
	if len(nums) > 3 {
		for i := len(nums) - 1; i >= 0; i-- {
			if len(k) == 3 {
				return nums[i+1]
			}
			k[nums[i]]++
			l = nums[i]
		}
	} else if len(nums) == 3 {
		for i := 0; i < len(nums); i++ {
			k[nums[i]]++
		}
		if len(k) <= 2 {
			return nums[len(nums)-1]
		} else {
			return nums[0]
		}
	}
	if len(k) == 3 {
		return l
	}
	return nums[len(nums)-1]
}

func main() {
	i := []int{1, 2, 3}
	fmt.Println(thirdMax(i))
}
