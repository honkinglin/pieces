package main

import (
	"fmt"

	"./mock"
	"./real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var m Retriever
	m = mock.Retriever{"this is a mock"}
	fmt.Println(download(m))

	var r Retriever
	r = real.Retriever{}
	fmt.Println(download(r))
}
