package main

import "fmt"

// go 没有继承跟多态只有封装

type treeNode struct {
	value       int
	left, right *treeNode
}

// 定义一个函数 xxx 指定外部对象调用，相当于其他语言的 this.xxx()
// node 为接收者 print 为调用方法
func (node treeNode) print() {
	fmt.Print(node.value)
}

// go 语言都是值传递 所以需要将指针传递进去 外部调用方依然是 node.setValue() 外部不需要再转成指针形式
// 要改变类型必须使用指针接收着
// 结构过大也考虑使用指针接收者
// 一致性：如有指针接收者，最好都是指针接收者
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.value = value
}

// 定义工厂函数 可返回局部变量的地址
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	// fmt.Println(root)
	// fmt.Println(root.left)
	// fmt.Println(root.right)
	// fmt.Println(root.left.right)
	// fmt.Println(root.right.left)

	var pNode *treeNode
	pNode.setValue(2)
	fmt.Println(pNode)
	// root.print()
	// root.left.setValue(4)
	fmt.Println(root.left)
}
