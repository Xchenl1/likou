package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 因为之前做过一道题 先用切片存储每个节点的地址
//func levelOrder(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	s := make([]int, 0)       //存储每层节点值
//	k := make([]*TreeNode, 0) //用于存储树的地址
//	k = append(k, root)
//	for len(k) != 0 {
//		if k[0] != nil {
//			s = append(s, k[0].Val)
//		}
//		if k[0].Left != nil {
//			k = append(k, k[0].Left)
//		}
//		if k[0].Right != nil {
//			k = append(k, k[0].Right)
//		}
//		k = k[1:]
//	}
//	return s
//}

//func main() {
//	s3 := &TreeNode{3, nil, nil}
//	s2 := &TreeNode{2, s3, nil}
//	s1 := &TreeNode{1, nil, s2}
//	fmt.Println(levelOrder(s1))
//}
