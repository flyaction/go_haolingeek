package main

import (
	"fmt"
	"reflect"
)

var container = []string{"zero", "one1", "two"}

func main() {

	fmt.Print(reflect.TypeOf(container))
	fmt.Print("\n")

	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	fmt.Print(reflect.TypeOf(container))
	fmt.Print("\n")

	fmt.Printf("The element is %q.\n", container[1])
}
