package main

import "fmt"

func main() {
	fmt.Println("hello go")
	fmt.Printf("hello %s \n", "go")
	err := fmt.Errorf("error %s", "go")
	fmt.Println(err)
}
