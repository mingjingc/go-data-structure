package tree

import "strings"

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children [256]*trieNode
	isValue bool
}

func NewTrie() *Trie {
	return &Trie{
		root:newTrieNode(),
	}
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: [256]*trieNode{},
		isValue:  false,
	}
}

func (t *Trie) Add(value string)  {
	value = strings.Trim(value," ")
	if len(value)==0{
		return
	}
	node:=t.root
	for i:=0;i<len(value);i++{
		b := value[i]
		if node.children[b]==nil {
			node.children[b] = newTrieNode()
		}
		node = node.children[b]
	}
	node.isValue = true
}
func (t *Trie) Search(value string) bool {
	value = strings.Trim(value," ")
	if len(value) == 0{
		return false
	}
	node:=t.root
	for i:=0;i<len(value);i++{
		b := value[i]
		if node.children[b]==nil{
			return false
		}
		node = node.children[b]
	}
	return true
}
func (t *Trie) Delete(value string)  {
	value = strings.Trim(value," ")
	if len(value)==0{
		return
	}
	node := t.root
	var pre *trieNode
	for i:=0;i<len(value);i++{
		b := value[i]
		if node.children[b]==nil{
			return
		}
		pre = node
		node = node.children[b]
	}
	node.isValue = false
	if t.nodeHasEmptyChildren(node) {
		pre.children[value[len(value)-1]]=nil
	}
}

func (t *Trie) nodeHasEmptyChildren(node *trieNode) bool {
	for _,n:=range node.children {
		if n!=nil{
			return false
		}
	}
	return true
}