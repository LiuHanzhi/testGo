package main

import (
	"flag"
	"fmt"
)

var myVar string

func init() {
	flag.StringVar(&myVar, "myVar", "", "myVar")
}

func main() {
	flag.Parse()

	fmt.Println("myVar:", myVar)
}
