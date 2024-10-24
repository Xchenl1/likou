package main

import (
	"fmt"
	"sort"
)

type TreeNode00 struct {
	Val   int
	Left  *TreeNode00
	Right *TreeNode00
}

//	func getMinimumDifference(root *TreeNode00) int {
//		node := make([]*TreeNode00, 0)
//		sum := 100000
//		node = append(node, root)
//		i := 0
//		for i < len(node) && node[i] != nil {
//			k := node[i]
//			node = append(node, k.Left, k.Right) //先添加到切片中
//			if k.Left != nil {
//				leftnum := node[2*i+1].Val //取左子树的值
//				cha := k.Val - leftnum
//				if cha < 0 { //变负数为正数
//					a := -cha
//					if a < sum {
//						sum = a
//					}
//				} else if cha < sum {
//					sum = cha
//				}
//			}
//			if k.Right != nil {
//				rightnum := node[2*i+2].Val
//				cha := k.Val - rightnum
//				if cha < 0 {
//					cha1 := -cha
//					if cha1 < sum {
//						sum = cha1
//					}
//				} else if cha < sum {
//					sum = cha
//				}
//			}
//			i++
//		}
//		return sum
//	}
func getMinimumDifference(root *TreeNode00) int {
	node := make([]*TreeNode00, 0)
	num := make([]int, 0)
	num = append(num, root.Val)
	node = append(node, root)
	i := 0
	for i < len(node) {
		k := node[i]
		if k.Left != nil {
			node = append(node, k.Left)
			num = append(num, k.Left.Val)
		}
		if k.Right != nil {
			node = append(node, k.Right)
			num = append(num, k.Right.Val)
		}
		i++
	}
	sort.Ints(num)
	min := 100000
	k := num[0]
	for i := 1; i < len(num); i++ {
		if num[i]-k < min {
			min = num[i] - k
		}
		k = num[i]
	}
	return min
}
func main() {
	//s6 := &TreeNode00{699, nil, nil}
	s5 := &TreeNode00{2, nil, nil}
	s4 := &TreeNode00{3, s5, nil}

	//s3 := &TreeNode00{384, nil, s5}

	s2 := &TreeNode00{1, nil, s4}

	fmt.Println(getMinimumDifference(s2))
}
