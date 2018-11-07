package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	t := time.Now()

	fmt.Println(t)

	fmt.Println(t.Unix())

	fmt.Println(t.Weekday())

	fmt.Println(reflect.TypeOf(t))

	const base_format = "2006-01-02 15:04:05"
	str_time := t.Format(base_format)
	fmt.Println(str_time)

	fmt.Println(t.Format("2006"))

	fmt.Println(string([]byte(string(str_time)))[:10])

}
