package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// tree 二叉查找树
type tree struct {
	data  int
	left  *tree
	right *tree
}

// InitTree 初始化一个二叉查找树
func InitTree(data int) *tree {
	return &tree{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (p *tree) Add(data int) error {

	t := &tree{
		data:  data,
		left:  nil,
		right: nil,
	}
	tmp := p
	for tmp != nil {
		if tmp.data > data {
			if tmp.left == nil {
				tmp.left = t
				return nil
			}
			tmp = tmp.left
		} else {
			if tmp.right == nil {
				tmp.right = t
				return nil
			}
			tmp = tmp.right
		}
	}

	return nil
}

func (p *tree) Search(data int) bool {
	if p == nil {

		return false
	}

	tmp := p
	for tmp != nil {
		if tmp.data == data {
			return true
		} else if tmp.data > data {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}

	return false
}

// LevelRange 层序遍历打印
func (p *tree) LevelRange() {
	tmp := p
	var queue []*tree

	queue = append(queue, tmp)
	for len(queue) > 0 {

		tmpQueue := queue[0]
		queue = queue[1:]
		fmt.Println(tmpQueue.data)
		if tmpQueue.left != nil {
			queue = append(queue, tmpQueue.left)
		}

		if tmpQueue.right != nil {
			queue = append(queue, tmpQueue.right)
		}
	}
}

// PreRange 先序遍历二叉树
func beforeRange(root *tree) {
	if root == nil {
		return
	}
	log.Info(root.data)
	beforeRange(root.left)
	beforeRange(root.right)
}

// afterRange 二叉树后序遍历
func afterRange(root *tree) {
	if root == nil {
		return
	}
	afterRange(root.left)
	afterRange(root.right)
	log.Info(root.data)
}

// midRange 二叉树中序遍历
func midRange(root *tree) {
	if root == nil {
		return
	}

	midRange(root.left)
	log.Info(root.data)
	midRange(root.right)
}

func (p *tree) Delete(data string) error {

	if p == nil {

		return fmt.Errorf("%s", "tree not init")
	}

	// 删除的节点没有子节点

	//删除的节点只有一个节点
	//删除的节点有两个节点
	return nil
}
