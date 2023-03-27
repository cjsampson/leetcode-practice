package main

type Tree struct {
	root *Node
}

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}
