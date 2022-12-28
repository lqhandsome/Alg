package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Stack    []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    []int{},
		MinStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	if x < this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, x)
	} else {
		this.MinStack = append(this.MinStack, this.MinStack[len(this.MinStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) <= 0 {
		return -1
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) Min() int {
	return this.MinStack[len(this.MinStack)-1]
}

func main() {

	minStack := Constructor()
	minStack.Push(-10)
	minStack.Push(2)
	minStack.Push(3)
	minStack.Push(-4)
	fmt.Println(minStack.Min())
}
