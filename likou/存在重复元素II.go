package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	//用哈希表来解决问题
	k1 := map[int]bool{}
	//遍历数组 找到相同的数组值
	for i, v := range nums {
		if k1[v] == true {
			for i1 := 0; i1 < i; i1++ {
				//找到第一个索引值  判断两个值得差是否小于k
				if nums[i1] == v {
					if i-i1 <= k {
						return true
					}
					continue
				}
			}
		}
		k1[v] = true
	}
	return false
}

//	func containsNearbyDuplicate(num []int) bool {
//		k1 := map[int]int{}
//		for _, v := range num {
//			k1[v]++
//			if k1[v] >= 2 {
//				return true
//			}
//		}
//		return false
//	}
func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
}
