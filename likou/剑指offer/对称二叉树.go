package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isSymmetric(root *TreeNode) bool {
// 不用这种方法写 用迭代的方法写
//
//	if root == nil {
//		return true
//	}
//
//	if root != nil && root.Left == nil && root.Right == nil {
//		return true
//	}
//
// var k func(node *TreeNode, node1 *TreeNode) bool
//
//	k = func(node *TreeNode, node1 *TreeNode) bool {
//		if node.Val == node1.Val {
//			return true
//		}
//		if node.Left != nil && node1.Right != nil {
//			k(node.Left, node1.Right)
//		} else if node.Right != nil && node1.Left != nil {
//			k(node.Right, node1.Left)
//		} else if node.Left == nil && node1.Right != nil || node.Right != nil && node1.Left == nil {
//			return false
//		}
//	}
//
// k(root.Left, root.Right)
//
//		p := root
//		q := root
//		node := make([]*TreeNode, 0) //切
//		node = append(node, p, q)
//		for len(node) > 0 {
//			k1 := node[0]
//			k2 := node[1]
//			node = node[2:]
//			if k1 == nil && k2 == nil {
//				continue
//			}
//			if k1 == nil || k2 == nil {
//				return false
//			}
//			if k1.Val != k2.Val {
//				return false
//			}
//			node = append(node, k1.Left, k2.Right)
//			node = append(node, k1.Right, k2.Left)
//		}
//		return true
//	}
func isSymmetric(root *TreeNode) bool {
	k := make([]*TreeNode, 0)
	k = append(k, root, root)
	for len(k) > 0 {
		a := k[0]
		b := k[1]
		k = k[2:]
		if a == nil && b == nil {
			continue
		}
		if a == nil && b != nil {
			return false
		}
		if b == nil && a != nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		k = append(k, a.Left, b.Right)
		k = append(k, a.Right, b.Left)
	}
	return true
}
func main() {
	s4 := &TreeNode{5, nil, nil}
	s3 := &TreeNode{2, nil, s4}
	s2 := &TreeNode{2, s4, nil}
	s1 := &TreeNode{1, s2, s3}
	fmt.Println(isSymmetric(s1))
}
