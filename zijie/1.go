package main

import (
	"fmt"
	"sync"
)

// InnerStruct 定义嵌套内部结构体
type InnerStruct struct {
	InnerField string
}

// MyStruct 定义结构体
type MyStruct struct {
	Field1 string
	Field2 int
	Inner  InnerStruct
}

// ConcurrentPrint 并发打印多个结构体内容
func ConcurrentPrint(structs []MyStruct) {
	// 使用 sync.WaitGroup 控制并发
	var wg sync.WaitGroup
	wg.Add(len(structs))

	resultCh := make(chan string)

	// 启动协程，执行并发操作
	for _, s := range structs {
		go func(str MyStruct) {
			defer wg.Done()
			result := fmt.Sprintf("Field1: %s, Field2: %d, InnerField: %s", str.Field1, str.Field2, str.Inner.InnerField)
			resultCh <- result
		}(s)
	}

	// 打印管道中数据
	go func() {
		for result := range resultCh {
			fmt.Println(result)
		}
	}()

	wg.Wait()
	// 关闭管道
	close(resultCh)
}

func main() {
	// 初始化结构体
	myStructs := []MyStruct{
		{
			Field1: "Value 1",
			Field2: 2,
			Inner: InnerStruct{
				InnerField: "Inner Value 1",
			},
		},
		{
			Field1: "Value 3",
			Field2: 4,
			Inner: InnerStruct{
				InnerField: "Inner Value 2",
			},
		},
	}

	ConcurrentPrint(myStructs)
}
