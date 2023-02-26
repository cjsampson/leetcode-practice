package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
	//	Output: [[2,2,2],[2,2,0],[2,0,1]]
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	color := 2
	expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	got := floodFill(image, sr, sc, color)
	if reflect.DeepEqual(got, expected) {
		fmt.Println("correct floodfill")
		return
	}

	fmt.Println("incorrect floodfill")
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var thing [][]int

	fmt.Println("image: ", image)
	return thing
}

func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}
	next := []int{}

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			visited[node] = true
			fn(node)
			for _, n := range bfs_frontier(node, nodes, visited) {
				next = append(next, n)
			}
		}
	}
}

func bfs_frontier(node int, nodes map[int][]int, visited map[int]bool) []int {
	next := []int{}
	iter := func(n int) bool {
		_, ok := visited[n]
		return !ok
	}
	for _, n := range nodes[node] {
		if iter(n) {
			next = append(next, n)
		}
	}
	return next
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type NodeWithLevel struct {
	node  *Node
	level int
}

func bfsTree(root *Node) *Node {
	if root == nil { //to handle edge case
		return root
	}
	q := []NodeWithLevel{
		{
			node:  root,
			level: 0,
		},
	}
	visited := []*Node{}
	for len(q) > 0 {
		vertex := q[0]
		node, level := vertex.node, vertex.level
		visited = append(visited, node)
		q = q[1:] //dequeue first node in queue(fifo)

		if node.Left != nil { //have both left and right since it's a perfect binary tree
			leftNode := NodeWithLevel{
				node:  node.Left,
				level: level + 1,
			}
			q = append(q, leftNode) //append left-child to back of queue(fifo)

			rightNode := NodeWithLevel{
				node:  node.Right,
				level: level + 1,
			}
			q = append(q, rightNode) //append right-child to back of queue(fifo)
		}
	}
	return root
}
