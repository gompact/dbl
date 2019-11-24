package dbl

import "testing"

func TestDoubleLinkedList_Prepend(t *testing.T) {
	dbl := buildList()

	dbl.Prepend(-1)

	if dbl.Head.Value != -1 {
		t.Errorf("expected list's head value to be -1, found %v", dbl.Head.Value)
	}

	if dbl.Head.next.Value != 0 {
		t.Errorf("expected list's second item value to be 0, found %v", dbl.Head.next.Value)
	}
}

func TestDoubleLinkedList_Append(t *testing.T) {
	dbl := buildList()

	if l := dbl.Length(); l != 3 {
		t.Errorf("expected list length to be 3, found %d", l)
	}

	if dbl.Head.Value != 0 {
		t.Errorf("expected list's head value to be 0, found %v", dbl.Head.Value)
	}

	if dbl.Tail.Value != 2 {
		t.Errorf("expected list's head value to be 2, found %v", dbl.Head.Value)
	}
}

func buildList() *DBL {
	dbl := NewList()
	dbl.Append(0)
	dbl.Append(1)
	dbl.Append(2)
	return dbl
}

func TestDoubleLinkedList_Get(t *testing.T) {
	dbl := buildList()

	if v, _ := dbl.Get(0); v.Value != 0 {
		t.Errorf("expected Get(0) to return a node of value 0, found %d", v.Value)
	}

	if v, _ := dbl.Get(1); v.Value != 1 {
		t.Errorf("expected Get(1) to return a node of value 1, found %d", v.Value)
	}

	if v, _ := dbl.Get(2); v.Value != 2 {
		t.Errorf("expected Get(2) to return a node of value 2, found %d", v.Value)
	}

	if _, ok := dbl.Get(4); ok {
		t.Errorf("expected Get(4) to return not ok, but ok(true) was returned")
	}
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	dbl := buildList()

	dbl.Remove(1)

	if l := dbl.Length(); l != 2 {
		t.Errorf("expected list length to be 2, found %d", l)
	}

	if v, _ := dbl.Get(0); v.Value != 0 {
		t.Errorf("expected Get(0) to return a node of value 0, found %d", v.Value)
	}

	if v, _ := dbl.Get(1); v.Value != 2 {
		t.Errorf("expected Get(1) to return a node of value 2, found %d", v.Value)
	}

	if dbl.Head.next != dbl.Tail {
		t.Errorf("expected the list head.next to be equal to the tail")
	}

	if dbl.Tail.previous != dbl.Head {
		t.Errorf("expected the list tail.previous to be equal to the head")
	}
}

func TestDoubleLinkedList_RemoveAll(t *testing.T) {
	dbl := buildList()

	dbl.RemoveAll()

	if l := dbl.Length(); l != 0 {
		t.Errorf("expected RemoveAll to make the list empty, but found %d items", l)
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	dbl := buildList()

	n := dbl.Pop()
	if l := dbl.Length(); l != 2 {
		t.Errorf("expected Pop to reduce the list length down to 2, but found %d", l)
	}

	if n.Value != 2 {
		t.Errorf("expected the popped element value to be 2, found %d", n.Value)
	}

	n = dbl.Pop()
	if l := dbl.Length(); l != 1 {
		t.Errorf("expected Pop to reduce the list length down to 1, but found %d", l)
	}

	if n.Value != 1 {
		t.Errorf("expected the popped element value to be 1, found %d", n.Value)
	}

	n = dbl.Pop()
	if l := dbl.Length(); l != 0 {
		t.Errorf("expected Pop to reduce the list length down to 0, but found %d", l)
	}

	if n.Value != 0 {
		t.Errorf("expected the popped element value to be 0, found %d", n.Value)
	}

	if dbl.Head != nil {
		t.Errorf("expected the list head to be nil")
	}

	if dbl.Tail != nil {
		t.Errorf("expected the list tail to be nil")
	}
}