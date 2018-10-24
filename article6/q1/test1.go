package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// 方式1。
	_, ok1 := interface{}(container).([]string) //
	//_, ok2 := interface{}(container).(map[int]string)
	//if !(ok1 || ok2) {
	//	fmt.Printf("Error: unsupported container type: %T\n", container)
	//	return
	//}
	if !ok1 {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
}
