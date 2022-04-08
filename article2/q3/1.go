package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {

}

func main() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse() //解析参数
	fmt.Printf("Hello, %s!\n", name)
}
