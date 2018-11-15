package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string     //值方法
	Category() string //值方法
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// 示例2。

	dog.SetName("pig")
	fmt.Println(dog.Name())
	fmt.Println(dog.Category())

	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n", pet.Category(), pet.Name())
}
