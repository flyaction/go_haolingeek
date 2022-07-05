package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

//map的Value赋值
func main() {

	list = make(map[string]Student)

	student := Student{"action"}

	list["student"] = student
	//list["student"].Name = "xd" //错误 cannot assign to struct field list["student"].Name in map

	tempStudent := list["student"]
	tempStudent.Name = "xd"
	list["student"] = tempStudent

	fmt.Println(list["student"])

}
