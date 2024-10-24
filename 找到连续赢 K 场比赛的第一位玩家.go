package main

func findWinningPlayer(skills []int, k int) int {
	mp := make(map[int]int)
	// 使用map做标记
	for i, v := range skills {
		mp[v] = i
	}
	num := 0
	// 遍历数组
	for i := 1; i < len(skills); i++ {
		if skills[0] < skills[i] {
			// 交换次序
			k1 := skills[0]
			skills[0] = skills[len(skills)-1]
			skills[len(skills)-1] = k1
			//在交换时检查一下
			if num == k {
				return mp[skills[0]]
			}
			num = 0
		} else {
			num++
		}
		if i == len(skills)-1 {
			return mp[skills[0]]
		}

	}
	return num
}

func main() {
	num := []int{4, 2, 6, 3, 9}
	findWinningPlayer(num, 2)
}
