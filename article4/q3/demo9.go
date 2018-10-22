package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。

	fmt.Print(reflect.TypeOf(err), "\n")
	fmt.Print(reflect.TypeOf(n), "\n")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
}
