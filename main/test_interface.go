package main

import "fmt"

// 定义一个接口
type Greeter interface {
	Greet()
}

// 定义一个结构体
type Person struct {
	Name string
	Age  int
}

// 实现接口的方法
func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	var g Greeter
	g = &Person{Name: "Alice", Age: 30}
	g.Greet() // 输出: Hello, my name is Alice and I am 30 years old.
}
