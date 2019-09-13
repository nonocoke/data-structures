package linklist

import "fmt"

// ListNode is an element of a linked list.
type ListNode struct {
	value interface{}
	next *ListNode
}

// NewListNode return a initialization of a Element.
func NewListNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}

// GetNext returns the next list ListNode or nil.
func (e *ListNode) GetNext() *ListNode {
	return e.next
}

// GetValue returns the value of current ListNode.
func (e *ListNode) GetValue() interface{} {
	return e.value
}

// SingleLinkedList is an single linked list
type SingleLinkedList struct {
	head *ListNode
	size uint
}

// NewSingleLinkedList return a initialization of a SingleLinkedList.
// To use the SingleLinkedList from the l.head.next.
func NewSingleLinkedList() * SingleLinkedList {
	return &SingleLinkedList{NewListNode(0), 0}
}

// InsertAfter
func (l *SingleLinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)

	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext

	l.size ++
	return true
}

// InsertBefore
func (l *SingleLinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || l.head == p {
		return false
	}

	// 查找前驱结点 pre
	pre, cur := l.head, l.head.next
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	// 当前链表中不存在 p 结点
	if nil == cur {
		return false
	}

	newNode := NewListNode(v)

	pre.next = newNode
	newNode.next = cur

	l.size ++
	return true
}

// InsertToHead
func (l *SingleLinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail
func (l *SingleLinkedList) InsertToTail(v interface{}) bool {
	// 查找最后一个结点
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}

// FindByIndex
func (l *SingleLinkedList) FindByIndex(idx uint) *ListNode {
	if idx > l.size {
		return nil
	}

	var i uint = 0
	cur := l.head.next
	for ; i < idx; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode
func (l *SingleLinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	// 查找前驱结点 pre
	pre, cur := l.head, l.head.next
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}

	pre.next = p.next
	p = nil  // 置空

	l.size --
	return true
}

func (l *SingleLinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
