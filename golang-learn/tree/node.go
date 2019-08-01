package tree

import "fmt"

// go 没有继承跟多态只有封装

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 定义一个函数 xxx 指定外部对象调用，相当于其他语言的 this.xxx()
// node 为接收者 print 为调用方法
func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

// go 语言都是值传递 所以需要将指针传递进去 外部调用方依然是 node.setValue() 外部不需要再转成指针形式
// 要改变类型必须使用指针接收着
// 结构过大也考虑使用指针接收者
// 一致性：如有指针接收者，最好都是指针接收者
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored")
		return
	}
	node.Value = value
}

// 定义工厂函数 可返回局部变量的地址
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
