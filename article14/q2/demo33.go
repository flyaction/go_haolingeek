package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
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
	var dog1 *Dog
	if dog1 == nil {
		fmt.Println("The first dog is nil.")
	} else {
		fmt.Println("The first dog is not nil.")
	}

	dog2 := dog1
	if dog2 == nil {
		fmt.Println("The second dog is nil.")
	} else {
		fmt.Println("The second dog is not nil.")
	}

	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n", dog2)
	fmt.Printf("The type of second dog is %s.\n", reflect.TypeOf(dog2).String())

	fmt.Println(dog1, dog2, pet)
	fmt.Println()

	// 示例2。
	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}

	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
}
