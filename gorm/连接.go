package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	Id    uint    `gorm:"size:10"`
	Name  string  `gorm:"size:16"`
	Age   int     `gorm:"size:3"`
	Email *string `gorm:"column:邮箱;size:128"`
	Data  string  `gorm:"default:2016-11-09;comment:日期"`
}

var mysqlLogger logger.Interface

var DB *gorm.DB

func init() {
	username := "root"
	passwd := "CHEN2580"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gorm"
	Timeout := "10s" //10秒连接超时
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, passwd, host, port, Dbname, Timeout)

	mysqlLogger = logger.Default.LogMode(logger.Info)
	//连接数据库mysql 获得db类型实例，用于后边的操作  可以修改这些策略
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Info, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  true,        // 使用彩色打印
	//	},
	//)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	//NamingStrategy: schema.NamingStrategy{
	//	//	//TablePrefix:   "j_", //表明前缀
	//	//	//SingularTable: true, //是否单数表名
	//	//	//NoLowerCase:   true, //不要小写转换
	//	//	Logger: mysqlLogger,
	//	//},
	//	//自定义日志
	//	//Logger: newLogger,
	//	//默认日志
	//	Logger: mysqlLogger,
	//	//SkipDefaultTransaction: true,
	//})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败！" + err.Error())
	}
	//fmt.Println(db)
	// SELECT * FROM `students` ORDER BY `students`.`name` LIMIT 1
	//var model Student
	//session := DB.Session(&gorm.Session{Logger: newLogger})
	//session.First(&model)
	DB = db
	//var studenta Student
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	//DB.Take(&studenta)
	//fmt.Println(studenta)
	//db.First(&studenta)
	//fmt.Println(studenta)
	//DB.Last(&studenta)
	//fmt.Println(studenta)
	//fmt.Println(mysqlLogger)
	//连接成功
	//fmt.Println(db)
}

//
//func main() {
//	//第二种方式创建日志
//	DB = DB.Session(&gorm.Session{Logger: mysqlLogger})
//	//DB.AutoMigrate(&Student{})
//	//email := "asdasdasd"
//	//email1 := &email
//	//var studentt []Student
//	//for i := 2; i < 10; i++ {
//	//	studentt = append(studentt, Student{
//	//		uint(i),
//	//		fmt.Sprintf("陈+%d", i),
//	//		18 + i,
//	//		nil,
//	//		fmt.Sprintf("2023-2-%d", i),
//	//	})
//	//}
//	//DB.Create(studentt)
//	//在数据库中创建表 第三种方式显示日志
//	//DB.Debug().AutoMigrate(&Student{})
//	//fmt.Println(DB)
//	//var studentt Student
//	//DB.Find(&studentt)
//	//DB.Model(&Student{}).Where("age=?", 21).Update("Email", "起飞@qq.com")
//}
