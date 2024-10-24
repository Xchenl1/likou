package main

func addDigits(num int) int {
	//这一步是通过这个判断sum/10!=0
	sum := 100
	for sum/10 != 0 {
		sum = 0
		for num != 0 {
			k := num % 10
			sum += k
			num /= 10
		}
		num = sum
	}
	return sum
}

func main() {
	num := 122
	addDigits(num)
}
