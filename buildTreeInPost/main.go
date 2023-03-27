package main

import "fmt"

func main() {
	//inorder := []int{6, 9, 3, 15, 20, 7}
	//postorder := []int{6, 9, 15, 7, 20, 3}

	inorder := []int{10, 20, 30, 100, 150, 200, 300}
	postorder := []int{10, 30, 20, 150, 300, 200, 100}

	treeNode := buildTree(inorder, postorder)
	treeNode.TraverseInOrder()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) TraverseInOrder() {
	inOrderTraverse(t)
}

func inOrderTraverse(node *TreeNode) {
	if node == nil {
		fmt.Print("|")
		return
	}

	inOrderTraverse(node.Left)
	fmt.Printf("Node: %v\n", node.Val)
	inOrderTraverse(node.Right)
}

// in-order   1,2,3,4,5,6,7,8,9,10,11
// pre-order  8,4,2,1,3,6,5,7,10,9,11
// post-order 1,3,2,5,7,6,4,9,11,10,8

// inorder   = []int{6, 9, 3, 15, 20, 7}
// postorder = []int{6, 9, 15, 7, 20, 3}

// inorder   = []int{10, 20, 30, 100, 150, 200, 300}
// postorder = []int{10, 30, 20, 150, 300, 200, 100}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	length := len(postorder)
	root := &TreeNode{
		Val: postorder[length-1],
	}
	var i int
	for ; inorder[i] != root.Val; i++ {
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:length-1])

	return root
}
