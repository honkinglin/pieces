package main

import "fmt"

// go 语言一般不使用数组 一般都使用切片 slice
// slice 本身没有数据，是对底层 array 的一个 view 相当于数组地址的映射
// slice 底层有三个变量 ptr len cap， ptr为数组切片的第一个位置 len为切片区间的长度 cap代表愿数组的最后一个位置
// 只要切片区间在原数组的[ptr, cap]内都可以取得到值 但只能向后扩展，不可以向前扩展
// s[i] 不可以超越 len(s)，向后拓展不可以超越底层数组 cap(s)

func updateSlice(s []int) {
	s[0] = 100	
}

func printSlice() {
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		fmt.Printf("len(s)=%d, cap(s)=%d \n", len(s), cap(s))
		// 由于是值传递的关系，必须接收append的返回值 例如： s = append(s, val)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)
}

// 创建切片的另一种形式
func makeSlice() {
	fmt.Println("Creating slice")
	s1 := make([]int, 16) // 第二个参数为 len 的值
	s2 := make([]int, 10, 32) // 第三个参数为 cap 的值
	s3 := []int{2, 4, 6, 8}
	fmt.Printf("len(s1)=%d, cap(s1)=%d \n", len(s1), cap(s1))
	fmt.Printf("len(s2)=%d, cap(s2)=%d \n", len(s2), cap(s2))

	fmt.Println("Copying slice")
	copy(s2, s3)
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // 删除操作比较蛋疼 需要将删除的前几位数提取出来再将删除后几位数提取出来拼成一个新数组
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]

	fmt.Println(tail)
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:]", s1)
	s2 := arr[:]
	fmt.Println("arr[:]", s2)

	// fmt.Println("After updateSlice(s1)")
	// updateSlice(s1)
	// fmt.Println(s1)
	// fmt.Println("After updateSlice(s2)")
	// updateSlice(s2)
	// fmt.Println(s2)
	// fmt.Println(arr)

	s3 := arr[1:2]
	// slice 可以向后扩展，不可以向前扩展
	s4 := s3[1:2]
	fmt.Println(s3, s4)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))

	s5 := append(s3, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	// 添加元素时如果超越cap，系统会重新分配更大的底层数组 原本数组如果没有引用会被垃圾回收掉
	fmt.Println("s5, s5, s7=", s5, s6, s7)
	fmt.Println("arr =", arr)

	// printSlice()
	makeSlice()
}
