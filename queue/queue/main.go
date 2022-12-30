package main

import "fmt"

// 十个容量的队列
type Queue struct {
	tail uint
	head uint
	arr  [3]int
}

func main() {
	queue := Queue{}
	fmt.Println(cap(queue.arr))
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}

func (this *Queue) Push(i int) {
	if this.IsFull() {
		fmt.Println("队列已满")
		return
	}
	this.arr[this.tail] = i
	this.tail = (this.tail + 1) % uint(cap(this.arr))
}

func (this *Queue) Pop() int {
	if this.head == this.tail {
		fmt.Println("队列为空")
		return -1
	}
	value := this.arr[this.head]
	this.head = (this.head + 1) % uint(cap(this.arr))
	return value
}

// 判断队列是否为空
func (this *Queue) IsEmpty() bool {
	return this.head == this.tail
}

//判断队列是否满
func (this *Queue) IsFull() bool {
	if (this.tail+1)%uint(cap(this.arr)) == this.head {
		return true
	}
	return false
}
