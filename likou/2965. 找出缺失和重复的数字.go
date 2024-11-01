package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	mp := make(map[int]int, len(grid)*len(grid))
	for i := 1; i <= len(grid)*len(grid); i++ {
		mp[i] = 0
	}
	res := make([]int, 2)
	for _, v := range grid {
		for _, v1 := range v {
			mp[v1]++
		}
	}
	for i, v := range mp {
		if v == 2 {
			res[0] = i
		} else if v == 0 {
			res[1] = i
		}
	}
	return res
}

func main() {
	grid := [][]int{{1, 3}, {2, 2}}
	findMissingAndRepeatedValues(grid)
}
