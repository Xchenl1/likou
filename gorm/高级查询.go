package main

import "fmt"

//type Student struct {
//	ID     uint   `gorm:"size:3"`
//	Name   string `gorm:"size:8"`
//	Age    int    `gorm:"size:3"`
//	Gender bool
//	Email  *string `gorm:"size:32"`
//}

func PtrString(email string) *string {
	return &email
}
func main() {
	//首先可以自己生成表结构
	DB.AutoMigrate(Student{})
	//DB.AutoMigrate(&Student{})
	//var studentList []Student
	//DB.Find(&studentList).Delete(&studentList)
	//studentList = []Student{
	//	{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
	//	{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
	//	{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
	//	{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
	//	{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
	//	{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
	//	{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
	//	{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
	//	{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	//}
	//DB.Create(&studentList)
	var users []Student
	//// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名不是枫枫的
	//DB.Where("name <> ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名包含 如燕，李元芳的
	//DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	//fmt.Println(users)
	//// 查询姓李的
	//DB.Where("name like ?", "李%").Find(&users)
	//fmt.Println(users)
	//// 查询年龄大于23，是qq邮箱的
	//DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	//fmt.Println(users)
	//// 查询是qq邮箱的，或者是女的
	//DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	//fmt.Println(users)
	//结构体查询
	// 会过滤零值
	//DB.Where(&Student{Name: "李元芳", Age: 0}).Find(&users)
	//fmt.Println(users)
	//map不会过虑0值 不会输出结果
	DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
	// SELECT * FROM `students` WHERE `age` = 0 AND `name` = '李元芳'
	fmt.Println(users)
	// 排除年龄大于23的
	DB.Not("age > 23").Find(&users)
	fmt.Println(users)
	DB.Or("gender = ?", false).Or(" email like ?", "%@qq.com").Find(&users)
	fmt.Println(users)
	DB.Select("name", "age").Find(&users)
	fmt.Println(users)
	// 没有被选中，会被赋零值
	var agelist []int
	DB.Model(Student{}).Select("distinct age").Scan(&agelist)
	fmt.Println(agelist)
}
