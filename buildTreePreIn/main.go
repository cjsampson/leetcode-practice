package main

import "fmt"

func main() {
	inorder := []int{10, 20, 30, 100, 150, 200, 300}
	preorder := []int{100, 20, 10, 30, 200, 150, 300}
	root := buildTreePreIn(preorder, inorder)
	traverseinorder(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverseinorder(node *TreeNode) {
	if node == nil {
		fmt.Print("|")
		return
	}
	traverseinorder(node.Left)
	fmt.Print(node.Val)
	traverseinorder(node.Right)
}

// in-order   1,2,3,4,5,6,7,8,9,10,11
// pre-order  8,4,2,1,3,6,5,7,10,9,11
// post-order 1,3,2,5,7,6,4,9,11,10,8

// Inorder Traversal:   10 20 30 100 150 200 300
// Preorder Traversal:  100 20 10 30 200 150 300

// inorder traversal:   10 20 30
// preorder traversal:  20 10 30

// inorder traversal:   10
// preorder traversal:  10
func buildTreePreIn(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	index := getIndex(inorder, preorder[0])

	root.Left = buildTreePreIn(preorder[1:index+1], inorder[:index])
	root.Right = buildTreePreIn(preorder[index+1:], inorder[index+1:])

	return root
}

func getIndex(arr []int, target int) int {
	var index int
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			index = i
		}
	}

	return index
}
