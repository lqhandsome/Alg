package main

import "fmt"

func main() {
	arr := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
	InitHeap(arr)
	//InitHeadTwo(arr) // [0 20 16 19 13 4 1 7 5 8]

	fmt.Println(arr)
}

// 第一种堆化
// 思路 从第一个元素开启对话依次调用插入
func InitHeap(arr []int) {
	for i := 2; i < len(arr); i++ {
		InsertData(arr, i)
	}
}

// 插入元素
func InsertData(arr []int, index int) {
	for i := index; i > 1; i = i / 2 {
		if arr[i] > arr[i/2] {
			arr[i], arr[i/2] = arr[i/2], arr[i]
		}
	}
}

// 第二种堆化思路
// 从第一个非叶子结点开始堆化（index=0 不存储数据） n/2 n 是数组长度,然后从这个节点开始向下比较子节点，比较完后 再比较下一个非叶子节点
func InitHeadTwo(arr []int) {
	for i := (len(arr) - 1) / 2; i > 0; i-- {
		//arr := []int{0, 7, 5, 19, 8, 4, 1, 20, 13, 16}
		fmt.Println(i)
		handle(arr, i)
		fmt.Println(arr)
	}
}

func handle(arr []int, index int) {
	for i := index; i*2 < len(arr) && i < len(arr); {
		tmp := i
		if arr[tmp] < arr[i*2] {
			tmp = i * 2
		}
		if arr[tmp] < arr[i*2+1] {
			tmp = i*2 + 1
		}
		if tmp != i {
			arr[tmp], arr[i] = arr[i], arr[tmp]
			i = tmp
			continue
		}
		break
	}
}

// 从堆中删除堆顶元素
