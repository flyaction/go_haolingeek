package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "everyone", "The greeting object name.")

func init() {

}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
