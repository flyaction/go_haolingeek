package main

import (
	"fmt"
	"github.com/cornelk/hashmap"
	"time"
)

func main() {

	t := time.Now()

	m := hashmap.New[string, int]()

	m.Set("name", 123)

	value, _ := m.Get("name")

	fmt.Println(value)

	elapsed := time.Since(t)

	fmt.Println(elapsed)

}
