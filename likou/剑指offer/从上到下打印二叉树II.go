package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	s1 := make([][]int, 0)
	//s := make([]int, 0)       //存储每层节点值
	k := make([]*TreeNode, 0) //用于存储树的地址
	k = append(k, root)
	j := 0
	//答案是双切片存节点值
	for len(k) > 0 {
		s1 = append(s1, []int{})
		k1 := make([]*TreeNode, 0)
		for i := 0; i < len(k); i++ {
			//添加到s1[j]这一层中
			s1[j] = append(s1[j], k[i].Val)
			//将下边的节点存入切片中
			if k[i].Left != nil && k[i].Right != nil {
				k1 = append(k1, k[i].Left)
				k1 = append(k1, k[i].Right)
			} else if k[i].Left != nil && k[i].Right == nil {
				k1 = append(k1, k[i].Left)
			} else if k[i].Right != nil && k[i].Left == nil {
				k1 = append(k1, k[i].Right)
			}
		}
		j++
		k = k1
	}
	return s1
}

func main() {
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, s3, nil}
	s1 := &TreeNode{1, nil, s2}
	fmt.Println(levelOrder1(s1))
}
