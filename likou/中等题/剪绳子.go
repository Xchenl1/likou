package main

import "fmt"

func cuttingRope(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	sum := 0
	for i := 1; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			if i*j*(n-i-j) > sum {
				sum = i * j * (n - i - j)
			}
		}
	}
	return sum
}

func main() {
	n := 3
	fmt.Println(cuttingRope(n))
}
