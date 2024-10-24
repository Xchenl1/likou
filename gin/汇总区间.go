package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	s1 := make([]string, 0)
	if len(nums) == 0 {
		return s1
	}
	nums = append(nums, 1000000)
	var l int
	var start int
	l = nums[0]
	//不妨换个思路 确定开始值和最终值 因为如果每两个相邻的差为1 就可以继续向后走
	for i, z := range nums {
		//将前一个数组的值传给l
		if i >= 1 {
			if z-l != 1 {
				if i-start > 1 {
					tmp := strconv.Itoa(nums[start]) + "->" + strconv.Itoa(l)
					s1 = append(s1, tmp)
				} else if nums[start] == l {
					tmp := strconv.Itoa(nums[start])
					s1 = append(s1, tmp)
				}
				start = i
			}
			l = z
		}
	}
	return s1
}

func main() {
	num := []int{-3, -2}
	summaryRanges(num)
	for _, i := range summaryRanges(num) {
		fmt.Println(i)
	}
}
