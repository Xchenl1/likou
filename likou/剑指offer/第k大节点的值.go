package main

import "fmt"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func kthLargest(root *TreeNode1, k int) int {
	p := make([]int, 0)
	var bianli func(node *TreeNode1)
	bianli = func(node *TreeNode1) {
		if node == nil {
			return
		}
		p = append(p, node.Val)
		bianli(node.Left)
		bianli(node.Right)
	}
	bianli(root)
	Paixu(p)
	return p[len(p)-k]
}

func Paixu(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		a := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				l := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = l
				a = true
			}
		}
		if a == false {
			return
		}
	}
}
func main() {
	s1 := &TreeNode1{1, nil, nil}
	s2 := &TreeNode1{2, s1, nil}
	s3 := &TreeNode1{4, s2, nil}
	fmt.Println(kthLargest(s3, 2))
}
