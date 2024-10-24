package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	k := make([]int, 0)
	k = append(k, 0, 1)
	for i := 2; i <= n; i++ {
		k = append(k, k[i-1]+k[i-2])
	}
	return k[len(k)-1]
}
func main() {
	fmt.Println(fib(4))
}
