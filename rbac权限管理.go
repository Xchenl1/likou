package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, err := xormadapter.NewAdapter("mysql", "root:CHEN2580@tcp(127.0.0.1:3306)/e-commerce_system", true) //连接数据库
	if err != nil {
		fmt.Println(err)
	}
	e, err := casbin.NewEnforcer("./model.conf", a) //使用casbin
	if err != nil {
		fmt.Println(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		return
	}
	enforce, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		return
	}
	fmt.Println(enforce)
}
