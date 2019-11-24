package dbl

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

type DBL = DoubleLinkedList

func NewList() *DBL {
	return &DBL{}
}

func (l *DoubleLinkedList) Append(value interface{}) {
	n := Node{
		Value:    value,
		previous: nil,
		next:     nil,
	}

	if l.Head == nil {
		// first node
		l.Head = &n
	} else {
		l.Tail.next = &n
		n.previous = l.Tail
	}
	l.Tail = &n
}

func (l *DoubleLinkedList) Remove(index int) bool {
	currentIndex := 0
	node := l.Head
	for currentIndex < index && node != nil {
		currentIndex += 1
		node = node.next
	}

	// remove the node if it exists
	if currentIndex == index {
		if node.next != nil {
			node.next.previous = node.previous
			node.previous.next = node.next
		}
		// indicate to the gc that this object is not needed anymore by assigning
		// it and it's children to nil
		node.previous = nil
		node.next = nil
		node.Value = nil
		node = nil
		return true
	}

	return false
}

func (l *DoubleLinkedList) RemoveAll() {
	node := l.Head
	for node != nil {
		next := node.next
		node.previous = nil
		node.next = nil
		node.Value = nil
		node = nil
		node = next
	}
	l.Head = nil
	l.Tail = nil
}

func (l *DoubleLinkedList) Get(index int) (*Node, bool) {
	currentIndex := 0
	node := l.Head
	for currentIndex < index && node != nil {
		currentIndex += 1
		node = node.next
	}

	if currentIndex != index {
		return nil, false
	}

	return node, true
}

// Pop will return the last element in the list, and deletes it
func (l *DoubleLinkedList) Pop() *Node {
	// check if there is only one element in the list
	if l.Head == l.Tail {
		n := l.Head
		l.Head = nil
		l.Tail = nil
		return n
	} else if l.Tail != nil {
		n := l.Tail
		if l.Tail.previous == l.Head {
			l.Tail = l.Head
			l.Head.next = nil
		} else {
			l.Tail = l.Tail.previous
			l.Tail.next = nil
		}
		return n
	} else {
		return nil
	}
}

func (l *DoubleLinkedList) Length() int {
	length := 0
	node := l.Head
	for node != nil {
		length += 1
		node = node.next
	}
	return length
}
