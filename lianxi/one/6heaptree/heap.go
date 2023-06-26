package main

import "errors"

// LitterHeap
// 小顶堆,index=0不存储数据
type LitterHeap struct {
	len     int   // 数据长度
	currLen int   // 当前元素长度
	data    []int // 储存数据
}

func New(len int) *LitterHeap {
	heap := &LitterHeap{
		len:     len + 1,
		currLen: 1,
		data:    make([]int, len+1),
	}
	heap.data[0] = -1

	return heap
}

func (p *LitterHeap) add(data int) error {
	if p.isFull() {
		return errors.New("LitterHeap is full")
	}

	n := p.currLen
	p.data[p.currLen] = data
	p.currLen++
	// 堆化
	for n/2 > 0 {
		if p.data[n] > p.data[n/2] {
			p.data[n/2], p.data[n] = p.data[n], p.data[n/2]
		}
		n = n / 2
	}
	return nil
}

func (p *LitterHeap) isFull() bool {

	if p.currLen == p.len {
		return true
	}
	return false
}
