// Given a 2D board and a list of words from the dictionary, find all words in the board.
// Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
// Example:
// Input:
// board = [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]
// words = ["oath","pea","eat","rain"]
// Output: ["eat","oath"]
// Note:
// All inputs are consist of lowercase letters a-z.
// The values of words are distinct.
//    Hide Hint #1
// You would need to optimize your backtracking to pass the larger test. Could you stop backtracking earlier?
//    Hide Hint #2
// If the current candidate does not exist in all words' prefix, you could stop backtracking immediately. What kind of data structure could answer such query efficiently? Does a hash table work? Why or why not? How about a Trie? If you would like to learn how to implement a basic trie, please work on this problem: Implement Trie (Prefix Tree) first.

package main

type TrieNode struct {
	child map[byte]*TrieNode
	word  string
}

var (
	_board  [][]byte
	_result []string
)

func findWords(board [][]byte, words []string) []string {
	_board, _result = board, nil
	root := &TrieNode{
		child: make(map[byte]*TrieNode),
		word:  "",
	}
	for _, w := range words {
		node := root
		for _, b := range w {
			if _, ok := node.child[byte(b)]; ok {
				node = node.child[byte(b)]
			} else {
				newNode := &TrieNode{
					child: make(map[byte]*TrieNode),
					word:  "",
				}
				node.child[byte(b)] = newNode
				node = newNode
			}
		}
		node.word = w
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if _, ok := root.child[board[r][c]]; ok {
				backtracking(r, c, root)
			}
		}
	}

	return _result
}

func backtracking(r, c int, p *TrieNode) {
	l := _board[r][c]
	if currNode, ok := p.child[l]; ok {
		if currNode.word != "" {
			_result = append(_result, currNode.word)
			currNode.word = ""
		}

		_board[r][c] = 0
		rowOffset, colOffset := []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
		for i := 0; i < 4; i++ {
			newR, newC := r+rowOffset[i], c+colOffset[i]
			if newR < 0 || newR >= len(_board) || newC < 0 || newC >= len(_board[0]) {
				continue
			}
			if _, ok := currNode.child[_board[newR][newC]]; ok {
				backtracking(newR, newC, currNode)
			}
		}

		_board[r][c] = l
		if len(currNode.child) == 0 {
			delete(p.child, l)
		}
	}
}
