package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return n
	}
	if n == 1 {
		return n
	}
	i := 2
	num := make([]int, 0)
	num = append(num, 0, 1)
	for i <= n {
		nums := num[i-2] + num[i-1]
		num = append(num, nums%1000000007)
		i++
	}
	return num[len(num)-1]
}
func main() {
	fmt.Println(fib(45))
}
