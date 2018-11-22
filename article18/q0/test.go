package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("empty content")

	fmt.Println(err)
}
