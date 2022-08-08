package main

import "fmt"

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	//New("little pig").SetName("monster") // 不能调用不可寻址的值的指针方法。
	tmp := New("little pig")
	tmp.SetName("monster")

	// 示例2。例外  字典索引结果值不可寻址，但是可以使用自增自增表达式
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++

	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}

	map1["word"]++

	fmt.Println(map1) //map[the:0 word:1 counter:0]
}
