package bitmap

import (
	"errors"
)

// BitMap 位图
type BitMap struct {
	maxSize int // 一共保存多少个数据
	data    []byte
}

// NewBitMap 初始化位图
func NewBitMap(cap int) *BitMap {
	b := &BitMap{
		maxSize: cap,
		data:    make([]byte, cap/8+1), //byte=8bits
	}

	return b
}

func (b *BitMap) Add(data int) error {
	if data > b.maxSize {
		return errors.New("over flow")
	}

	index := data / 8
	y := data % 8
	b.data[index] = b.data[index] | (1 << y)

	return nil
}

func (b *BitMap) Get(data int) bool {
	if data > b.maxSize {
		return false
	}

	index := data / 8
	y := data % 8

	return (b.data[index] & (1 << y)) != 0
}
