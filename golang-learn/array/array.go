package main

import "fmt"

// go 语言中数组也是作为值拷贝传进去的并不是传入地址引用 所以函数内修改数组并不会影响外部的数据
// 如果要更改外部的值只能修改成指针形式，外部也使用 &arr 传入引用地址
func printArray(arr *[5]int) {
	arr[0] = 100
	// 通过range 关键字可以获取到数组的下标 第二个参数可以获取到具体的值
	for i, v := range arr {
		// fmt.Println(arr3[i])
		fmt.Println(i, v)
	}
}

func main()  {
	var arr1 [5]int
	// 使用 := 定义数组需要指定初始值
	arr2 := [3]int{1, 3, 5}
	// 也可以不指定长度
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[i])
	// }

	fmt.Println("printArray arr1")
	printArray(&arr1)

	//  [3]int 和 [5]int 不是一个类型
	// printArray(arr2)

	fmt.Println("printArray arr3")
	printArray(&arr3)

	fmt.Println("arr1 & arr3")
	fmt.Println(arr1, arr3)
}
