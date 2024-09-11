package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func (p *Person) BirthDay() {
	p.Age++
}
func (p Person) BirthDayNormal() {
	p.Age++
}

func main() {
	person := Person{
		Age:  10,
		Name: "Hanzhi",
	}

	fmt.Println("before birthday:", person.Age) // 10

	person.BirthDay()
	fmt.Println("after birthday:", person.Age) // 11

	person.BirthDayNormal()
	fmt.Println("after birthday:", person.Age) // 还是11，因为是值类型，所以不会改变（go默认是值传递，使用了指针后才是引用传递）
}
