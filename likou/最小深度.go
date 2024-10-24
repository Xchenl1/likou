package main

type TreeNode11 struct {
	Val   int
	Left  *TreeNode11
	Right *TreeNode11
}

// 广度优先遍历
func minDepth(root *TreeNode11) int {
	if root == nil {
		return 0
	}
	//创建一个树结构类型的切片
	s1 := make([]*TreeNode11, 0)
	//创建一个int类型的切片方便存储深度
	sum := make([]int, 0)
	s1 = append(s1, root)
	//默认只有根节点 所以深度为1
	sum = append(sum, 1)
	for i := 0; i < len(s1); i++ {
		//如果发现为空  立马返回深度
		if s1[i].Left == nil && s1[i].Right == nil {
			return sum[i]
		}
		//左边不为空 那么需要将左边的节点放切片中
		if s1[i].Left != nil {
			s1 = append(s1, s1[i].Left)
			sum = append(sum, sum[i]+1)
		}
		//同理
		if s1[i].Right != nil {
			s1 = append(s1, s1[i].Right)
			sum = append(sum, sum[i]+1)
		}
	}
	return 0
}

func main() {
	s5 := &TreeNode11{1, nil, nil}
	s4 := &TreeNode11{4, s5, s5}
	s3 := &TreeNode11{3, s4, nil}
	s2 := &TreeNode11{2, s3, s5}
	minDepth(s2)
}
