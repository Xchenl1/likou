package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Right, root.Left = root.Left, root.Right
	} else if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
	} else if root.Right != nil {
		root.Left = root.Right
		root.Right = nil
	} else if root.Left == nil && root.Right == nil && root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil && root != nil {
		return root
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

func main() {
	s4 := &TreeNode{5, nil, nil}
	//s3 := &TreeNode{3, nil, nil}
	//s2 := &TreeNode{2, s3, nil}
	//s1 := &TreeNode{1, s4, s2}
	mirrorTree(s4)
}
