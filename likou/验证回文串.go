package main

import "fmt"

func isPalindrome(s string) bool {
	a := []byte(s)
	i := 0
	for _, value := range a {
		if (value >= 48 && value <= 57) || (value >= 97 && value <= 122) {
			a[i] = value
			i++
		} else if value >= 65 && value <= 90 {
			a[i] = value + 32
			i++
		} else {
			continue
		}
	}
	for j := 0; j < i; j++ {
		fmt.Printf(string(a[j]))
		//fmt.Println(j)
	}
	fmt.Println(i)
	if i%2 == 0 {
		for j := 0; j < i/2; j++ {
			if a[j] == a[i-j-1] {
			} else {
				return false
			}
		}
	} else if i%2 != 0 {
		for j := 0; j <= i/2; j++ {
			if a[j] == a[i-j-1] {
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
