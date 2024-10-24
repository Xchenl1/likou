package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func countNodes(root *TreeNode) int {
	num := 0
	if root == nil {
		return 0
	}
	// 处理最边界情况
	if root.Right == nil && root.Left == nil {
		return 1
	}
	if root.Left != nil {
		num += countNodes(root.Left)
	}
	if root.Right != nil {
		num += countNodes(root.Right)
	}
	num++
	return num
}

func main() {
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3, Left: n6}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	fmt.Println(countNodes(n1))
}
