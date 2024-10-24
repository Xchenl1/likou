package main

import "fmt"

// s 只包含小写字母 delete删除map中的数据  方法不可行 处理不了多个相同的字母(可以解决 用map存储数量 不存索引)

func firstUniqChar(s string) int {
	nums := make(map[uint8]int)
	nums[s[0]] = 1
	for i, v := range s {
		if i == 0 {
			continue
		}
		if _, ok := nums[uint8(v)]; ok {
			nums[s[i]]++
		} else {
			nums[s[i]]++
		}
	}
	for i, _ := range s {
		if nums[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
