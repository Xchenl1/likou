package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 看了参考答案这道题目的思路是  先将节点的下一个val以及该节点的地址存储到map中(两个map解决)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == nil || q == nil {
		return nil
	}
	node := make(map[int]*TreeNode)
	k := make(map[int]bool)
	var s func(root *TreeNode)
	s = func(node1 *TreeNode) {
		if node1 == nil {
			return
		}
		if node1.Left != nil {
			node[node1.Left.Val] = node1
			s(node1.Left)
		}
		if node1.Right != nil {
			node[node1.Right.Val] = node1
			s(node1.Right)
		}
	}
	s(root)
	//将所有节点的val存储到map中之后就有判断利用 kmap来进行判断
	for p != nil {
		k[p.Val] = true
		p = node[p.Val]
	}
	for q != nil {
		if k[q.Val] == true {
			return q
		}
		q = node[q.Val]
	}
	return nil
}

func main() {
	s3 := &TreeNode{3, nil, nil}
	s2 := &TreeNode{2, s3, nil}
	s1 := &TreeNode{1, nil, s2}
	fmt.Println(lowestCommonAncestor(s1, s2, s3))
}
