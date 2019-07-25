package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// go语言没有while 直接一个for就相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 没有条件就是死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("../branch/abc.txt")
	// forever()
}
