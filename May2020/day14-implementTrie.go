// Implement a trie with insert, search, and startsWith methods.
// Example:
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true
// Note:
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.

package main

type TrieNode struct {
	links []*TrieNode
	isEnd bool
}

func (n *TrieNode) ContainsKey(ch byte) bool {
	return n.links[ch-byte('a')] != nil
}

func (n *TrieNode) Get(ch byte) *TrieNode {
	return n.links[ch-byte('a')]
}

func (n *TrieNode) Put(ch byte, node *TrieNode) {
	n.links[ch-byte('a')] = node
}

func (n *TrieNode) SetEnd() {
	n.isEnd = true
}

func (n *TrieNode) IsEnd() bool {
	return n.isEnd
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{
		root: &TrieNode{
			links: make([]*TrieNode, 26),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node *TrieNode = this.root
	w := []byte(word)
	for i := 0; i < len(w); i++ {
		chr := w[i]
		if !node.ContainsKey(chr) {
			node.Put(chr, &TrieNode{
				links: make([]*TrieNode, 26),
			})
		}
		node = node.Get(chr)
	}
	node.SetEnd()
}

func (this *Trie) searchPrefix(word string) *TrieNode {
	var node *TrieNode = this.root
	w := []byte(word)
	for i := 0; i < len(w); i++ {
		chr := w[i]
		if node.ContainsKey(chr) {
			node = node.Get(chr)
		} else {
			return nil
		}
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var node *TrieNode = this.searchPrefix(word)

	return node != nil && node.IsEnd()
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var node *TrieNode = this.searchPrefix(prefix)

	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
