package main

import "fmt"

/*
	int preIndex = 0, posIndex = 0;
	TreeNode* constructFromPrePost(vector<int>& pre, vector<int>& post) {
		TreeNode* root = new TreeNode(pre[preIndex++]);
		if (root->val != post[posIndex])
		root->left = constructFromPrePost(pre, post);
		if (root->val != post[posIndex])
		root->right = constructFromPrePost(pre, post);
		posIndex++;
		return root;
	}
*/

func main() {
	preorder := []int{100, 20, 10, 30, 200, 150, 300}
	postorder := []int{10, 30, 20, 150, 300, 200, 100}

	root := buildTree(preorder, postorder)
	inorder(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder  := []int{100, 20, 10, 30, 200, 150, 300}
// postorder := []int{10, 30, 20, 150, 300, 200, 100}
func buildTree(preorder, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}

	var index int
	for i := 0; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			index = i + 1
			break
		}
	}

	fmt.Printf("preorder first half: %v\n", preorder[1:index+1])
	fmt.Printf("preorder second half: %v\n", preorder[index+1:])

	fmt.Printf("postorder first half: %v\n", postorder[:index])
	fmt.Printf("postorder second half: %v\n", postorder[:index])

	root.Left = buildTree(preorder[1:index+1], postorder[0:index])
	root.Right = buildTree(preorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func inorder(node *TreeNode) {
	if node == nil {
		fmt.Print("|")
		return
	}

	inorder(node.Left)
	fmt.Print(node.Val)
	inorder(node.Right)
}
