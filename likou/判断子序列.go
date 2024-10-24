package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	//双指针法 s是子序列 t是总序列
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else if s[i] != t[j] {
			j++
		}
	}
	if i != len(s) {
		return false
	}
	return true
}
func main() {
	fmt.Println(isSubsequence("ac", "asc"))
}
