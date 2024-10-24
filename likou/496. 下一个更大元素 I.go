package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	sum := make([]int, 0)
	for _, v := range nums1 {
		k := 0
		for i, v1 := range nums2 {
			if v1 != v {
				continue
			} else {
				k = i
				break
			}
		}
		for j := k; j < len(nums2); j++ {
			if nums2[j] > v {
				sum = append(sum, nums2[j])
				break
			} else if j == len(nums2)-1 {
				sum = append(sum, -1)
				break
			}
		}
	}
	fmt.Println(sum)
	return sum
}

func main() {
	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	nextGreaterElement(nums1, nums2)
}
