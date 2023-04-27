package main

import "fmt"

// Node 单向链表
type Node struct {
	Data string
	Next *Node // 下一个节点
}

// InitNode 初始化一个节点
func InitNode(data string) *Node {

	return &Node{
		Data: "head",
		Next: &Node{
			Data: data,
			Next: nil,
		},
	}
}

// AddNode 添加一个节点
func (p *Node) AddNode(node *Node) error {
	if p == nil {
		fmt.Println("node is nil")

		return fmt.Errorf("%s", "node is nil")
	}

	head := p
	for head.Next != nil {
		head = head.Next
	}

	head.Next = node

	return nil
}

// DeleteNode 删除指定节点
func (p *Node) DeleteNode(data string) error {
	if p == nil {
		fmt.Println("node is nil")

		return fmt.Errorf("%s", "node is nil")
	}

	head := p.Next
	pre := p
	for head != nil {
		if head.Data == data {
			pre.Next = head.Next
			return nil
		}
		head = head.Next
		pre = pre.Next

	}

	fmt.Printf("data: %s not found\r\n", data)

	return fmt.Errorf("data: %s not found\r\n", data)
}

// Range 遍历链表
func (p *Node) Range() {
	if p == nil {

		return
	}

	head := p
	for head != nil {
		fmt.Print(head.Data, "->")
		head = head.Next
	}

	fmt.Println()

	return
}

// reversalList 反转单链表
func reversalList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *Node
	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = pre
		pre = tmp
		tmp = next
	}
	return pre
}
