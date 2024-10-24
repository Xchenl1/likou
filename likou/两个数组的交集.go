package main

import "fmt"

// 因为是找交集
func intersection(nums1 []int, nums2 []int) []int {
	//用map
	m1 := map[int]int{}
	for _, v := range nums1 {
		m1[v]++
	}
	m2 := make([]int, 0)
	//用V2的每个值m1中的值相比较
	for _, v := range nums2 {
		if _, ok := m1[v]; ok == true {
			m2 = append(m2, v)
		}
	}
	return m2
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
