package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// ArrStack 循环栈 栈尾不存储数据
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
		data: []string{"null", "null", "null", "null", "null"},
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
		return
	}

	tmpTail := p.tail
	for tmpTail != p.head {
		fmt.Print(p.data[tmpTail], "-->")
		tmpTail++
		if tmpTail > p.cap-1 {
			tmpTail = 0
		}
	}
	fmt.Println()
}

func (p *ArrStack) Add(data string) error {
	if p == nil || p.data == nil {
		log.Errorf("%s", " arrStack is nil")
		return fmt.Errorf("%s", " arrStack is nil")
	}

	if p.IsFull() {
		log.Errorf("%s", "stack is full")

		return fmt.Errorf("%s", "stack is full")
	}

	p.data[p.head] = data
	p.head = (p.head + 1) % (p.cap)
	return nil
}

// Pop 弹出一个元素
func (p *ArrStack) Pop() (data string, e error) {

	if p == nil || p.data == nil {
		log.Errorf("%s", "p or data is nil")

		return "null", fmt.Errorf("%s", "p or data is nil")
	}

	if p.IsEmpty() {

		log.Errorf("%s", "stack is empty")

		return "null", fmt.Errorf("%s", "stack is empty")
	}

	data = p.data[p.tail]
	p.tail = (p.tail + 1) % p.cap
	return
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

	if (p.head+1)%(p.cap) == p.tail {
		fmt.Printf("head: %v,tail: %v,cap: %v\r\n", p.head, p.tail, p.cap)
		return true
	} else {

		return false
	}
}
