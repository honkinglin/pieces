package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// 函数返回值类型写在最后面
// 可返回多个值
// 函数是一等公民可作为参数传入
// 没有默认参数、可选参数等复杂语法 只有一个可变参数列表

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// go语言中定义的变量一定要使用否则编译不通过，若不想使用可以使用 _ 定义
		q, _ := div(a, b)
		return q, nil
	default:
		// panic 相当于报错
		// panic("unsupported operator:" + op)
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

// 函数可以有多个返回值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d, %d)\n", opName, a ,b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表 相当于传入一个数组
func sum(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	fmt.Println(eval(3, 4, "/"))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	a, b := div(12, 4)
	fmt.Println(a, b)

	fmt.Println(apply(pow, 3, 4))
	// 匿名函数不需要指定名称
	fmt.Println(apply(func (a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1,2,3,4,5))
}
