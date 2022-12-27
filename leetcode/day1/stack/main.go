package main

import "fmt"

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{InStack: []int{}, OutStack: []int{}}
}

// 压入
func (this *CQueue) AppendTail(value int) {
	this.InStack = append(this.InStack, value)
}

// 弹出 先检查输出栈是否为空，如果为空就把输入栈的数据全部压入
func (this *CQueue) DeleteHead() int {
	if len(this.InStack) == 0 && len(this.OutStack) == 0 {
		return -1
	}

	// 输出栈有数据
	if len(this.OutStack) > 0 {
		value := this.OutStack[len(this.OutStack)-1]
		this.OutStack = this.OutStack[:len(this.OutStack)-1]
		return value
	}
	//输出栈无数据输入栈有数据
	//将输入栈数据搬移至输出栈
	for i := len(this.InStack); i > 0; i-- {
		this.OutStack = append(this.OutStack, this.InStack[i-1])
	}
	this.InStack = []int{}
	return this.DeleteHead()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	cqueue := Constructor()
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(1)
	cqueue.AppendTail(2)
	cqueue.AppendTail(3)
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	fmt.Println(cqueue.DeleteHead())
	cqueue.AppendTail(3)
	fmt.Println(cqueue.DeleteHead())
}
