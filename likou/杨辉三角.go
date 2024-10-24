package main

func generate(numRows int) [][]int {
	//var a [][]int = [][]int{{1}, {1, 2}}
	//fmt.Println(a[1][1])
	var a [][]int
	//首先要开辟二维数组的个数
	a = make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		//再进行开辟数组的a[i]中有多少个元素
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
	return a
}
func main() {
	generate(4)
}
