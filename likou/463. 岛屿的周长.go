package main

import (
	"fmt"
)

//	func islandPerimeter(grid [][]int) int {
//		sum := 0
//		//l := 0
//		h := 0
//		l := 0
//		k := make([]int, len(grid))
//		for i := 0; i < len(grid); i++ {
//			for j := 0; j < len(grid[i]); j++ {
//				if grid[i][j] == 1 {
//					if i < len(grid)-1 && grid[i+1][j] == 1 {
//						h++ //下边有多少个1
//					}
//					k[i]++
//					l++ //记录每一层有多少个1
//				}
//			}
//			if l > 0 {
//				k[i] = k[i]*4 - (k[i]-1)*2
//			}
//			sum += k[i] - h*2
//			h = 0
//			l = 0
//		}
//		return sum
//	}
//
// type pair struct{ x, y int }
//
// var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上、下、左、右
// // 深度优先遍历
//
//	func islandPerimeter(grid [][]int) (ans int) {
//		n, m := len(grid), len(grid[0])
//		var dfs func(x, y int)
//		dfs = func(x, y int) {
//			if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
//				ans++
//				return
//			}
//			if grid[x][y] == 2 {
//				return
//			}
//			grid[x][y] = 2
//			for _, d := range dir4 {
//				dfs(x+d.x, y+d.y)
//			}
//		}
//		for i, row := range grid {
//			for j, v := range row {
//				if v == 1 {
//					dfs(i, j)
//				}
//			}
//		}
//		return
//	}
//
// 既然是深度优先遍历 那么要先定义向哪个方向走

func islandPerimeter(grid [][]int) int {
	Driect := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右
	n := len(grid)                                      //有多少层
	m := len(grid[0])                                   //每一层有多少 因为每一层都是相等的 我就取了第一层
	num := 0
	var XZ func(x, y int)
	XZ = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			num++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, v := range Driect {
			XZ(x+v[0], y+v[1])
		}
	}
	for i, H := range grid {
		for j, L := range H {
			if L == 1 {
				XZ(i, j)
			}
		}
	}
	return num
}
func main() {
	//grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	grid := [][]int{{1, 1, 1}, {1, 0, 1}}
	fmt.Println(islandPerimeter(grid))
}
