package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age int

// 方式3。
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式3。
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object name.")
	cmdLine.IntVar(&age, "age", 29, "The greeting object age.")
}

func main() {
	// 方式3。
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %d!\n", age)
}

//go run test2.go -name="xd" -age=18