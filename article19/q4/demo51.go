package main

import "fmt"

func main() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}

/*
	go语言会把它携带的defer函数及其参数值另行存储到一个队列中。
	这个队列与该defer语句所属的函数是对应的，并且，他是先进后出(FILO)的，相当于一个栈

*/
