package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	// 只有一个根节点
	if root == nil {
		return 0
	}
	num := 0
	// 统计左边叶子
	if root.Left != nil {
		// 左叶子
		if root.Left.Left == nil && root.Left.Right == nil {
			num += root.Left.Val
		} else {
			num += sumOfLeftLeaves(root.Left)
		}
	}
	// 确保右边下个节点不是叶子
	if root.Right != nil && !(root.Right.Left == nil && root.Right.Right == nil) {
		num += sumOfLeftLeaves(root.Right)
	}
	return num
}

func main() {
	n5 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 15}
	n3 := &TreeNode{Val: 20, Left: n4, Right: n5}
	n2 := &TreeNode{Val: 9}
	n1 := &TreeNode{Val: 3, Left: n2, Right: n3}
	fmt.Println(sumOfLeftLeaves(n1))
}
