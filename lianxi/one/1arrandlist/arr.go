package main

import "fmt"

type Array struct {
	Data []int
	Len  uint
}

// InitArray 初始化一个Arrpy
func InitArray(len uint) *Array {
	return &Array{
		Data: make([]int, len, len),
		Len:  len,
	}
}

func (p *Array) Set(index uint, value int) bool {
	if p == nil {

		return false
	}

	if index > p.Len-1 {
		fmt.Println("超出数组")

		return false
	}

	p.Data[index] = value
	return true
}

// Range 遍历Array
func (p *Array) Range() {
	if p == nil {

		return
	}

	for _, v := range p.Data {
		fmt.Println(v)
	}
}
