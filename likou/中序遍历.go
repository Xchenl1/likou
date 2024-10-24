package main

//
//import "fmt"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func inorderTraversal(root *TreeNode) []int {
//	var a func(node *TreeNode)
//	var k []int
//	a = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		a(node.Left)
//		k = append(k, node.Val)
//		a(node.Right)
//	}
//	a(root)
//	return k
//}
//func main() {
//	s3 := &TreeNode{3, nil, nil}
//	s2 := &TreeNode{2, s3, nil}
//	s1 := &TreeNode{1, nil, s2}
//	fmt.Println(inorderTraversal(s1))
//}
