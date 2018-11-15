package main

import "fmt"

func main() {

	//切片
	a := []int{1, 2, 3}[0:2]
	fmt.Println("切片", a)

	//字典
	b := map[int]string{1: "a"}[1]
	fmt.Println("字典", b)
}
