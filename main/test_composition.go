package main

import "fmt"

//在 Go 语言中，组合（Composition） 是一种代码复用的技术，它通过将一个结构体嵌套到另一个结构体中，允许外层结构体直接使用内层结构体的字段和方法。
//组合是 Go 语言中实现代码复用和模拟继承的主要方式。

// 定义一个基础结构体
type Animal struct {
	Name string
}

// 给基础结构体定义一个方法
func (a Animal) Speak() {
	fmt.Println("Animal speaks")
}

// 定义一个结构体，通过组合实现复用
type Dog struct {
	Animal // 嵌入结构体
	Breed  string
}

func main() {
	// 创建一个 Dog 实例
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// 直接访问嵌入的 Animal 结构体的字段和方法
	fmt.Println(d.Name)  // 输出: Buddy
	d.Speak()            // 输出: Animal speaks
	fmt.Println(d.Breed) // 输出: Golden Retriever
}
