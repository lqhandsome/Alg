package main

import (
	"fmt"
	"runtime"
	"time"
)

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {

	// 长子串
	var maxLen int = firstLen
	// 短子串
	var minLen int = secondLen
	if firstLen < secondLen {
		maxLen = secondLen
		minLen = firstLen
	}
	// 长子串index
	var maxIndex, i int
	for ; i <= len(nums)-maxLen; i++ {
		if getSum(nums, i, maxLen) > getSum(nums, maxIndex, maxLen) {
			maxIndex = i
		}
	}

	var j int
	if maxIndex+maxLen > j+minLen && j+minLen > maxIndex {
		j = maxIndex + maxLen
	}

	//	短子串index
	var minIndex int = j
	for j <= len(nums)-minLen {

		if maxIndex+maxLen > j+minLen && j+minLen > maxIndex {
			j = maxIndex + maxLen
			j++
			continue
		}
		if getSum(nums, j, minLen) > getSum(nums, minIndex, minLen) {
			minIndex = j
		}
		j++
	}
	fmt.Printf("max: %v,min: %v\r\n", maxIndex, minIndex)
	return getSum(nums, maxIndex, maxLen) + getSum(nums, minIndex, minLen)
}

func getSum(num []int, index int, len int) int {
	var sum int
	for len > 0 {
		sum += num[index]
		index++
		len--
	}
	return sum
}

func main() {
	x := 0
	threads := runtime.GOMAXPROCS(0)
	fmt.Println(threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x = x + 1
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("x =", x)
	return
	a := &TreeNode{
		Data: "A",
	}
	b := &TreeNode{
		Data: "B",
	}
	c := &TreeNode{
		Data: "C",
	}
	d := &TreeNode{
		Data: "D",
	}
	e := &TreeNode{
		Data: "E",
	}
	f := &TreeNode{
		Data: "F",
	}
	g := &TreeNode{
		Data: "G",
	}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g

	preRange(a)
	fmt.Println()
	inRange(a)
	fmt.Println()
	afterRange(a)
	fmt.Println()
	preRangeFor(a)
	midRangeFor(a)
}

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func preRange(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%s --->", root.Data)
	preRange(root.Left)
	preRange(root.Right)
}

func inRange(root *TreeNode) {
	if root == nil {
		return
	}

	inRange(root.Left)
	fmt.Printf("%s --->", root.Data)
	inRange(root.Right)
}

func afterRange(root *TreeNode) {
	if root == nil {
		return
	}

	afterRange(root.Left)
	afterRange(root.Right)
	fmt.Printf("%s --->", root.Data)
}

// midRangeFor 中序非递归遍历
func midRangeFor(root *TreeNode) {
	if root == nil {
		return
	}

	// 用一个切片
	stack := make([]*TreeNode, 0, 0)
	// 栈或者root还有数据
	for root != nil || len(stack) != 0 {
		//左节点为空则开始取出一个节点
		if root == nil {
			fmt.Println(stack[len(stack)-1].Data)
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}

		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}
}

// preRangeFor 前序非递归遍历
func preRangeFor(root *TreeNode) {
	if root == nil {
		return
	}

	// 用一个切片
	stack := make([]*TreeNode, 0, 0)
	// 栈或者root还有数据
	for root != nil || len(stack) != 0 {
		//左节点为空则开始取出一个节点
		if root == nil {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}

		if root != nil {
			fmt.Println(root.Data)
			stack = append(stack, root)
			root = root.Left
		}
	}
}
