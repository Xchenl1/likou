package main

import "fmt"

// 一个结构体类型
type AnimalCategory struct {
	kingdom string // 界。 //如果Kingdom小写的话 表示非公开成员 只能在该包内访问
	phylum  string // 门。 //如果首字母大写 那么就可以在所有包中访问
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

// 类型方法
// 对于这种具体的函数 只有AnimalCategory类型的实例对象才可以使用该方法
func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

// Animal
// 多级嵌套的时候，最外层的会屏蔽内层的
type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。嵌入字段
	//String //引申，这也是嵌入字段 并且会引发屏蔽
}

// Animal1
// 同级嵌套多个结构体时，如果拥有同名字段或者方法，会报出编译错误。换句话说，拥有相似方法属性的结构体只能嵌套一个。
type Animal1 struct {
	Animal
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func main() {
	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: AnimalCategory{"123", "123", "1", "2", "3", "3", "12"},
	}
	fmt.Println(animal.Category())
	fmt.Printf("The animal: %s\n", animal)
}
func Test() {

}
