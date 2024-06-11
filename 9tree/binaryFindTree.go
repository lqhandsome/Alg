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

func mainn() {
	root := newBinaryFindTree(7)
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 10,
	}))
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 3,
	}))
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 5,
	}))
	fmt.Println(root.AddNode(&BinaryFindTree{
		Data: 1,
	}))
	root.levelRangeTree()
	fmt.Println()
	root.inRangeTree()
	fmt.Println()
	newRoot, e := DeleteNode(root, 7)
	if e != nil {
		panic(e)
	}
	newRoot.levelRangeTree()
	fmt.Println()
	newRoot.inRangeTree()
	fmt.Println()

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

// 二叉查找树的删除
func DeleteNode(root *BinaryFindTree, data int) (*BinaryFindTree, error) {

	if root == nil {
		fmt.Println("root is nil")
		return nil, errors.New("root is nil")
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
		return nil, errors.New("node not found")
	}

	// 情况1 没有子节点
	if p.RightTree == nil && p.LeftTree == nil {
		// 如果删除的是头结点
		if pp == nil {
			root.Data = 0
			fmt.Println("删除的是头结点")
			return nil, nil
		}
		if pp.LeftTree == p {
			pp.LeftTree = p.LeftTree
		} else {
			pp.RightTree = p.LeftTree
		}
		return root, nil
	}

	//情况2 有且只有一个节点
	if (p.RightTree != nil && p.LeftTree == nil) || p.RightTree == nil && p.LeftTree != nil {
		if pp == nil {
			root.Data = 0
			fmt.Println("删除的是头结点")
			if root.LeftTree == nil {
				return root.RightTree, nil
			} else {
				return root.LeftTree, nil
			}
		}
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
		return root, nil
	}
	//情况3 有两个节点
	if p.RightTree != nil && p.LeftTree != nil {
		if pp == nil {
			fmt.Println("删除的是头结点")
			minTree := p.RightTree
			minTreePP := p
			for minTree.LeftTree != nil {
				minTreePP = minTree
				minTree = minTree.LeftTree
			}
			minTreePP.RightTree = minTree.RightTree
			minTree.LeftTree = p.LeftTree
			minTree.RightTree = p.RightTree
			p.Data = minTree.Data
			return root, nil
		}
		minTree := p.RightTree
		minTreePP := p
		for minTree.LeftTree != nil {
			minTreePP = minTree
			minTree = minTree.LeftTree
		}
		minTreePP.RightTree = minTree.RightTree
		minTree.LeftTree = p.LeftTree
		minTree.RightTree = p.RightTree
		if pp.LeftTree == p {
			pp.LeftTree = minTree
		} else {
			pp.RightTree = minTree
		}

		return root, nil
	}
	return nil, fmt.Errorf("unknown error")
}
