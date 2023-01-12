package main

import (
	"fmt"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{
		Data: "A",
	}
	b := &TreeNode{
		Data: "B",
	}
	c := &TreeNode{
		Data: "C",
	}
	d := &TreeNode{
		Data: "D",
	}
	e := &TreeNode{
		Data: "E",
	}
	f := &TreeNode{
		Data: "F",
	}
	g := &TreeNode{
		Data: "G",
	}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g

	//a.PreOrderTraversal()
	PreOrderTraversal(a)
	fmt.Println()
	InOrderTraversal(a)
	fmt.Println()
	PostOrderTraversal(a)
}

// 添加一个节点
func (p *TreeNode) AddChildNode(parent *TreeNode, child *TreeNode, left bool) {

}

func (p *TreeNode) PreOrderTraversal() {
	if p != nil {
		fmt.Println(p.Data)
		p.Left.PreOrderTraversal()
		p.Right.PreOrderTraversal()
	}
}

//先序遍历
func PreOrderTraversal(tree *TreeNode) {
	if tree != nil {
		fmt.Print(tree.Data, "-->")
		PreOrderTraversal(tree.Left)
		PreOrderTraversal(tree.Right)
	}
}

//中序遍历
func InOrderTraversal(tree *TreeNode) {
	if tree != nil {
		InOrderTraversal(tree.Left)
		fmt.Print(tree.Data, "-->")
		InOrderTraversal(tree.Right)
	}
}

//后序遍历
func PostOrderTraversal(tree *TreeNode) {
	if tree != nil {
		PostOrderTraversal(tree.Left)
		PostOrderTraversal(tree.Right)
		fmt.Print(tree.Data, "-->")
	}
}
