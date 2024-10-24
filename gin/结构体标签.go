package main

import (
	"encoding/json"
	"fmt"
)

type Artical struct {
	Title    string `json:"-"` //-也不参与序列化
	Free     bool   `json:"2"`
	Password string `json:"1"`
	name     string `json:"name"` //小写字母开头不会参与序列化
}

func main() {
	artiacl1 := Artical{
		Title:    "12",
		Free:     false,
		Password: "asd",
		name:     "chen",
	}
	//结构体转json
	jsonDate, error := json.Marshal(artiacl1)
	if error != nil {
		fmt.Println(error)
		return
	}
	jsonStr := string(jsonDate)
	fmt.Println(jsonStr)
	fmt.Println(artiacl1)
}
