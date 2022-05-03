package main

import (
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0} // {nil, 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)

	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)

}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node

	// new tail
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		// 루트를 삭제하고 보니까 아무것도 남지 않는 경우
		// root와 tail 모두 nil 처리
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	// 삭제하려는 노드가 테일인 경우?
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		// 루트 노드의 다음 다음 (2번째)를 루트의 다음으로 집어 넣는다.
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	// 마지막에 next가 nil인 node 출력
	fmt.Printf("%d\n", node.val)
}
