package main

func firstUniqChar(s string) byte {
	//map不行因为便利的顺序时随机的
	k1 := make(map[int32]int, 0)
	for _, v := range s {
		k1[v]++
	}
	for _, v := range s {
		if k1[v] == 1 {
			return byte(v)
		}
	}
	return ' '
}
func main() {
	s := "abaccdeff"
	firstUniqChar(s)
}
