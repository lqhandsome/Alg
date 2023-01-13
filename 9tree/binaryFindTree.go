package main

import (
	"errors"
	"fmt"
)

type BinaryFindTree struct {
	Data      int
	LeftTree  *BinaryFindTree
	RightTree *BinaryFindTree
}

func newBinaryFindTree(data int) *BinaryFindTree {
	return &BinaryFindTree{
		Data: data,
	}
}

func main() {
	root := newBinaryFindTree(5)
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 10,
	}))
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 3,
	}))
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 5,
	}))
	root.levelRangeTree()
	fmt.Println()
	root.inRangeTree()
	fmt.Println()

	fmt.Println(root.FindNode(1))
	fmt.Println(root.FindNode(3))
	fmt.Println(root.FindNode(5))
	fmt.Println(root.FindNode(9))
	fmt.Println(root.FindNode(10))
}
func (root *BinaryFindTree) inRangeTree() {
	if root == nil {
		return
	}
	root.LeftTree.inRangeTree()
	fmt.Print(root.Data, "-->")
	root.RightTree.inRangeTree()
}
func (root *BinaryFindTree) levelRangeTree() {
	if root == nil {
		return
	}
	//用一个切片模拟队列
	stackArr := []*BinaryFindTree{}
	for root != nil || len(stackArr) != 0 {
		fmt.Print(root.Data, "-->") // 后序遍历
		if root.LeftTree != nil {
			stackArr = append(stackArr, root.LeftTree)
		}
		if root.RightTree != nil {
			stackArr = append(stackArr, root.RightTree)
		}
		if len(stackArr) != 0 {
			root = stackArr[0]
			stackArr = stackArr[1:]
		} else {
			break
		}
	}
}
func (root *BinaryFindTree) AddNode(node *BinaryFindTree) error {
	if root == nil {
		fmt.Println("root is nil")
		return errors.New("root is nil")
	}
	if node == nil {
		fmt.Println("node is nil")
		return errors.New("node is nil")
	}
	if node.Data >= root.Data {
		if root.RightTree == nil {
			root.RightTree = node
		} else {
			root.RightTree.AddNode(node)
		}
	} else {
		if root.LeftTree == nil {
			root.LeftTree = node
		} else {
			root.LeftTree.AddNode(node)
		}
	}
	return nil
}

// 二叉树的查找
func (root *BinaryFindTree) FindNode(data int) (*BinaryFindTree, error) {

	if root == nil {
		return nil, errors.New("not found")
	}
	if data == root.Data {
		return root, nil
	}
	if data < root.Data {
		return root.LeftTree.FindNode(data)
	} else {
		return root.RightTree.FindNode(data)
	}
}

//二叉查找树的删除
func (root *BinaryFindTree) DeleteNode(data int) error {

	if root == nil {
		fmt.Println("root is nil")
		return errors.New("root is nil")
	}
	// 当前节点
	p := root
	//当前节点的父节点
	var pp *BinaryFindTree

	// 查找节点
	for p != nil && p.Data != data {
		pp = p
		if p.Data > data {
			p = p.LeftTree
		} else {
			p = p.RightTree
		}
	}
	// 退出条件是 p == nil
	if p == nil {
		fmt.Println("node not found")
		return errors.New("node not found")
	}

	// 情况1 没有子节点
	if p.RightTree == nil && p.LeftTree == nil {
		// 如果删除的是头结点
		if pp == nil {
			root.Data = 0
			return nil
		}
		if pp.LeftTree == p {
			pp.LeftTree = p.LeftTree
		} else {
			pp.RightTree = p.LeftTree
		}
	}

	//情况2 有且只有一个节点
	if (p.RightTree != nil && p.LeftTree == nil) || p.RightTree == nil && p.LeftTree != nil {
		if p.LeftTree != nil {
			if pp.LeftTree == p {
				pp.LeftTree = p.LeftTree
			} else {
				pp.RightTree = p.LeftTree
			}
		} else {
			if pp.LeftTree == p {
				pp.LeftTree = p.RightTree
			} else {
				pp.RightTree = p.RightTree
			}
		}
	}

	//情况3 有两个节点
	return nil
}
