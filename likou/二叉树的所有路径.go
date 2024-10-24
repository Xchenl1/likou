package main

import (
	"strconv"
)

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

func binaryTreePaths(root *TreeNode3) []string {
	s1 := make([]string, 0)
	//采用匿名函数的方法
	var a func(root *TreeNode3, k string)
	a = func(root *TreeNode3, k string) {
		if root != nil {
			//加上节点值
			k += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				s1 = append(s1, k)
			} else {
				k += "->"
				a(root.Left, k)
				a(root.Right, k)
			}
		}
	}
	a(root, "")
	return s1
}

func main() {
	s5 := &TreeNode3{1, nil, nil}
	s4 := &TreeNode3{4, s5, s5}
	s3 := &TreeNode3{3, s4, nil}
	s2 := &TreeNode3{2, s3, nil}
	binaryTreePaths(s2)
}
