package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外部必须要有 var 定义符 go没有全局变量只有包变量概念
var aa = 3
var ss = "ss"
var bb = true

// 第一种定义空值变量方式
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 第二种指定类型定义变量
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b , s)
}

// 第三种不指定类型 go 语言会自动推断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b ,c, s)
}

// 第四种使用 := 符号快捷定义变量 但只能在函数内部使用
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/* ---------------------------------------------------------------------------- */

// bool string

// 其中 (u)int 为不指定位数根据操作系统位数来
// unitptr 为指针类型
// (u)int (u)int8 (u)int16 (u)int32 (u)int64 unitptr

// rune 为字节型32字节 类似于char
// byte rune

// complex为复数 complex64实部和虚部为32位 complex128实部和虚部为64位
// float32 float64 complex64 complex128

func triangle() {
	var a, b int = 3, 4
	var c int
	// go语言必须强制类型转换
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

// 欧拉公式
func euler() {
	c := cmplx.Exp(1i * math.Pi) + 1
	// c := cmplx.Pow(math.E, 1i * math.Pi) + 1
	fmt.Printf("%.3f\n", cmplx.Abs(c))
}

/* ---------------------------------------------------------------------------- */
// 定义常量
func consts() {
	const filename = " abc.txt"
	const a, b = 3, 4
	// 也可以括号扩起来一起定义
	// const (
	// 	filename = "abc.txt"
	// 	a, b = 3, 4
	// )
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums() {
	// iota 具有自增值的含义
	const (
		cpp = iota
		java
		python
		golang
		javascript
	)
	// 也可以使用某种公式自增下去
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

// 主函数
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println("go语言内置类型")
	euler()
	triangle()

	fmt.Println("go语言定义常量")
	consts()
	enums()
}
