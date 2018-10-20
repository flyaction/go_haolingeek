package main

import (
	"flag"
	"fmt"
)

var name string
var age int


func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object name.")
	flag.IntVar(&age, "age", 23, "The greeting object age.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %d!\n",age)
}
