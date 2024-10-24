package main

import "fmt"

//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 想法是在A树中找到与B树相同头节点
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	s := make([]*TreeNode, 0)
	s = append(s, A)
	for len(s) > 0 {
		p := s[0]
		s = s[1:]
		if p.Val == B.Val && Bijiao(p, B) {
			return true
		}
		if p.Left != nil {
			s = append(s, p.Left)
		}
		if p.Right != nil {
			s = append(s, p.Right)
		}
	}
	return false
}
func Bijiao(s1 *TreeNode, s2 *TreeNode) bool {
	if s2 == nil {
		return true
	}
	if s2 == nil || s1 == nil {
		return false
	}
	if s1.Val == s2.Val {
		return Bijiao(s1.Left, s2.Left) && Bijiao(s1.Right, s2.Right) //递归
	} else {
		return false
	}
}

func main() {
	s4 := &TreeNode{3, nil, nil}
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, s3, nil}
	s1 := &TreeNode{1, s2, s3}
	fmt.Println(isSubStructure(s1, s4))
}
