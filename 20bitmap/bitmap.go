package main

import "fmt"

type BitMap struct {
	Data  []byte
	NBits int //整数最大
}

// go 中 byte（8 bit）可以表示八个数字
func InitBitMap(nbits int) *BitMap {

	return &BitMap{
		Data:  make([]byte, (nbits>>3)+1),
		NBits: nbits,
	}
}
func (p *BitMap) Set(k int) {
	if k > p.NBits {
		return
	}

	p.Data[k/8] |= 1 << (k % 8)
}

func (p *BitMap) Get(k int) bool {
	if k > p.NBits {
		return false
	}
	return p.Data[k/8]&(1<<(k%8)) != 0
}
func main() {

	fmt.Println(5 >> 2)
	return
	//var a byte
	bitMap := InitBitMap(100)
	bitMap.Set(10)
	fmt.Println(bitMap.Get(10))
	fmt.Println(bitMap.Get(1))
}
