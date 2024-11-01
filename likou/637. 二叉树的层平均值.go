package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Num struct {
	nums  int // 统计每层值和
	count int // 统计每层个数
}

func averageOfLevels(root *TreeNode) []float64 {
	var nums []Num
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(nums) {
			nums = append(nums, Num{nums: node.Val, count: 1})
		} else {
			nums[level].count++
			nums[level].nums += node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	jieguo := make([]float64, 0)
	for _, v := range nums {
		s := float64(v.nums) / float64(v.count)
		jieguo = append(jieguo, s)
	}
	return jieguo
}
