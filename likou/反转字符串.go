package main

import "fmt"

func reverseString(s []byte) {
	//知道总长度
	for i := 0; i < len(s)/2; i++ {
		l := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = l
	}
	fmt.Println(s)
}

// 反转元音字母
func reverseVowels(s string) string {
	l1 := []byte(s)
	//k := 0
	//p := 0
	//for i := 0; i < len(l1); i++ {
	//	if (l1[i] == 'a' || l1[i] == 'e' || l1[i] == 'i' || l1[i] == 'o' || l1[i] == 'u' || l1[i] == 'A' || l1[i] == 'E' || l1[i] == 'I' || l1[i] == 'O' || l1[i] == 'U') && k == 0 {
	//		p = i
	//		k++
	//	} else if (l1[i] == 'a' || l1[i] == 'e' || l1[i] == 'i' || l1[i] == 'o' || l1[i] == 'u' || l1[i] == 'A' || l1[i] == 'E' || l1[i] == 'I' || l1[i] == 'O' || l1[i] == 'U') && k > 0 {
	//		q := l1[p]
	//		l1[p] = l1[i]
	//		l1[i] = q
	//		p = i
	//	}
	//}
	for i := 0; i < len(l1)/2; i++ {
		if (l1[i] == 'a' || l1[i] == 'e' || l1[i] == 'i' || l1[i] == 'o' || l1[i] == 'u' || l1[i] == 'A' || l1[i] == 'E' || l1[i] == 'I' || l1[i] == 'O' || l1[i] == 'U') && (l1[len(l1)-i-1] == 'a' || l1[len(l1)-i-1] == 'e' || l1[len(l1)-i-1] == 'i' || l1[len(l1)-i-1] == 'o' || l1[len(l1)-i-1] == 'u' || l1[len(l1)-i-1] == 'A' || l1[len(l1)-i-1] == 'E' || l1[len(l1)-i-1] == 'I' || l1[len(l1)-i-1] == 'O' || l1[len(l1)-i-1] == 'U') {
			l := l1[i]
			l1[i] = l1[len(l1)-i-1]
			l1[len(l1)-i-1] = l
		}
	}
	l2 := string(l1)
	return l2
}

func main() {
	l := "hello"
	l1 := []byte(l)
	fmt.Println(l1)
	reverseString(l1)

	fmt.Println(reverseVowels(l))
}
