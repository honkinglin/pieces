package main

import (
	"io/ioutil"
	"fmt"
)

// if 的条件里可以赋值
// if 的条件里赋值的变量作用域就在这个 if 语句里
func ifOerator() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// switch 会自动 breack，除非使用 fallthrough
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		// panic 相当于报错
		panic("unsupported operator:" + op)
	}
	return result
}

// 也可以不指定具体 switch 的变量 直接 case 一个判断表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	fmt.Println(g)
	return g
}

func main() {
	ifOerator()
	// eval()
	grade(100)
}
