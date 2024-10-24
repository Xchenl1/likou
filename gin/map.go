package main

import "fmt"

func main() {
	//定义创建第一种方式
	var m map[string]string
	//分配空间
	m = make(map[string]string, 2)
	m["name"] = "zhangsan"
	fmt.Println(m)
	//第二种方式
	m1 := map[string]string{
		"name": "chen",
		"xing": "ling",
	}

	fmt.Println(m1["name"])
	fmt.Println(m1)

	//查找 v表示值 ok表示true or false
	v, ok := m1["name"]
	fmt.Println(v, ok)
	//如果不存在v1为空值 比如string为nil int为0  ok1为false
	v1, ok1 := m1["ll"]
	fmt.Println(v1, ok1)
	//删除
	delete(m1, "name")
	fmt.Println(m1)
	//遍历
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	//map扩展
	var student = []map[string]string{
		{
			"name":     "zhangsan",
			"yuanxiao": "hekeyuan",
		},
		{
			"name":     "lisi",
			"yuanxiao": "hekeyuan",
		},
	}
	fmt.Println(student)
	//写一个自动追加的map 切片中的每一个元素是map
	var q []map[string]string
	//	q=make([]map[string]string,3)
	for i := 1; i <= 3; i++ {
		q = append(q, map[string]string{
			"name":  fmt.Sprintf("%s%d", "张", i),
			"dizhi": "wojia" + fmt.Sprintf("%d", i),
		})
	}
	for k, v := range q {
		fmt.Println(k, v)
	}
	fmt.Println(q)

	//map的value是切片
	var m11 map[string][]string = map[string][]string{
		"likes": []string{"chang", "tiao", "rap"},
	}
	like2 := m11["likes"]
	fmt.Println(like2)
	//添加就是
	m11["xingming"] = []string{
		"asd",
		"asds",
		"asdaq",
	}

}
