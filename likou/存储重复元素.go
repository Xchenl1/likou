package main

import "fmt"

func containsDuplicate(nums []int) bool {
	//这个方法不太行  时间复杂度太高了
	//for i := 0; i < len(nums); i++ {
	//	for j := i; j < len(nums); j++ {
	//		if j == i {
	//			continue
	//		} else if nums[i] == nums[j] {
	//			return true
	//		}
	//	}
	//}
	//return false
	//可以使用哈希表来做 合理利用map函数的特点
	k := map[int]bool{}
	for _, v := range nums {
		if k[v] == true {
			return true
		}
		k[v] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 4, 3, 5}
	fmt.Println(containsDuplicate(nums))
}
