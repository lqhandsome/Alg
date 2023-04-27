package main

import "testing"

func TestDoubleList(t *testing.T) {
	doubleList := InitDoubleNode("a")
	doubleList.AddNode(&DoubleNode{
		Data: "b",
		Pre:  nil,
		Next: nil,
	})
	doubleList.Range()
}
