package main

import "fmt"

// 这种写法的整体思路是取四个变量 然后将他们分别记录 顺时针依次打印最外层的数据
// 然后进入内层
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	left := 0
	top := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1
	nums := make([]int, 0)
	for left <= right && top <= bottom {
		//先顺时针将最外层的数据存入切片中  首先是第一层
		for i := left; i <= right; i++ {
			nums = append(nums, matrix[top][i])
		}
		//这里有一个问题 会出现重复的情况从上到下遍历
		for i1 := top + 1; i1 <= bottom; i1++ {
			nums = append(nums, matrix[i1][right])
		}
		if left < right && top < bottom {
			//从右到左遍历
			for i := right - 1; i >= left; i-- {
				nums = append(nums, matrix[bottom][i])
			}
			//以及从下到上遍历
			for i := bottom - 1; i > 0; i-- {
				nums = append(nums, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return nums
}

// 123
// 456
// 789
func main() {
	obj := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(obj))
}
