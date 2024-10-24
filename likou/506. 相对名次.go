package main

import (
	"sort"
	"strconv"
)

// 结构体排序
func findRelativeRanks(score []int) []string {
	var desc = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	type pair struct{ score, idx int } //记录分数和下标
	arr := make([]pair, len(score))
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score //从大到小排序
	})
	ans := make([]string, len(score))
	for i, p := range arr {
		if i < 3 {
			ans[p.idx] = desc[i]
		} else {
			ans[p.idx] = strconv.Itoa(i + 1)
		}
	}
	//fmt.Println(arr)
	return ans
}

func main() {
	score := []int{10, 3, 8, 9, 4}
	findRelativeRanks(score)
}
