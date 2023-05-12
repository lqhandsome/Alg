package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// ListStack 循环栈
type ListStack struct {
	data *Node
	head *Node
	tail *Node
	cap  uint
}
type Node struct {
	Data string
	Next *Node // 下一个节点
}

// InitListStack 初始化以链表为储存的循环栈
func InitListStack(cap uint) *ListStack {

	head := &Node{
		Data: "nil",
		Next: nil,
	}

	tmpHead := head
	for cap > 0 {
		cap--
		tmp := &Node{
			Data: "nil",
			Next: nil,
		}
		tmpHead.Next = tmp
		tmpHead = tmpHead.Next
	}
	tmpHead.Next = head

	return &ListStack{
		data: head,
		head: head,
		tail: head,
		cap:  cap,
	}
}

func (p *ListStack) Range() error {
	if p == nil || p.data == nil {

		log.Errorf("range fail,error: %s\r\n", "p or p.data is nil")

		return fmt.Errorf("range fail,error: %s", "p or p.data is nil")
	}
	if p.IsEmpty() {

		return nil
	}

	tmpTmp := p.tail
	for tmpTmp != p.head {
		fmt.Printf("%s ->", tmpTmp.Data)
		tmpTmp = tmpTmp.Next
	}

	log.Info()

	return nil
}

func (p *ListStack) IsEmpty() bool {
	if p == nil || p.data == nil {
		log.Warnf("r %s\r\n", "p or p.data is nil")

		return true
	}

	if p.head == p.tail {

		return true
	}

	return false
}

func (p *ListStack) IsFull() bool {
	if p == nil || p.data == nil {
		log.Warnf("r %s\r\n", "p or p.data is nil")

		return true
	}

	if p.head.Next == p.tail {

		return true
	}

	return false
}

// Add 往栈插入一个数据
func (p *ListStack) Add(data string) error {
	if p == nil || p.data == nil {
		log.Errorf("r %s\r\n", "p or p.data is nil")

		return fmt.Errorf("r %s\r\n", "p or p.data is nil")
	}

	if p.IsFull() {

		log.Errorf("%s\r\n", "stack is full ")

		return fmt.Errorf("r %s\r\n", "stack is full ")
	}

	p.head.Data = data
	p.head = p.head.Next
	return nil
}

// Pop 往栈插入一个数据
func (p *ListStack) Pop() (data string, e error) {
	if p == nil || p.data == nil {
		log.Errorf("r %s\r\n", "p or p.data is nil")

		return "", fmt.Errorf("r %s\r\n", "p or p.data is nil")
	}

	if p.IsEmpty() {
		log.Warnf("%s", "p is empty")

		return
	}

	data = p.tail.Data
	p.tail = p.tail.Next

	return
}
