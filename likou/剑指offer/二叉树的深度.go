package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root != nil && root.Left == nil && root.Right == nil {
		return 1
	}
	nums := make([]int, 0)
	k1 := make([]*TreeNode, 0)
	nums = append(nums, 1)
	k1 = append(k1, root)
	for i := 0; i < len(k1); i++ {
		if k1[i] == nil {
			continue
		}
		if k1[i].Left != nil {
			k1 = append(k1, k1[i].Left)
			nums = append(nums, nums[i]+1)
		}
		if k1[i].Right != nil {
			k1 = append(k1, k1[i].Right)
			nums = append(nums, nums[i]+1)
		}
	}
	return nums[len(nums)-1]
}

func main() {
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, s3, nil}
	s1 := &TreeNode{1, nil, s2}
	fmt.Println(maxDepth(s1))
}
