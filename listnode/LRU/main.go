package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LRU struct {
	List   *Node
	Len    uint
	MaxLen uint
}

func main() {
	lru := GetLRU()
	lru.Add(1)
	lru.Add(2)
	lru.Add(3)
	lru.Add(4)

	lru.Add(5)
	lru.Add(6)
	lru.Add(5)
	lru.Foreach()
	lru.Get(3)
	lru.Foreach()
}

func GetLRU() *LRU {
	return &LRU{MaxLen: 5}
}

// 新增一个缓存
func (p *LRU) Add(value int) {
	// 数据已经存在
	head := p.List
	for head != nil {
		if head.Value == value {
			tmpHead := p.List
			for tmpHead.Next.Value != value {
				tmpHead = tmpHead.Next
			}
			findTmp := tmpHead.Next
			tmpHead.Next = tmpHead.Next.Next
			findTmp.Next = p.List
			p.List = findTmp
			return
		}
		head = head.Next
	}

	// 队列满
	if p.Len == p.MaxLen {
		node := &Node{Value: value}
		node.Next = p.List
		p.List = node
		var i uint = 0
		head := p.List
		for ; i <= p.MaxLen-2; i++ {
			head = head.Next
		}
		head.Next = nil
		return
	}
	node := &Node{Value: value}
	node.Next = p.List
	p.List = node
	p.Len++
}

func (p *LRU) Get(value int) int {
	head := p.List
	for head != nil {
		if head.Value == value {
			tmpHead := p.List
			for tmpHead.Next.Value != value {
				tmpHead = tmpHead.Next
			}
			findTmp := tmpHead.Next
			tmpHead.Next = tmpHead.Next.Next
			findTmp.Next = p.List
			p.List = findTmp
			return (value)
		}
		head = head.Next
	}
	return -1
}
func (p *LRU) Foreach() {
	head := p.List
	for head != nil {
		fmt.Print(head.Value, "-")
		head = head.Next
	}
	fmt.Println()
}
