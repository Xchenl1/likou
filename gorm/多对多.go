package main

//type Tag struct {
//	ID   uint
//	Name string
//	//Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
//}
//
//type Article struct {
//	ID    uint
//	Title string
//	Tags  []Tag `gorm:"many2many:article_tags;"`
//}
//type ArticleTag struct {
//	ArticleID uint `gorm:"primaryKey"`
//	TagID     uint `gorm:"primaryKey"`
//	CreatedAt time.Time
//}

//func (a *ArticleTag) BeforeCreate(tx *gorm.DB) (err error) {
//	a.CreatedAt = time.Now()
//	return nil
//}
//
//func main() {
//	DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
//	var article Article
//	DB.Preload("Tags").Take(&article, 2)
//	var tag []Tag
//	DB.Find(&tag, []int{1})
//	DB.Model(&article).Association("Tags").Replace(&tag)
//var tag []Tag
//DB.Find(&tag, []int{1, 2})
//DB.Create(&Article{
//	Title: "我无敌",
//	Tags:  tag,
//})
//var article Article
//DB.Preload("Tags").Take(&article, 1)
//fmt.Println(article)
//DB.SetupJoinTable(&Tag{}, "Tags", &ArticleTag{})
//err := DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
//fmt.Println(err)
//DB.SetupJoinTable(&ArticleTag{}, "articletag")
//DB.AutoMigrate(&Tag{}, &Article{})
////多对多添加
//DB.Create(&Article{
//	Title: "python",
//	Tags: []Tag{{
//		Name: "a11sd",
//	},
//	},
//})
//
////创建文章已有文章
//DB.Create(&Tag{
//	Name: "ASDAS",
//	Articles: []Article{
//		{
//			Title: "golang",
//		},
//	},
//})
////var tag Tag
////DB.Preload("articles").Take(&tag, 1)
////fmt.Println(tag)
//var article Article
//DB.Preload("tags").Take(&article, 1)
//fmt.Println(article)
//}
