package main

import "fmt"

/*
	关于fmt.printf 格式化打印：
		%T	: 数据类型
		%v	: 变量值
		%x	: 16进制打印
		%o	: 8 进制打印
*/
func main() {
	n := int8(100)
	fmt.Printf("%T\n", n)

}
