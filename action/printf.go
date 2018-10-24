package main

import "fmt" //一定不要忘了

//printf格式化 总结

type point struct {
	x, y int
}

func test(i, j int) int { return i + j }

func main() {
	p := point{1, 2}

	fmt.Printf("%d\n", p) // {1 2}

	fmt.Printf("%+v\n", p) // {x:1 y:2}

	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}

	//输出类型
	fmt.Printf("%T\n", p) // main.point

	//输出函数签名
	fmt.Printf("%T\n", test) //func(int ,int) int

	//输出bool值
	flag := true
	fmt.Printf("%t\n", flag) // true

	//尝试将一个字符串作为参数来输出bool值，不要尝试这样做
	fmt.Printf("%t\n", "true") //%!t(string=true)

	//输出十进制形式输出
	fmt.Printf("%d\n", 123) // 123

	//输出一个字符，参数对应ASCII码
	fmt.Printf("%c\n", 98) // b

	//输出一个整数的二进制形式的值
	fmt.Printf("%b\n", 98) // 1100010

	//输出一个字符的二进制形式的值
	fmt.Printf("%b\n", 'b') // 1100010

	//如果参数是数字，则以十六进制形式输出
	fmt.Printf("%x\n", 456) // 1c8

	//如果参数是字符串，则打印字符串的每一个字符的ASCII码
	fmt.Printf("%x\n", "hex this") // 6865782074686973

	//浮点数形式输出，注意小数位数为6位
	fmt.Printf("%f\n", 78.53) // 78.530000

	//注意这里前后不对应，不会报错，但是不会自动转换
	fmt.Printf("%d\n", 78.53) // %!d(float64=78.53)

	//科学计数法的形式，注意参数要为小数，不为小数，可以乘1.0
	fmt.Printf("%e\n", 123400000000.0) //1.234000e+11 注意参数为小数

	//科学计数法的形式，注意参数要为小数，不为小数，可以乘1.0
	fmt.Printf("%E\n", 123000000000.0) //1.234000E+11

	//输出字符串
	fmt.Printf("%s\n", "\"ddadjfaskdafjasfsaf") //"ddadjfaskdafjasfsaf

	//保留字符串两端的双引号
	fmt.Printf("%q\n", "\"dddddddd\"") // "\"dddddddd\""

	//输出指针（地址）的值
	fmt.Printf("%p\n", &p) //0xc420012090

	//最小宽度为6，默认右对齐，不足6位时，空格补全，超过6位时，不会截断
	fmt.Printf("|%6d|%6d|\n", 12, 1234567) // |    12|1234567|

	//最小6个宽度（包含小数点)，2位小数，超过6位时，不会截断
	fmt.Printf("|%6.2f|%6.2f|\n", 12, 222.333) // |%!f(int=    12)|222.33|

	//使用 - 表示左对齐
	fmt.Printf("|%-6.2f|%-6.2f|\n", 12.2, 3.33) //|12.20 |3.33  |

	//最小6个宽度，右对齐，不足6个时，空格补齐，超过6位时，不会截断
	fmt.Printf("|%6s|%6s|\n", "foo", "abcdefgg") //|   foo|abcdefgg|

	////最小6个宽度，右对齐，不足6个时，空格补齐，超过6位时，不会截断
	fmt.Printf("|%-6s|%-6s|\n", "foo", "abcdeefgth") //|foo   |abcdeefgth|

	//不会输出内容，相反，会将内容以字符串的形式返回
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s) //a string
}
