package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (h *LinkedList) AddEnd(data int) {
	newNode := &Node{data: data}
	if h.head == nil {
		h.head = newNode
	} else {
		cur := h.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newNode
	}
}

func (h *LinkedList) Print() {
	cur := h.head
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}

func main() {
	list := &LinkedList{}
	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)
	list.AddEnd(4)
	list.Print()
}
