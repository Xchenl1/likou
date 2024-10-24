package main

import (
	"fmt"
)

// 通过深度优先的方式解决
//
//	func maxValue(grid [][]int) int {
//		i, j := 0, 0
//		sum := 0
//		var fun func(i, j, i1, j1 int) bool //左边是向下走 右边是向右走
//		fun = func(i, j, i1, j1 int) bool {
//			if i >= len(grid) || j1 >= len(grid[0]) {
//				return false
//			}
//			if grid[i][j] > grid[i1][j1] {
//				sum += grid[i][j]
//				return true
//			}
//			return false
//		}
//		for i < len(grid)-1 || j < len(grid[0])-1 {
//			if fun(i+1, j, i, j+1) && i+1 < len(grid[0]) {
//				i++
//				continue
//			} else {
//				j++
//				sum += grid[i][j]
//				continue
//			}
//		}
//		return sum + grid[0][0]
//	}
//
//	func maxValue(grid [][]int) int {
//		m := len(grid)
//		n := len(grid[0])
//
//		// 创建一个二维数组用于保存到达每个位置的最大值
//		dp := make([][]int, m)
//		for i := range dp {
//			dp[i] = make([]int, n)
//		}
//
//		// 初始化第一行和第一列的最大值
//		dp[0][0] = grid[0][0]
//		for i := 1; i < m; i++ {
//			dp[i][0] = dp[i-1][0] + grid[i][0]
//		}
//		for j := 1; j < n; j++ {
//			dp[0][j] = dp[0][j-1] + grid[0][j]
//		}
//
//		// 逐行逐列计算最大值
//		for i := 1; i < m; i++ {
//			for j := 1; j < n; j++ {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
//			}
//		}
//
//		return dp[m-1][n-1]
//	}
//
//	func maxValue(grid [][]int) int {
//		m, n := len(grid), len(grid[0])
//		f := make([][]int, m)
//		for i := range f {
//			f[i] = make([]int, n)
//		}
//		for i, g := range grid {
//			for j, x := range g {
//				if i > 0 {
//					f[i][j] = max(f[i][j], f[i-1][j])
//				}
//				if j > 0 {
//					f[i][j] = max(f[i][j], f[i][j-1])
//				}
//				f[i][j] += x
//			}
//		}
//		return f[m-1][n-1]
//	}
//
// 动态规划：先计算每个位置的最大值 比如要求右下角的最大值(只能向右或向下移动) 那么要先找到左边和上边的最大值 依次寻找
func maxValue(grid [][]int) int {
	k := make([][]int, len(grid))
	for i := range k {
		k[i] = make([]int, len(grid[0]))
	} //生成一个跟grid一摸一样的数组 数据全部为0
	k[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		k[i][0] = max(k[i][0], k[i-1][0])
		k[i][0] += grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		k[0][i] = max(k[0][i], k[0][i-1])
		k[0][i] += grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			k[i][j] = max(k[i-1][j], k[i][j])
			k[i][j] = max(k[i][j-1], k[i][j])
			k[i][j] += grid[i][j]
		}
	}
	return k[len(grid)-1][len(grid[0])-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(grid))
}
