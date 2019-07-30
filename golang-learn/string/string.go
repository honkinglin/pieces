package main

import (
	"fmt"
	"unicode/utf8"
)

// rune 相当于go的char
// 使用range 遍历pos, rune对
// 使用utf8.RuneCountInString 获得字符数量
// 使用len获得字节长度
// 使用[]byte获得字节

func main() {
	s := "Yes我爱Github!" // UTF-8编码 可变长 中文占三个字节
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// rune是四个字节 字符串底层是3个字节转成byte会造成误解 将字符串转成 rune 类型就无需关心底层是几个字节 
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
