package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
	fmt.Println(reflect.TypeOf(name))
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

//上面函数的实现也可以是这样的。
//func getTheFlag() *int {
//	return flag.Int("num", 1, "The number of greeting object.")
//}
