package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func hasPathSum(root *TreeNode, targetSum int) bool {
	//可以按照广度优先白遍历来写
	//1.如果根是空 直接返回false 因为不可能有相等的长度
	if root == nil {
		return false
	}
	var a []*TreeNode
	var b []int
	b = append(b, root.Val)
	k := len(a)
	a = append(a, root)
	for len(a) > 0 {
		k = len(a)
		if k > 0 {
			a1 := a[0]
			a = a[1:]
			b1 := b[0]
			b = b[1:]
			if a1.Left == nil && a1.Right == nil {
				if b1 == targetSum {
					return true
				}
			}
			if a1.Left != nil {
				a = append(a, a1.Left)
				b = append(b, a1.Left.Val+b1)
			}
			if a1.Right != nil {
				a = append(a, a1.Right)
				b = append(b, a1.Right.Val+b1)
			}
		}

	}
	return false
}

func main() {
	s4 := &TreeNode{1, nil, nil}
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, nil, s3}
	s1 := &TreeNode{1, s4, s2}
	hasPathSum(s1, 6)
}
