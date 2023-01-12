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
	fmt.Print("先序遍历")
	PreOrderTraversal(a)
	fmt.Println()
	fmt.Print("中序遍历")
	InOrderTraversal(a)
	fmt.Println()
	fmt.Print("后序遍历")
	PostOrderTraversal(a)
	fmt.Println()
	fmt.Print("非递归前序遍历")
	beforeRangeTree(a)
	fmt.Println()
}

//前序非递归遍历二叉树
func beforeRangeTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	//用一个切片模拟栈
	stackArr := []*TreeNode{}
	for tree != nil || len(stackArr) != 0{

		if tree == nil {
			// 如果节点为空，且堆栈中有数据，就弹出一个并遍历其右节点
			if len(stackArr) != 0{
				//fmt.Print(stackArr[len(stackArr)-1].Data,"-->") // 中序遍历
				tree =stackArr[len(stackArr)-1].Right
				stackArr = stackArr[:len(stackArr)-1]
			}
		}else {
			fmt.Print(tree.Data,"-->") // 前序遍历
			stackArr = append(stackArr,tree)
			tree = tree.Left
		}
	}
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
