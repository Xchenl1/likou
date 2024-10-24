package main

import "fmt"

// 对于栈 想法非常巧妙
func validateStackSequences(pushed []int, popped []int) bool {
	s := make([]int, 0)
	j := 0
	for _, v := range pushed {
		s = append(s, v)
		for len(s) > 0 && popped[j] == s[len(s)-1] {
			s = s[:len(s)-1]
			j++
		}
	}
	return len(s) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}
