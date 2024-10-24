package main

import "fmt"

//	func findContinuousSequence(target int) [][]int {
//		arr := make([]int, target-1)
//		var res [][]int
//		for i := 0; i < len(arr); i++ {
//			arr[i] = i + 1
//		}
//
//		var (
//			sum   = 0
//			left  = 0
//			right = 0
//		)
//
//		for right <= len(arr)-1 && left <= right {
//
//			sum = sum + arr[right]
//
//			if sum == target {
//				tmp := []int{}
//				for i := left; i <= right; i++ {
//					tmp = append(tmp, arr[i])
//				}
//				res = append(res, tmp)
//				sum = sum - arr[left]
//				left++
//			}
//			if sum > target {
//				sum = sum - arr[left]
//				sum = sum - arr[right]
//				left++
//			} else if sum < target {
//				right++
//			}
//		}
//
//		return res
//	}
func findContinuousSequence(target int) [][]int {
	//双指针法
	left := 0
	right := 0
	k := make([]int, 0)
	for i := 0; i < target; i++ {
		k = append(k, i+1)
	}
	sum := 0
	m := make([][]int, 0)
	for left <= right && right < len(k) {
		if sum == target {
			m = append(m, k[left:right])
			sum = sum - k[left]
			left++
		}
		if sum < target {
			sum += k[right]
			right++
		}
		if sum > target {
			sum -= k[left]
			left++
		}
	}
	return m
}
func main() {
	target := 15
	fmt.Println(findContinuousSequence(target))
}
