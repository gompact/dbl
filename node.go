package dbl

type Node struct {
	Value    interface{}
	previous *Node
	next     *Node
}
