package main

import "fmt"

var name string = "action"

func main() {

	var name int = 20

	{
		var name float64 = 3.45

		fmt.Print(name)

	}

	fmt.Print(name)

}
