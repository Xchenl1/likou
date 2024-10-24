package main

import "sort"

//	func removeDuplicates(nums []int) int {
//		//if len(nums) == 1 {
//		//	return len(nums)
//		//} else if len(nums) == 2 {
//		//	if nums[0] == nums[1] {
//		//		nums = nums[:1]
//		//		return len(nums)
//		//	} else {
//		//		return len(nums)
//		//	}
//		//}
//		sum := -1
//		//长度大于2
//		for i := 0; i <= len(nums)-1; i++ {
//			if nums[i] == sum {
//				nums = append(nums[:i], nums[i+1:]...)
//				i = i - 1
//				continue
//			}
//			sum = nums[i]
//		}
//		return sum
//	}
func removeDuplicates(nums []int) int {
	m := make(map[int]int) //todo:通过map来去重数组
	for i, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = i
	}
	//fmt.Println(m)
	nums1 := make([]int, 0) //todo:然后取出key 按照升序排序
	for i, _ := range m {
		nums1 = append(nums1, i)
	}
	sort.Ints(nums1)
	for i, v := range nums1 { //todo:将前len(nums1)个元素替换为nums1的元素
		nums[i] = v
	}
	return len(nums1)
}
func main() {
	nums := []int{1, 1}
	removeDuplicates(nums)
}
