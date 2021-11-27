package main

import (
	"fmt"
)

func main() {
	fmt.Println("字符串基本类型")
	StringType()
	fmt.Println("字符串修改")
	StringModify()
	StringIter()
	res := StatChineseLength("yuanding袁丁")
	fmt.Printf("中文字符长度: %d\n", res)
}

/*
	关于字符串修改
		1. 字符串一旦被定义不可修改
	修改需要转换成数组操作
		1. 含有中文字符的需要定义成 `rune` 类型
		2. 英文字符定一次 `byte` 类型
		3. rune 实际上是 int32 的别名
*/
func StringModify() {
	s1 := "袁丁"
	// s[1] = '元'  无法直接修改字符串
	s2 := []rune(s1)
	s2[0] = '元'
	fmt.Println(string(s2))
	s3 := "abcdef"
	s4 := []byte(s3)
	fmt.Println(s3)
	fmt.Println(s4)
	s4[1] = 'h'
	fmt.Println(string(s4))
}

/*
	关于字符串类型
		双引号表示一个字符串
		单引号表示单个字符
		单独的汉字也是一个字符，只是字符多占了几个字节
		1 字节 = 8 bit
		多行字符串使用 `` 包裹
*/
func StringType() {
	s1 := "yuanding"
	s2 := "袁丁"
	s3 := "袁"
	s4 := '袁'
	fmt.Printf("%v\t%v\t%v\t%v\n", s1, s2, s3, s4)
	fmt.Printf("%T\t%T\t%T\t%T\n", s1, s2, s3, s4)

}

/*
	strings.Split : 分割   返回一个数组
	strings.Contains	: 包含 返回bool
	strings.HasPrefix	: 前缀匹配 返回bool
	strings.HasSuffix	: 后缀匹配 返回bool
*/
func _() {}

/*
	关于字符串迭代:
		1. `len()` 函数仅返回字符串所占用的字节长度，而非字符长度
		2. 使用 `for _,c := range STRING` 可以逐字符迭代
*/
func StringIter() {
	s1 := "yuanding袁丁"
	fmt.Println(len(s1)) //14 单个英文字符一个字节，每个中文字符三个字节
	fmt.Println("逐字节迭代，循环次数 14:")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("type: %T,value: %v,char: %c\n", s1[i], s1[i], s1[i])
	}
	fmt.Println("逐字符迭代 循环次数 10:")
	for _, c := range s1 {
		fmt.Printf("type: %T,value: %v,char: %c\n", c, c, c)
	}
}

/*
	如何判断字符串中汉字的个数？
		1.初始化一个计数器
		2.找出汉字的编码范围
		3.逐字符输出字符判断是否在范围内
		4.返回计数器值
*/
func StatChineseLength(a string) int {
	var startCode = int32(0x2e80)
	var endCode = int32(0x9fff)
	var count = 0
	for _, c := range a {
		if startCode <= c && c <= endCode {
			count++
		}
	}
	return count
}
