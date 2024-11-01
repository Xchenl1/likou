package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	var sum []int
	var a func(node *TreeNode)
	a = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum = append(sum, node.Val)
		a(node.Left)
		a(node.Right)
	}
	a(root)
	return sum
}

func main() {
	s4 := &TreeNode{4, nil, nil}
	s3 := &TreeNode{3, s4, nil}
	s2 := &TreeNode{2, s3, nil}
	fmt.Println(preorderTraversal(s2))
}
