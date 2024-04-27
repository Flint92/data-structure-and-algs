package tree

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isWord   bool
	children map[rune]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	parent := t.root
	if parent == nil {
		return
	}

	for _, ch := range word {
		chNode := parent.children[ch]
		if chNode == nil {
			parent.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		parent = parent.children[ch]
	}

	parent.isWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.find(word)
	return current != nil && current.isWord == true
}

func (t *Trie) StartWith(prefix string) bool {
	return t.find(prefix) != nil
}

func (t *Trie) Candidates(word string) []rune {
	current := t.find(word)
	if current == nil {
		return []rune{}
	}

	var candidates []rune
	for candidate := range current.children {
		candidates = append(candidates, candidate)
	}

	return candidates
}

func (t *Trie) find(word string) *TrieNode {
	parent := t.root
	if parent == nil {
		return nil
	}

	for _, ch := range word {
		chNode := parent.children[ch]
		if chNode == nil {
			return nil
		}
		parent = chNode
	}

	return parent
}
