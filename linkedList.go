package main

import "fmt"

type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *List) Insert(value any) {
	currentNode := l.Head
	node := &Node{Value: value, Next: nil, Prev: nil}

	if currentNode == nil {
		l.Head = node
		l.Tail = node
		l.Length++
		return
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = node
	currentNode.Next.Prev = currentNode
	l.Tail = node
	l.Length++
}

func (l List) DisplayList() {
	Head := l.Head
	for Head != nil {
		fmt.Println(Head.Value)
		Head = Head.Next
	}
}

func (l *List) DisplayInReverse() {
	Tail := l.Tail
	for Tail != nil {
		fmt.Println(Tail.Value)
		Tail = Tail.Prev
	}
}

func (l *List) RemoveHead() {
	if l.Head.Next == nil {
		l.Head = nil
		l.Length = 0
		return
	}
	l.Head = l.Head.Next
	l.Length--
}
