package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		//为什么要减一 因为用'A'+0 表示A这个字符  'A'+25:表示'Z'这个字符
		columnNumber--
		//P := 'A' + 0
		//fmt.Println(string(P))
		ans = append(ans, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	//b := string(ans)
	//fmt.Println(b)
	return string(ans)
}

func main() {
	fmt.Println(convertToTitle(701))
}
