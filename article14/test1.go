package main

import (
	"fmt"
)

type Car struct {
	Name string
	Num  int
}

func (c Car) Move() {
	fmt.Println(c.Name, "--", c.Num)
}

type Xiaorui interface {
	Move()
}

func Test(x Xiaorui) {
	x.Move()
}

func main() {
	//结构体
	c := Car{Name: "aa", Num: 5}
	c.Move()

	//接口
	Test(c)

}
