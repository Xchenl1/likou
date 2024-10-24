package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	//双指针法 g是孩子 s是饼干
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	j := 0
	count := 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			count++
			j++
			i++
		} else if s[j] < g[i] {
			j++
		}
	}
	return count
}

func main() {
	g := []int{1, 2, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}
