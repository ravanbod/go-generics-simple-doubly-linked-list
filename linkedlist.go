package main

import "errors"

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T] // Storing tail node to insert without traversing
}

type Node[T comparable] struct {
	prev *Node[T]
	next *Node[T]
	val  T
}

func (l *List[T]) insert(val ...T) { // It accepts arbitrary number of T
	if l.head == nil {
		l.head = &Node[T]{val: val[0]}
		l.tail = l.head
		val = val[1:]
	}
	for _, val := range val {
		l.tail.next = &Node[T]{val: val}
		l.tail.next.prev = l.tail
		l.tail = l.tail.next
	}

}

func (l *List[T]) returnAsSlice() []T {
	var r []T
	for current := l.head; current != nil; current = current.next {
		r = append(r, current.val)
	}
	return r
}

func (l *List[T]) returnReverseAsSlice() []T {
	var r []T
	for current := l.tail; current != nil; current = current.prev {
		r = append(r, current.val)
	}
	return r
} 

func (l *List[T]) findByValue(val T) (int, *Node[T], error) { // It will return error if it doesn't find node with value = val
	for index, current := 0, l.head; current != nil; index, current = index+1, current.next {
		if current.val == val {
			return index, current, nil
		}
	}
	return 0, nil, errors.New("Node Not found")
}

func (l *List[T]) deleteByValue(val T) error { // It will return error if it doesn't find node with value = val
	_, node, err := l.findByValue(val)
	if err != nil {
		return err
	}
	if l.head == l.tail { // It will that if list has 1 node.
		l.head = nil
		l.tail = nil
	} else if node == l.head { // It deletes head node
		l.head = l.head.next
		l.head.prev.next = nil
		l.head.prev = nil
	} else if node == l.tail { // It deletes tail node
		l.tail = l.tail.prev
		l.tail.next.prev = nil
		l.tail.next = nil
	} else { // other nodes ...
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return nil
}
