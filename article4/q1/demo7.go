package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	//var name string                                                   // [1]
	//flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	// 方式1。
	var name = *flag.String("name", "everyone", "The greeting object.") //类型*string代表的是字符串的指针类型，而不是字符串类型

	// 方式2。
	//name := *flag.String("name", "everyone", "The greeting object.")

	flag.Parse()

	fmt.Printf("Hello, %v!\n", name)
	fmt.Println()
	fmt.Println(reflect.TypeOf(name))
}
