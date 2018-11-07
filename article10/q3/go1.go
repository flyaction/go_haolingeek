package main

import "fmt"

func main() {
	go fmt.Println("1")
	fmt.Println("2")
}
