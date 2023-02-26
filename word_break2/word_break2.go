package main

import (
	"fmt"
)

var x = []int{0, 0, 1, -1}
var y = []int{1, -1, 0, 0}

func main() {
	board := [][]byte{{111, 97, 97, 110}, {101, 116, 97, 101}, {105, 104, 107, 114}, {105, 102, 108, 118}}
	words := []string{"oath", "pea", "eat", "rain"}

	res := findWords(board, words)
	fmt.Printf("res:%v\n", res)
}

func findWords(board [][]byte, words []string) set {
	trie := initTrie()
	root := trie.root
	result := set{}

	// insert possible words into trie
	for _, w := range words {
		trie.insert(w)
	}

	// loop through each board piece to find possible words on board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// current char on board
			c := board[i][j]
			// if trie root contains current board char - continue searching
			if root.children[c-'a'] != nil {
				dfs(board, i, j, root, c, result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i int, j int, tn *trieNode, c byte, result set) {
	if outside(board, i, j) || board[i][j] == '#' || tn.children[c-'a'] == nil {
		return
	}

	// hold current piece temporarily
	temp := board[i][j]
	// set visited board piece to arbitrary indicator
	board[i][j] = '#'
	// check trie node to see if the current board piece exists
	tn = tn.children[temp-'a']

	// base case - have we matched a word in the trie
	if tn != nil && tn.isWordEnd {
		result.add(tn.word)
	}

	// move horizontally and vertically on board to check for words in trie
	for k := 0; k < len(x); k++ {
		// check (1,0), (-1,0), (0,1), (0,-1)
		rowOffset, columnOffset := i+x[k], j+y[k]

		// are the offsets out of bounds?
		if outside(board, rowOffset, columnOffset) {
			continue
		}

		// if offsets not out of bounds - continue checking board for words
		dfs(board, rowOffset, columnOffset, tn, board[rowOffset][columnOffset], result)
	}

	// set the board piece back from to original char
	board[i][j] = temp
	return
}

func outside(board [][]byte, i, j int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board) {
		return true
	}
	return false
}

const (
	ALPHABET_SIZE = 26
)

func initTrie() trie {
	return trie{&trieNode{}}
}

type trie struct {
	root *trieNode
}

type trieNode struct {
	children  [ALPHABET_SIZE]*trieNode
	isWordEnd bool
	word      string
}

func (t *trie) insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
	current.word = word
}

type set map[string]struct{}

func (s set) add(word string) {
	s[word] = struct{}{}
}

func (s set) has(word string) bool {
	_, ok := s[word]
	return ok
}
