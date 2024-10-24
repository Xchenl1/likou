package main

import (
	"fmt"
	"time"
)

type ArticleModel struct {
	ID    uint
	Title string
	Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
}

type TagModel struct {
	ID       uint
	Name     string
	Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
}

type ArticleTagModel struct {
	ArticleID uint `gorm:"primaryKey"` // article_id
	TagID     uint `gorm:"primaryKey"` // tag_id
	CreatedAt time.Time
}

func main() {
	DB.AutoMigrate(&ArticleModel{}, &TagModel{})
	DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
	DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
	DB.Create(&ArticleModel{
		Title: "asd",
		Tags: []TagModel{
			{Name: "xaiochen"},
			{Name: "花花"},
		},
	})
	var article []ArticleModel
	DB.Preload("Tags").Take(&article)
	fmt.Println(article)
}
