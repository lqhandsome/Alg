package main

import "fmt"

// DoubleNode 双向链表
type DoubleNode struct {
	Data string
	Pre  *DoubleNode
	Next *DoubleNode
}

func InitDoubleNode(data string) *DoubleNode {
	head := &DoubleNode{
		Data: "head",
		Pre:  nil,
		Next: nil,
	}

	node := &DoubleNode{
		Data: data,
		Pre:  head,
		Next: nil,
	}

	head.Next = node

	return head
}

func (p *DoubleNode) AddNode(node *DoubleNode) error {
	if p == nil {
		fmt.Println("doubleNode is nil")

		return fmt.Errorf("%v", "doubleNode is nil")
	}

	head := p

	for head.Next != nil {

		head = head.Next
	}

	head.Next = node
	node.Pre = head

	return nil
}

func (p *DoubleNode) Range() {
	if p == nil {

		return
	}

	head := p
	for head != nil {
		fmt.Printf("%v-->", head.Data)
		head = head.Next
	}
}
