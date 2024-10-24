package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//	if (p == nil && q != nil) || (p != nil && q == nil) {
//		return false
//	}
//	if p.Val != q.Val {
//		return false
//	}
//	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
//}

//func main() {
//	s3 := &TreeNode{3, nil, nil}
//	s2 := &TreeNode{2, s3, nil}
//	s1 := &TreeNode{1, nil, s2}
//
//	//s33 := &TreeNode{3, nil, nil}
//	s22 := &TreeNode{1, nil, nil}
//	s11 := &TreeNode{1, nil, s22}
//	fmt.Printf("%v", isSameTree(s1, s11))
//}
