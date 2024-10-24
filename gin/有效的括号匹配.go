package main

import "fmt"

func isValid(s string) bool {
	sum := len(s)
	if sum%2 == 1 {
		return false
	}
	var a []byte
	a = make([]byte, sum+1)
	k := 1
	j := 1
	a[1] = s[0]
	//用栈的方法做
	for j < sum {
		if (a[k] == '(') && (s[j] == ')') {
			k--
			j++
		} else if (a[k] == '[') && (s[j] == ']') {
			k--
			j++
		} else if (a[k] == '{') && (s[j] == '}') {
			k--
			j++
		} else {
			a[k+1] = s[j]
			k++
			j++
		}
	}
	if k != 0 {
		return false
	}
	return true
}
func main() {
	var s1 string = "()[]{}"
	fmt.Println(isValid(s1))
}
