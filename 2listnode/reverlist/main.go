package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {

	head := &Node{
		Value: 1,
	}

	one := &Node{Value: 2}
	two := &Node{Value: 3}
	three := &Node{Value: 4}
	head.Next = one
	one.Next = two
	two.Next = three

	head2 := &Node{
		Value: 1,
	}

	one2 := &Node{Value: 2}
	two2 := &Node{Value: 3}
	three2 := &Node{Value: 4}
	head2.Next = one2
	one2.Next = two2
	two2.Next = three2
	//two.Next = one
	//
	Foreach(mergeList(head, head2))
	//head = reverseNode(head)
	//Foreach(head)

	//fmt.Println(checkCircles(head))
}

// 合并两个有序链表
func mergeList(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	newHead := &Node{}
	returnHead := newHead
	for head1 != nil && head2 != nil {
		if head1.Value < head2.Value {
			newHead.Next = head1
			head1 = head1.Next
		} else {
			newHead.Next = head2
			head2 = head2.Next
		}
		newHead = newHead.Next
		fmt.Println(head1, head2)

	}
	if head1 != nil {
		newHead.Next = head1
	}
	if head2 != nil {
		newHead.Next = head2
	}
	return returnHead.Next

}

// 求链表的中间结点
func midNode(head *Node) *Node {

	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 删除链表的倒数第n个节点
func DeleteKNode(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	sentinel := &Node{
		Next: head,
	}
	length := 0
	deleteHead, tmpHead := sentinel, sentinel.Next

	for tmpHead != nil {
		length++
		tmpHead = tmpHead.Next
	}
	for i := 0; i < length-k; i++ {
		deleteHead = deleteHead.Next

	}
	deleteHead.Next = deleteHead.Next.Next
	return sentinel.Next
}

//反转链表
func reverseNode(head *Node) *Node {

	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	var prev *Node
	for tmp != nil {
		next := tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}
	return prev
}

// 检测链表有无环
func checkCircles(head *Node) *Node {

	if head == nil || head.Next == nil {
		fmt.Println("没有环")
		return nil
	}
	ptr, slow, fast := head, head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			fmt.Println("有环", slow)
			return ptr
		}
	}
	fmt.Println("没有环")
	return nil
}

func Foreach(head *Node) {
	for head != nil {
		fmt.Print(head.Value, "-")
		head = head.Next
	}
	fmt.Println()
}
