package main

import (
	"fmt"
	"sort"
)

// 1.实现sort接口 2.匿名函数

type animal struct {
	Name   string
	Age    int
	Gender int
}
type animal1 []animal

func (a animal1) Len() int {
	return len(a)
}

func (a animal1) Less(i, j int) bool {

	if a[i].Age == a[j].Age {
		return a[i].Gender > a[j].Gender
	}
	return a[i].Age < a[j].Age
}

func (a animal1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	//var animal2 animal1
	var animal2 animal1
	//animal2 = []animal{
	//	{"小狗", 10},
	//	{"小猪", 11},
	//	{"小鹿", 12},
	//	{"小猫", 10},
	//}
	animal2 = []animal{
		{"小狗", 10, 1},
		{"小猪", 11, 0},
		{"小鹿", 12, 1},
		{"小猫", 10, 0},
	}
	//animal2 := func(a animal) animal1 {
	//	var animal3 []animal
	//	animal3 = append(animal3, a)
	//	return animal3
	//}
	//fmt.Println(animal2)
	//升序
	sort.Sort(animal2)
	for i, value := range animal2 {
		fmt.Print(i + 1)
		fmt.Println(value)
	}
	//
	//sort.Slice(animal3, func(i, j int) bool {
	//	return animal3[i].Age < animal3[j].Age
	//})
	//fmt.Println(animal3)
	//sort.SliceStable(animal3, func(i, j int) bool {
	//	return animal3[i].Age < animal3[j].Age
	//})
	//fmt.Println(animal3)
}
