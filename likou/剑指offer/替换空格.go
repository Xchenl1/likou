package main

import "fmt"

func replaceSpace(s string) string {
	var result string
	for _, v := range s {
		//v 类型为int32
		//空格的ASCII
		if v == 32 {
			result += "%20"
		} else {
			result += string(v)
		}
	}
	return result
}

//func replaceSpace(s string) string {
//	//无法直接改变字符串  字符串不可变  通过字节数组表示 一旦创建就不可边 如果想改变可以转化为[]rune或者[]byte
//	s1 := []rune(s)
//	for i, v := range s1 {
//		if v == ' ' {
//			s1[i] = '\u0032' //不能这样写 很有可能会改变原来
//		}
//	}
//	return string(s1)
//}

func main() {
	s := "we are happy."
	fmt.Println(replaceSpace(s))
}
