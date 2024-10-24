package main

import "fmt"

func reverseBits(num uint32) uint32 {
	//充分利用特性& | ^ &:同一为一 |:有一为一 可以说是保留原来的数据位  ^:相同为0 不同为1
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		k := num & 1
		j := k << (31 - i)
		res = res | j
		fmt.Printf("res:%b\n", res)
		num >>= 1
		//fmt.Printf("num:%b\n", num)
	}
	return res
}

func main() {
	var a uint32 = 0b11111111111111111111111111111101
	reverseBits(a)
}
