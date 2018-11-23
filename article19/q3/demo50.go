package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("Enter defer function.")

		// recover函数的正确用法。
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	// recover函数的错误用法。
	fmt.Printf("no panic: %v\n", recover())

	// 引发panic。
	panic(errors.New("something wrong"))

	// recover函数的错误用法。
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}

/*
	panic
	停止当前函数执行
	一直向上返回，执行每一层的defer
	如果没有遇见recover，程序退出

	recover
	仅在defer调用中使用
	获取panic的值
	如果无法处理，可重新panic


	意料之中的：使用error 。如文件打不开
	意料之外的: 使用panic 。如数组越界

*/
