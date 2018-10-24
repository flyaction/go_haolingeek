package main

import (
	"fmt"
	"reflect"
)

var container = []string{"zero", "one", "two"}

func main() {

	value, ok := interface{}(container).(map[int]string)
	//value,ok := interface{}(container).([]string)

	fmt.Print(reflect.TypeOf(value))

	fmt.Print("\n")

	fmt.Print(reflect.TypeOf(ok))

	fmt.Print("\n")

	fmt.Print(ok)

	fmt.Print("\n")

	fmt.Print(value)

}
