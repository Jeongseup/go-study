package main

import (
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (ll *LinkedList) AddNode(val int) {
	// 처음 만들어져서 루트랑 테일이 없는 경우
	if ll.root == nil {
		ll.root = &Node{val: val}
		ll.tail = ll.root
		return
	}

	ll.tail.next = &Node{val: val}
	// 기존 테일을 저장
	prev := ll.tail
	ll.tail = ll.tail.next
	// 새로운 테일에 기존 prev를 저장
	ll.tail.prev = prev
}

func (ll *LinkedList) RemoveNode(node *Node) {
	// 지우려고 하는 노드가 루트 노드인 경우
	if node == ll.root {
		ll.root = ll.root.next
		ll.root.prev = nil
		// 맨 앞 노드는 연결 해제
		node.next = nil
		return
	}

	// 지우려고 하는 노드가 루트 노드가 아닌 경우
	// 지우련는 노드 탐색
	prev := node.prev

	// 지우려고 하는 노드가 테일인지 확인
	if node == ll.tail {
		prev.next = nil
		ll.tail.prev = nil
		ll.tail = prev
	} else { // 테일이 아닌 중간 노드
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
}

func (ll *LinkedList) PrintNodes() {
	node := ll.root
	// node의 next가 nil일때 까지 루프
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (ll *LinkedList) PrintReverse() {
	node := ll.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

	list.PrintReverse()
}
