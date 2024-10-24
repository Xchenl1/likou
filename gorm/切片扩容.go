package main

import "fmt"

type Cat struct {
	Name string
	Age  int64
}

func (c *Cat) String() string {
	return fmt.Sprintf("%#v", c)
}
func main() {
	//s6 := make([]int, 0)
	////s6 := make([]int, 1,10)
	//println(s6)
	//fmt.Printf("The capacity of s6: %d\n", cap(s6))
	//for i := 1; i <= 5; i++ {
	//	s6 = append(s6, i)
	//	println(s6)                                                   //一旦触发扩容，地址信息会变
	//	fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6)) //长度在稳定增加，但是容量会跳着增加
	//}
	//fmt.Println()

	//var b = [6]int{1, 2, 3, 4, 5, 6}
	//for key, value := range b {
	//	fmt.Printf("value的值：%d,value的地址：%x,切片该元素的地址：%x\n", value, &value, &b[key])
	//	//Value的值不会变，value 的地址不变
	//	//println(reflect.TypeOf(value))
	//	value += 1  //验证下会不会改变原，需要再次以理解”值传递“的概念
	//	b[key] += 1 //这样才会变
	//}
	//fmt.Printf("value的值：%d", b) //值没有加1
	Cat1 := Cat{
		Name: "abc1",
		Age:  12,
	}

	Cat2 := Cat{
		Name: "abc2",
		Age:  13,
	}

	cats := make([]Cat, 2) //这里定义的是指针接口体
	cats[0] = Cat1
	cats[1] = Cat2

	for _, cat := range cats {
		fmt.Printf("Cat:%+v\n", cat)
		cat.Age = 14 // 这里会改变数据
	}
	fmt.Printf("Cat:%+v\n", cats)
}
