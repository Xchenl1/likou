package main

import (
	"fmt"
	"gorm.io/gorm"
)

//type Student1 struct {
//	Id    uint    `gorm:"size:10"`
//	Name  string  `gorm:"size:16"`
//	Age   int     `gorm:"size:3"`
//	Email *string `gorm:"column:youxiang;size:128"`
//	Data  string  `gorm:"default:2016-11-09;comment:日期"`
//}

func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}

func main() {
	//DB.AutoMigrate(Student1{})
	////单行插入
	//s1 := Student1{
	//	Name: "ASD",
	//	Age:  18,
	//	Data: "2022-11-09",
	//}
	//err := DB.Create(&s1).Error
	//fmt.Println(err)
	//多行插入 用切片进行插入
	//var studentslist []Student1
	//for i := 0; i < 5; i++ {
	//	studentslist = append(studentslist, Student1{
	//		Name: fmt.Sprintf("陈令统%d", i+1),
	//		Age:  20 + i,
	//	})
	//}
	//err11 := DB.Create(&studentslist).Error
	//fmt.Println(err11)
	//单表查询
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysql.Logger("asda"),
	//})
	//var student2 Student1
	////DB.Take(&student2)
	////fmt.Println(student2)
	////DB.First(&student2)
	////fmt.Println(student2)
	//DB.Last(&student2)
	//fmt.Println(student2)
	////打印id为2的数据
	//var s2 Student1
	//DB.Take(&s2, 2)
	//fmt.Println(s2)
	//var s3 Student1
	//DB.Take(&s3, fmt.Sprintf("name='%s'", "陈令统1"))
	//fmt.Println(s3)
	//var studentlist []Student1
	//count := DB.Find(&studentlist).RowsAffected
	//fmt.Println(count)
	//for _, value := range studentlist {
	//	fmt.Println(value)
	//}
	////用json
	//DB.Find(&studentlist)
	//data, _ := json.Marshal(studentlist)
	//fmt.Println(string(data))
	//其他条件 主键和其他条件
	//DB.Find(&studentlist, []int{1, 2, 3})
	//DB.Find(&studentlist, 1, 2, 3, 4)
	//DB.Find(&studentlist, "name in ?", []string{"ASD"})
	//data, _ := json.Marshal(studentlist)
	//fmt.Println(string(data))

	//save更新
	//var student2 Student1
	//DB.Find(&student2, 11)
	//student2.Name = "陈令统NB"
	//student2.Age = 0
	//DB.Save(&student2)

	//批量更新
	//var student2 []Student1
	//DB.Find(&student2, []int{1, 2, 3}).Update("data", "2015-1-1")
	//更新多列

	//删除
	//var studnet2 Student1
	//DB.Delete(&studnet2, 1)
}
