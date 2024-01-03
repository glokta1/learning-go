package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

// if the original data structure is a pointer and you want to modify it in-place
// it requires an invalid pointer-to-pointer receiver
// solution: wrap the pointer in a struct and define methods on that struct
type LinkedList[T comparable] struct {
	head *Node[T]
}

// inserts a Node at the beginning of the list
func (ll *LinkedList[T]) Insert(x T) {
	// create node with value x
	new_node := &Node[T]{nil, x}

	// fix links
	new_node.next = ll.head
	ll.head = new_node
}

// Deletes a Node with value x
func (ll *LinkedList[T]) Delete(x T) {
	// traverse to node with value x
	var prev, curr *Node[T]
	for prev, curr = nil, ll.head; curr != nil && curr.val != x; prev, curr = curr, curr.next {
	}

	// x is at first node
	if prev == nil {
		ll.head = curr.next
		return
	}

	prev.next = curr.next
}

func (ll LinkedList[T]) String() string {
	var out string
	for p := ll.head; p != nil; p = p.next {
		out += fmt.Sprintf("%v->", p.val)
	}

	return out
}

func main() {
	// head is a pointer to a node
	var ll LinkedList[int]
	ll.Insert(9)
	ll.Insert(12)
	ll.Insert(3)
	fmt.Println(ll)
	ll.Delete(3)

	fmt.Println(ll)
}
