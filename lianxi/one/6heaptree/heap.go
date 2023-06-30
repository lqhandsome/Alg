package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

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
		if p.data[n] < p.data[n/2] {
			p.data[n/2], p.data[n] = p.data[n], p.data[n/2]
		}
		n = n / 2
	}
	return nil
}

// deleteData 从堆中删除元素
// 当元素存在时返回nil
func (p *LitterHeap) deleteData(data int) error {

	if p.currLen == 1 {
		return errors.New("not found data")
	}

	if p.data[1] > data {
		return errors.New("not found data")
	}
	var n int
	for i := 1; i < p.currLen; i++ {
		if p.data[i] == data {
			n = i
		}
	}

	if n == 0 {
		return errors.New("not found data")
	}

	p.currLen--
	p.data[n] = p.data[p.currLen]
	p.data[p.currLen] = 0
	log.Info(n, p.currLen, p.data)
	for n*2+1 <= p.currLen-1 {
		tmpMin := n
		if p.data[n] > p.data[n*2+1] {
			tmpMin = n*2 + 1
		}

		if n*2+2 <= p.currLen-1 {
			if p.data[tmpMin] > p.data[n*2+2] {
				tmpMin = n*2 + 2
			}
		}
		// swap
		p.data[n], p.data[tmpMin] = p.data[tmpMin], p.data[n]
		n = tmpMin
	}

	return nil
}
func (p *LitterHeap) isFull() bool {

	if p.currLen == p.len {
		return true
	}
	return false
}
