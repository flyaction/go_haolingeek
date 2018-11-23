package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Enter function main.")

	panic(errors.New("something wrong"))

	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")

}
