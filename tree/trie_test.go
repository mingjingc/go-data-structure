package tree

import "testing"

func TestTrie(t *testing.T)  {
	trie := NewTrie()
	trie.Add("hello")
	trie.Add("find")
	trie.Add("thank")
	if !trie.Search("hello"){
		t.Error("Trie Search failed")
	}
	trie.Delete("hello")
	trie.Delete("find")
	trie.Delete("thank")
	if trie.Search("hello"){
		t.Error("Trie Search failed")
	}
}