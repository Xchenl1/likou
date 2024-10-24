package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetric(root *TreeNode) bool {
	return duibi(root, root)
}
func duibi(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && duibi(p.Left, q.Right) && duibi(p.Right, q.Left)
}

func main() {
	//	s3 := &TreeNode{3, nil, nil}
	//	s2 := &TreeNode{2, s3, nil}
	//	s1 := &TreeNode{1, nil, s2}
	//	maxDepth(s1)
}
