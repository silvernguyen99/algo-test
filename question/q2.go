package test

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func (dl *DoubleLinkedList) PushBack(data int) *DoubleLinkedList {
	newNode := &Node{
		Data: data,
	}

	if dl.Head == nil {
		dl.Head = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Prev = dl.Tail
	}

	dl.Tail = newNode
	return dl
}

var errEmpty = errors.New("error - DoubleLinkedList is empty")

func (dl *DoubleLinkedList) PushFront(data int) *DoubleLinkedList {
	newNode := &Node{
		Data: data,
	}

	if dl.Tail == nil {
		dl.Tail = newNode
	} else {
		newNode.Next = dl.Head
		dl.Head.Prev = newNode
	}
	dl.Head = newNode
	return dl
}

func (dl *DoubleLinkedList) PopBack() (int, error) {
	if dl.Tail == nil {
		return 0, errEmpty
	}
	v := dl.Tail.Data
	dl.Tail = dl.Tail.Prev
	if dl.Tail == nil {
		dl.Head = nil
	} else {
		if dl.Tail.Next != nil {
			dl.Tail.Next = nil
		}
	}

	return v, nil
}

func (dl *DoubleLinkedList) PopFront() (int, error) {
	if dl.Head == nil {
		return 0, errEmpty
	}
	v := dl.Head.Data
	dl.Head = dl.Head.Next

	if dl.Head == nil {
		dl.Tail = nil
	} else {
		if dl.Head.Prev != nil {
			dl.Head.Prev = nil
		}
	}

	return v, nil
}

func (dl *DoubleLinkedList) String() string {
	if dl.Head == nil {
		fmt.Print("list is empty")
		return ""
	}

	str := ""
	cur := dl.Head
	for cur != nil {
		str += fmt.Sprint(cur.Data) + " "
		cur = cur.Next
	}

	return str
}
