package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	//先排序 从小到大排序
	sort.Slice(nums, func(i, j int) bool {
		x := nums[i]
		y := nums[j]
		return x < y
	})
	i := 0
	k := 0
	j := 0
	for i < len(nums) {
		//保存0的数量
		if nums[i] == 0 {
			j++
			i++
			continue
		}
		//判断第一个不是0的数字 存到k中
		if nums[i] != 0 && k == 0 {
			k = nums[i]
			i++
			continue
		}
		if nums[i] != k+1 && j == 0 {
			return false
		}
		if nums[i] != k+1 && j != 0 {
			j--
			k++
			continue
		}
		if nums[i] != 0 && k != 0 {
			k = nums[i]
		}
		i++
	}
	return true
}

func main() {
	nums := []int{0, 0, 1, 2, 5}
	fmt.Println(isStraight(nums))
}
