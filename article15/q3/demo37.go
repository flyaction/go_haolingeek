package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	dogP := &dog

	dogPtr := uintptr(unsafe.Pointer(dogP)) //内存起始存储地址 // uintptr一个无符号整数，能够存储指针值

	//unsafe.Offsetof函数用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位
	/*
		1.一个指针值（比如*Dog类型的值）可以被转换为一个unsafe.Pointer类型的值，反之亦然。
		2.一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值,反之亦然。
		3.一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此。
	*/

	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))

	fmt.Println(dogPtr)
	fmt.Println(nameP)
	fmt.Println(&(dogP.name))

	fmt.Printf("nameP == &(dogP.name)? %v\n", nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	// 示例2。
	// 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

}
