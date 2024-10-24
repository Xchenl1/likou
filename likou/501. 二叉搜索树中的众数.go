package main

type TreeNode0 struct {
	Val   int
	Left  *TreeNode0
	Right *TreeNode0
}

func findMode(root *TreeNode0) []int {
	var s, k func(node *TreeNode0)
	sum := 0
	base, max := 0, 0
	F := make([]int, 0)
	k = func(node *TreeNode0) {
		if node.Val == base {
			sum++
		} else {
			sum = 1
			base = node.Val
		}
		if sum == max {
			F = append(F, node.Val)
		} else if sum > max {
			max = sum
			F = []int{node.Val}
		}
	}
	s = func(node *TreeNode0) {
		if node == nil {
			return
		}
		s(node.Left)
		//由于中序遍历是左中右
		k(node)
		s(node.Right)
	}
	s(root)
	return F
}

func main() {
	s5 := &TreeNode0{1, nil, nil}
	s4 := &TreeNode0{1, s5, s5}
	s3 := &TreeNode0{2, s4, nil}
	s2 := &TreeNode0{2, s3, nil}
	findMode(s2)
}
