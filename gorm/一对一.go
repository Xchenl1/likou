package main

//type User struct {
//	ID       uint
//	Name     string
//	Age      int
//	Gender   bool
//	UserInfo *UserInfo // 通过UserInfo可以拿到用户详情信息
//}
//
//type UserInfo struct {
//	UserID uint // 外键
//	user   User
//	ID     uint
//	Addr   string
//	Like   string
//}

func main() {
	//DB.AutoMigrate(&User{}, &UserInfo{})
	//添加
	//DB.Create(&User{
	//	ID:     1,
	//	Name:   "xiaochen",
	//	Age:    17,
	//	Gender: true,
	//	UserInfo: &UserInfo{
	//		Addr: "xxxxx",
	//		Like: "篮球",
	//	},
	//})

	//查询
	//var userinfo UserInfo
	//DB.Preload("User").Take(&userinfo)
	//fmt.Println(userinfo)
	//var user User
	//DB.Preload("UserInfo").Take(&user)
	//fmt.Println(user)

	//删除
	//var user User
	//DB.Take(&user)
	//DB.Select("UserInfo").Delete(&user, 1)
}
