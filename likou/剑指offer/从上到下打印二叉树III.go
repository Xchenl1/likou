package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func Nixu(nums []int) []int {
	for i := 0; i < len(nums)/2; i++ {
		k := nums[i]
		nums[i] = nums[len(nums)-1-i]
		nums[len(nums)-1-i] = k
	}
	return nums
}

func levelOrder2(root *TreeNode) [][]int {
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
	for i := 0; i < len(s1); i++ {
		if i%2 != 0 {
			Nixu(s1[i])
		}
	}
	return s1
}

func main() {
	s4 := &TreeNode{5, nil, nil}
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, s3, nil}
	s1 := &TreeNode{1, s4, s2}
	fmt.Println(levelOrder2(s1))
}
