package main

import (
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

func main() {
	fmt.Println("")
	var root *Node
	root = &Node{val: 0} // {nil, 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	// 마지막에 next가 nil인 node 출력
	fmt.Printf("%d\n", node.val)
}

func AddNode(root *Node, val int) {
	var tail *Node
	tail = root

	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node
}
