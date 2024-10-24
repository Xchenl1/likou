package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	//这是所有的长度
	for i := 0; i < len(matrix); i++ {
		//由于是递增的 第一行的最后一个数据最大 进行比较
		if matrix[i][len(matrix[i])-1] == target {
			return true
		} else if matrix[i][len(matrix[i])-1] < target {
			continue
		} else if matrix[i][len(matrix[i])-1] > target {
			for j := len(matrix[i]) - 1; j >= 0; j-- {
				if matrix[i][j] == target {
					return true
				} else if matrix[i][j] < target {
					break
				}
			}
		}
	}
	return false
}

func main() {
	//obj := [][]int{{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}
	obj := [][]int{{-1, 3}}
	fmt.Println(findNumberIn2DArray(obj, -1))
}
