package main

import "fmt"

func getRow(rowIndex int) []int {
	var a [][]int
	a = make([][]int, rowIndex+1)
	for i := 1; i <= rowIndex+1; i++ {
		a[i-1] = make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 {
				a[i-1][j] = 1
			} else if j == i-1 {
				a[i-1][j] = 1
			} else {
				a[i-1][j] = a[i-2][j-1] + a[i-2][j]
			}
		}
	}
	return a[rowIndex]
}

func main() {
	fmt.Println(getRow(4))
}
