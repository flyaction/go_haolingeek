package main

import "fmt"

type handle func(str string)

func exec(f handle) {
	f("hello")
}

func main() {

	var p = func(str string) {
		fmt.Println("first", str)
	}
	exec(p)

	exec(func(str string) {
		fmt.Println("second", str)
	})

}
