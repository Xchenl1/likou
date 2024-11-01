package main

func distributeCandies(n int, limit int) int {
	var res int
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i+j > n {
				break
			} else if n-i-j <= limit {
				res++
			}

		}
	}
	return res
}

func main() {
	distributeCandies(3, 3)
}
