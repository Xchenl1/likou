package main

import "fmt"

func arrangeCoins(n int) int {
	//其实就是前n项和
	i := 1
	for i <= n {
		if (i*(i+1))/2 > n {
			return i - 1
		}
		i++
	}
	return 0
}
func main() {
	n := 20
	fmt.Println(arrangeCoins(n))
}
