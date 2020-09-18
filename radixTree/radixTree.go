package radixtree

type node struct {
	path     string
	children []*node
	isWord   bool
}

func NewRadixTree() (n *node) {
	n = new(node)
	n.path = ""
	return
}

func (n *node) Add(word string) {
	if len(word)==0{
		return
	}

	currentNode := n
	prefixNode := currentNode.findPrefixNode(word[0])
	if prefixNode != nil{
		currentNode = prefixNode
	}

	walk:
	for {

		lcp := lengthCommonPrefix(word,currentNode.path)

		// currentNode is root
		if lcp==0 {
			child := &node{
				path:     word,
				isWord:   true,
			}
			currentNode.children = append(n.children, child)
			return
		}

		if lcp<len(currentNode.path){
			child := &node{
				path:     currentNode.path[lcp:],
				children: currentNode.children,
				isWord:   currentNode.isWord,
			}
			currentNode.path = currentNode.path[:lcp]
			currentNode.children = []*node{child}
		}

		if lcp<len(word) {
			word = word[lcp:]
			prefixNode = currentNode.findPrefixNode(word[0])
			if prefixNode!=nil{
				currentNode = prefixNode
				continue walk
			}
			child:= &node{
				path:     word,
				isWord:   true,
			}
			currentNode.children = append(currentNode.children,child)
			return
		} else  {
			currentNode.isWord = true
			return
		}
	}
}

func (n *node) Delete(word string) {
	if len(word)==0{
		return
	}

	currentNode := n
	walk:
	for {
		for i:=0;i<len(currentNode.children);i++{
			prefixEndI := prefixEndPosition(word,currentNode.children[i].path)
			if prefixEndI>0 {
				if prefixEndI == len(word) {
					if !currentNode.children[i].isWord {
						return
					}

					if len(currentNode.children[i].children)>0{
						currentNode.children[i].isWord = false
						return
					}

					copy(currentNode.children[i:],currentNode.children[i+1:])
					currentNode.children = currentNode.children[:len(currentNode.children)-1]
					return
				}
			}
			word = word[prefixEndI:]
			currentNode = currentNode.children[i]
			continue walk
		}
		return
	}
}

func (n *node) Search(word string) (found bool) {
	if len(word)==0{
		found = true
		return
	}

	currentNode := n
	walk:
	for {
		for i:=0; i<len(currentNode.children); i++{
			prefixEndI := prefixEndPosition(word,currentNode.children[i].path)
			if prefixEndI>0{
				if prefixEndI == len(word) {
					found =  currentNode.children[i].isWord
					return
				}
				word = word[prefixEndI:]
				currentNode = currentNode.children[i]
				continue walk
			}
		}
		return
	}
	return
}

func (n *node) findPrefixNode(startByte byte) *node {
	for i:=0;i<len(n.children);i++{
		if n.children[i].path[0]== startByte{
			return n.children[i]
		}
	}
	return nil
}

func lengthCommonPrefix(s1,s2 string) int {
	max := min(len(s1),len(s2))
	i:=0
	for i<max && s1[i]==s2[i] {
		i++
	}
	return i
}

func prefixEndPosition(s,prefix string)  int {
	i:=0
	for ;i<len(prefix);i++{
		if i>=len(s) {
			return -1
		}
		if s[i]!=prefix[i]{
			return -1
		}
	}
	return i
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
