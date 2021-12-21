package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	nodePtr := node
	nodePtr = nil
	fmt.Println(nodePtr)
}

func main() {

}