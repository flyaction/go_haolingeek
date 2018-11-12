package main

import "fmt"

type Language struct {
	name   string
	ranges string
}

//func (ll Language) String() string  {
//	return fmt.Sprintf("%s%s",ll.name,ll.ranges)
//}

func (ll Language) String() string {
	return ll.name
}

func main() {
	l := Language{
		name:   "php",
		ranges: "all worlds",
	}

	fmt.Println(l)
}
