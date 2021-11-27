package main

import "fmt"

/*
	四种赋值方式:
		1. 声明变量类型后赋值：	var x int; x = 1
		2. 声明变量类型并赋值：	var x int = 1
		2. 自动推导变量类型：	var x = 1
		3. 简略写法的自动推导(仅能在函数中使用)： x := 1
*/
/*
	多变量赋值：
		1. var x,y,z int = 1, 2, 3
		2. var x,y,z int; x,y,z = 1,2,3
		3. x,y,z != 1,2,3
		4. 常用于声明全局变量
			var (
				x int
				y int
				z int
			)
*/

// 此语句会报错: `syntax error: non-declaration statement outside function body`
// xxx := "woojd"

/*
	go 中变量声明必须使用
*/

/*
	fmt.print 家族:
		1.fmt.Print普通打印
		2.fmt.Printf("%s", "yuadning") 格式化打印
		3.fmt.Println 打印后自动追加换行符
*/
func main() {
	// 赋值方式
	var varInt_1 int       //0
	var varString_1 string // ‘’
	var varFloat_1 float32
	var isOk bool // false
	// 声明了但未赋值，会使用一个默认值
	fmt.Println(varInt_1, varString_1, varFloat_1)
	varInt_1 = 1
	fmt.Println(varInt_1)
	// 赋值方式2
	varInt_2 := 2
	fmt.Println(varInt_2)
	// 格式化赋值
	// 赋值方式3
	var varString_2 = "yuand"
	var varInt_3 int
	varInt_3 = 26
	var formatStr = "Name: %s, age: %d"
	var String_3 = fmt.Sprintf(formatStr, varString_2, varInt_3)
	fmt.Println(String_3)
}
