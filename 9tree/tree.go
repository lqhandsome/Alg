package main

import (
	"fmt"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//	a := &TreeNode{
//		Data: "A",
//	}
//	b := &TreeNode{
//		Data: "B",
//	}
//	c := &TreeNode{
//		Data: "C",
//	}
//	d := &TreeNode{
//		Data: "D",
//	}
//	e := &TreeNode{
//		Data: "E",
//	}
//	f := &TreeNode{
//		Data: "F",
//	}
//	g := &TreeNode{
//		Data: "G",
//	}
//	a.Left = b
//	a.Right = c
//	b.Left = d
//	b.Right = e
//	c.Left = f
//	c.Right = g
//
//	//a.PreOrderTraversal()
//	fmt.Print("先序遍历")
//	PreOrderTraversal(a)
//	fmt.Println()
//	fmt.Print("中序遍历")
//	InOrderTraversal(a)
//	fmt.Println()
//	fmt.Print("后序遍历")
//	PostOrderTraversal(a)
//	fmt.Println()
//	fmt.Print("非递归前序遍历")
//	beforeRangeTree(a)
//	fmt.Println()
//	fmt.Print("非递归后续遍历")
//	afterRangeTree(a)
//	fmt.Println()
//	fmt.Print("非递归层续遍历")
//	levelRangeTree(a)
//}

//前序非递归遍历二叉树
func beforeRangeTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	//用一个切片模拟栈
	stackArr := []*TreeNode{}
	for tree != nil || len(stackArr) != 0 {

		if tree == nil {
			// 如果节点为空，且堆栈中有数据，就弹出一个并遍历其右节点
			if len(stackArr) != 0 {
				//fmt.Print(stackArr[len(stackArr)-1].Data,"-->") // 中序遍历
				tree = stackArr[len(stackArr)-1].Right
				stackArr = stackArr[:len(stackArr)-1]
			}
		} else {
			fmt.Print(tree.Data, "-->") // 前序遍历
			stackArr = append(stackArr, tree)
			tree = tree.Left
		}
	}
}
func afterRangeTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 上一个访问的节点
	var pre *TreeNode
	//用一个切片模拟栈
	stackArr := []*TreeNode{}
	for tree != nil || len(stackArr) != 0 {

		// 先将所有左子树压入
		for tree != nil {
			stackArr = append(stackArr, tree)
			tree = tree.Left
		}

		// 弹出一个，这时候是第二次访问
		if len(stackArr) != 0 {
			tmp := stackArr[len(stackArr)-1]
			tree = stackArr[len(stackArr)-1]

			stackArr = stackArr[:len(stackArr)-1]
			//如果当前节点没有右子树，或者当前节点的右子树刚刚访问过，就要打印
			if tree.Right == nil || tree.Right == pre {
				fmt.Print(tree.Data, "-->") // 后序遍历
				// 因为此节点没有是返回上来的，已经访问过，所以需要置为nil
				tree = nil
				pre = tmp
			} else {
				stackArr = append(stackArr, tree)
				tree = tree.Right
			}
		}

	}
}

func levelRangeTree(tree *TreeNode) {
	if tree == nil {
		return
	}
	//用一个切片模拟队列
	stackArr := []*TreeNode{}
	for tree != nil || len(stackArr) != 0 {
		fmt.Print(tree.Data, "-->") // 后序遍历
		if tree.Left != nil {
			stackArr = append(stackArr, tree.Left)
		}
		if tree.Right != nil {
			stackArr = append(stackArr, tree.Right)
		}
		if len(stackArr) != 0 {
			tree = stackArr[0]
			stackArr = stackArr[1:]
		} else {
			break
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
