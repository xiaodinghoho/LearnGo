package main

import "fmt"

func main() {
	aboutIf()
	aboutFor()
	multiplication()
}

func aboutIf() {
	if 0 > 1 {
		fmt.Println("GO")
	} else if 1 < 2 {
		age := 19
		fmt.Println(age)
	} else {
		fmt.Println("GO")
	}
	// 代码作用域的问题，age仅在if 子语句下起作用，所以函数体外部不起作用
	//fmt.Println(age)   // 报错写法
}

func aboutFor() {
	// 基本格式 初始值；判断条件；循环
	for i := 1; i < 10; i++ {
		fmt.Printf("%d  ", i)
	}
	fmt.Printf("\n")
	// 变种,注意作用域
	i := 5
	for ; i < 10; i++ {
		fmt.Printf("%d  ", i)
	}
	// 变种2,类似while
	var a = 7
	for a < 10 {
		fmt.Printf("%d  ", a)
		a++
	}
	// 无限循环
	// for {
	// 	fmt.Print("inf")
	// }

	//键值循环
	s := "abcdefgh"
	for i, v := range s {
		fmt.Printf("位置:%d 数值:%c\n", i, v)
	}

}

func multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d  ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
