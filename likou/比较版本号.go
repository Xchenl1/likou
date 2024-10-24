package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version11 := strings.Split(version1, ".")
	version22 := strings.Split(version2, ".")
	//fmt.Println(version11, version22)
	for i := 0; i < len(version11) || i < len(version22); i++ {
		x, y := 0, 0
		if i < len(version11) {
			x, _ = strconv.Atoi(version11[i])
		}
		if i < len(version22) {
			y, _ = strconv.Atoi(version22[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
func main() {
	fmt.Println(compareVersion("1.01", "1.001.1"))
}
