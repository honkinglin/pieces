package main

import (
	"fmt"

	"../../tree"
)

// 封装
// 名字一般使用CamelCase
// 首字母大写：public
// 首字母小写：private

// 包
// 每个目录一个包
// main包 包含可执行入口
// 为结构定义的方法必须放在同一个包内，可以是不同文件

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right.Left)

	var pNode *tree.TreeNode
	pNode.SetValue(2)
	fmt.Println(pNode)
	root.Print()
	root.Left.SetValue(4)
	fmt.Println(root.Left)
}
