package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func invertTree(root *TreeNode2) *TreeNode2 {
	if root == nil {
		return root
	}
	//这一步比较关键 同时翻转
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func qianxu(root *TreeNode2) []int {
	var a []int
	var v func(root *TreeNode2)
	v = func(root *TreeNode2) {
		if root == nil {
			return
		}
		a = append(a, root.Val)
		v(root.Left)
		v(root.Right)
	}
	v(root)
	return a
}

//func invertTree(root *TreeNode2) *TreeNode2 {
//	if root == nil {
//		return nil
//	}
//	root.Left, root.Right = root.Right, root.Left
//	invertTree(root.Left)
//	invertTree(root.Right)
//	return root
//}

// 反转二叉树
func main() {
	a3 := &TreeNode2{3, nil, nil}
	a11 := &TreeNode2{1, nil, nil}
	a2 := &TreeNode2{2, a11, a3}
	fmt.Println(qianxu(a2))
	p := invertTree(a2)
	fmt.Println(qianxu(p))
}
