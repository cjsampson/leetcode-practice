package structures

import "fmt"

func InitTrie() *trie {
	return &trie{&TrieNode{}}
}

type trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children [26]*TrieNode
	IsWord   bool
	Word     string
}

func (t *trie) Insert(word string) {
	current := t.Root
	for _, letter := range word {
		index := letter - 'a'
		if current.Children[index] == nil {
			current.Children[index] = &TrieNode{}
		}
		current = current.Children[index]
	}
	current.IsWord = true
	current.Word = word
}

func (t *trie) Has(word string) bool {
	current := t.Root
	for _, letter := range word {
		// find index into trie of current letter
		index := letter - 'a'
		fmt.Printf("string(index): %v\n",string(index))
		// if current index not found - word doesnt exist
		if current.Children[index] == nil {
			return false
		}
		// set current to the current index
		current = current.Children[index]
	}

	return word == current.Word
}
