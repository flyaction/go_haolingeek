package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

// 方式3。
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式3。
	cmdLine.StringVar(&name, "name", "boys", "The greeting object name.")
	cmdLine.StringVar(&name, "name1", "girls", "The greeting object age.")
}

func main() {
	// 方式3。
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
