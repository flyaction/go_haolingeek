package main

import "fmt"

var block = "package"

func main() {
	fmt.Printf("The block is %s.\n", block)
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
