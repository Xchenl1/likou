package main

//type User struct {
//	ID       uint      `gorm:"size:4"`
//	Name     string    `gorm:"size:8"`
//	Articles []Article // 用户拥有的文章列表
//}

//type Article struct {
//	ID     uint   `gorm:"size:4"`
//	Title  string `gorm:"size:16"`
//	UserID uint   `gorm:"size:4"` // 属于   这里的类型要和引用的外键类型一致，包括大小
//	User   User   // 属于
//}

func main() {
	//var user User
	//var article Article
	//DB.Take(&article, 1)
	//DB.Select("User").Delete(&user)
	//DB.AutoMigrate(&Article{}, &User{})
	//var user1 User
	//DB.Create(&User{
	//	ID:   1,
	//	Name: "小陈",
	//	Articles: []Article{{
	//		Title: "一千零一夜",
	//	}, {
	//		Title: "python",
	//	},
	//	},
	//})
	//创建文章关联用户
	//	DB.Create(&Article{
	//		Title: "阿帅",
	//		User: User{
	//			Name: "小张",
	//		},
	//	})
	//}
	//var user User
	//DB.Take(&user, 1)
	//var artical Article
	//DB.Take(&artical, 4)
	//DB.Model(&user).Association("Articles").Append(&artical)
	//给文章绑定用户

}
