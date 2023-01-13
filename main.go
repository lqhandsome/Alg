package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type demo struct {
	Val int
}

func (d *demo) change() {
	v := reflect.ValueOf(d).Elem()
	v.SetPointer(unsafe.Pointer(new(demo)))
	println(d)
}

func (d *demo) myVal() {
	fmt.Printf("my val: %#v\n", d)
}

func (d demo) change2() {
	d = demo{} // 对方法接收器的分配不会传播到其他调用
	d.myVal()
}

func (d *demo) change3() {
	d.Val = 3
	d.myVal()
}

func main() {
	//return
	d := &demo{
		Val: 10,
	}
	println(d)
	d.change()
	return
	d.myVal()
	d.change()
	d.myVal()
	d.Val = 2
	d.myVal()
	d.change2()
	d.myVal()
	d.change3()
	d.myVal()
}
