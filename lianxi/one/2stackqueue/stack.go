package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// ArrStack 循环栈
type ArrStack struct {
	data []string
	head uint // 栈头
	tail uint // 栈尾
	cap  uint
}
type Node struct {
	Data string
	Next *Node // 下一个节点
}

// ListStack 循环栈
type ListStack struct {
	data *Node
	head *Node
	tail *Node
	cap  uint
}

// InitArrStack 初始化数组为储存的顺序栈
func InitArrStack(cap uint) *ArrStack {
	if cap == 0 {
		cap = 10
	}

	return &ArrStack{
		data: make([]string, cap),
		head: 0,
		tail: 0,
		cap:  cap,
	}
}

// Range 打印ArrStack
func (p *ArrStack) Range() {
	if p == nil {
		log.Errorf("%s", "stack is nil")
		return
	}

	if p.head == p.tail {
		fmt.Println("stack is empty")
	}

	tmpHead := p.head
	for tmpHead != p.tail {
		fmt.Print(p.data[tmpHead])
		tmpHead++
		if tmpHead > p.cap {
			tmpHead = 0
		}
	}
	fmt.Println()
}

func (p *ArrStack) Add(data string) error {
	if p == nil || p.data == nil {
		log.Errorf("%s", " arrStack is nil")
		return fmt.Errorf("%s", " arrStack is nil")
	}

	return nil
}

// IsEmpty 栈是否为空
func (p *ArrStack) IsEmpty() bool {
	if p == nil || p.data == nil {
		return true
	}

	if p.tail == p.head {
		return true
	}
	return false
}

// IsFull 栈是否满
func (p *ArrStack) IsFull() bool {
	if p == nil || p.data == nil {

		return true
	}

	if (p.tail+1+p.cap)%p.cap == p.head {

		return true
	} else {

		return false
	}
}
