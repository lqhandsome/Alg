package main

import "testing"

func TestListInit(t *testing.T) {

	node := InitNode("a")
	node.AddNode(&Node{
		Data: "b",
		Next: nil,
	})
	node.Range()
	node.DeleteNode("a")
	node.DeleteNode("c")
	node.DeleteNode("b")
	node.Range()
}

func TestRever(t *testing.T) {
	node := InitNode("a")
	node.AddNode(&Node{
		Data: "b",
		Next: nil,
	})
	node.AddNode(&Node{
		Data: "c",
		Next: nil,
	})
	node.Range()

	node.Next = reversalList(node.Next)
	node.Range()
}

// 链表中间节点
func TestMidNode(t *testing.T) {
	node := InitNode("a")
	node.AddNode(&Node{
		Data: "b",
		Next: nil,
	})
	node.AddNode(&Node{
		Data: "c",
		Next: nil,
	})
	midNode := midList(node.Next)
	midNode.Range()
}
