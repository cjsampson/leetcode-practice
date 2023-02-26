package main

import (
	"fmt"
	"leetcode/word_break22/structures"
)

var x = []int{0,0,1,-1}
var y = []int{1,-1,0,0}

func main() {
	board := [][]byte{{111, 97, 97, 110}, {101, 116, 97, 101}, {105, 104, 107, 114}, {105, 102, 108, 118}}
	words := []string{"oath", "pea", "eat", "rain"}

	set := findWords(board, words)
	fmt.Println("---------------")
	fmt.Printf("set:%v\n", set)
	fmt.Println("---------------")
}

func findWords(board [][]byte, words []string) structures.Set {
	result := structures.Set{}

	trie := structures.InitTrie()
	root := trie.Root
	for _, word := range words {
		trie.Insert(word)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			currentChar := board[i][j]
			if root.Children[currentChar - 'a'] != nil {
				dfs(board, i, j, root, currentChar, result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i int, j int, trieNode *structures.TrieNode, currentChar byte, result structures.Set) {
	// if current board piece is outside boundary, has already been visited or doesnt exist within the trie - return
	if outside(board, i, j) || board[i][j] == '#' || trieNode.Children[currentChar - 'a'] == nil {
		return
	}

	temp := board[i][j]
	board[i][j] = '#'

	trieNode = trieNode.Children[temp - 'a']
	if trieNode != nil && trieNode.IsWord {
		result.Add(trieNode.Word)
	}

	for k := 0; k < len(x); k++ {
		rowOffset, columnOffset := i + x[k], j + y[k]

		if outside(board, rowOffset, columnOffset) {
			continue
		}

		dfs(board, rowOffset, columnOffset, trieNode, board[rowOffset][columnOffset], result)
	}

	board[i][j] = temp
	return
}

func outside(board [][]byte, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j > len(board) {
		return true
	}

	return false
}
