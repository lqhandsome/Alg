package main

func main() {

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
