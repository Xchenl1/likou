package main

//	type ListNode struct {
//		Val  int
//		Next *ListNode
//	}
//
//	func reversePrint(head *ListNode) []int {
//		var nums []int
//		for head != nil {
//			nums = append(nums, head.Val)
//			head = head.Next
//		}
//		for i := 0; i < len(nums)/2; i++ {
//			m := nums[i]
//			nums[i] = nums[len(nums)-1-i]
//			nums[len(nums)-1-i] = m
//		}
//		return nums
//	}
//
// func main() {
//
// }
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reversePrint(root *ListNode) []int {
	a := make([]int, 0)
	for root != nil {
		a = append(a, root.Val)
		root = root.Next
	}
	for i := 0; i < len(a)/2; i++ {
		a[i] = a[len(a)-i-1]
	}
	return a
}

func main() {

}
