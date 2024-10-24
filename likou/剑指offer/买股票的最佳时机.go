package main

func maxProfit(prices []int) int {
	min := 10000
	max := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > max {
			max = v - min
		}
	}
	return max
}
func main() {

}
