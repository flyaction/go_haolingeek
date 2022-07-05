package main

import "fmt"

type People struct {
	Name string
	Age  int
}

//map的遍历赋值
func main() {
	//定义map
	list := make(map[string]*People)

	//定义student数组
	peoples := []People{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "xi", Age: 22},
	}

	//遍历结构体数组，依次赋值给map
	for i := 0; i < len(peoples); i++ {
		list[peoples[i].Name] = &peoples[i]
	}

	//打印map
	for k, v := range list {
		fmt.Println(k, "=>", v.Name)
	}

}
