package list

import (
	"unsafe"
)

// XORList
// head                  													   tail
//
//	↓                     														↓
//
// A|add(B) -> B|addr(A)^addr(C) -> C|addr(B)^addr(D) -> D|addr(C)^addr(E) -> E|addr(D)
//
// 处于节点B，想求节点C的地址, addr(A)^(link(B))==addr(A)^(addr(A)^addr(C))=0^addr(C)=addr(C)
// 处于节点B，想求节点A的地址, addr(C)^(link(B))==addr(C)^(addr(A)^addr(C))=0^addr(A)=addr(A)
type XORList struct {
	head *XORListNode
	tail *XORListNode
}

type XORListNode struct {
	val int
	np  *XORListNode
}

func NewXORList() *XORList {
	return &XORList{}
}

func (list *XORList) Insert(val int) {
	newNode := newNode(val)
	if list.head == nil {
		list.head = newNode
	} else {
		list.tail.np = xor(list.tail.np, newNode)
		newNode.np = xor(list.tail, nil)
	}
	list.tail = newNode
}

func (list *XORList) Search(val int) bool {
	curr := list.head
	var prev *XORListNode
	var next *XORListNode

	for curr != nil {
		if curr.val == val {
			return true
		}
		next = xor(prev, curr.np)
		prev = curr
		curr = next
	}

	return false
}

func (list *XORList) Remove(val int) bool {
	curr := list.head
	var prev *XORListNode
	var next *XORListNode

	for curr != nil {
		if curr.val == val {
			if curr == list.head && curr == list.tail {
				list.head = nil
				list.tail = nil
			} else if curr == list.head {
				list.head = curr.np
				list.head.np = xor(curr, list.head.np)
			} else if curr == list.tail {
				list.tail = prev
				if prev != nil { // always true
					list.tail.np = xor(curr, prev.np)
				}

			} else {
				next = xor(prev, curr.np)
				if prev != nil { // always true
					prev.np = xor(xor(prev.np, curr), next)
				}
				next.np = xor(prev, xor(curr, next.np))
			}
			return true
		}
		next = xor(prev, curr.np)
		prev = curr
		curr = next
	}

	return false
}

func newNode(data int) *XORListNode {
	return &XORListNode{val: data, np: nil}
}

func xor(prev, next *XORListNode) *XORListNode {
	return (*XORListNode)(unsafe.Pointer(uintptr(unsafe.Pointer(prev)) ^ uintptr(unsafe.Pointer(next))))
}
