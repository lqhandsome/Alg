package main

import "fmt"

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
