package list

import "math/rand"

type skipListNode struct {
	key int
	value interface{}

	levelsNext []*skipListNode
}

const MaxLevel = 5
type SkipList struct {
	head *skipListNode
	length int
	level int
}

func NewSkipList() *SkipList  {
	return &SkipList{
		head:   &skipListNode{
			key:        0,
			value:      nil,
			levelsNext: make([]*skipListNode,MaxLevel),
		},
		length: 0,
		level:  0,
	}
}

func (sl *SkipList) Put(key int, value interface{})  {
	level := sl.randomLevel()
	newNode:= &skipListNode{
		key:        key,
		value:      value,
		levelsNext: make([]*skipListNode,level),
	}

	p:=sl.head
	update:=make([]*skipListNode,level)
	for i:=level-1;i>=0;i--{
		 for p.levelsNext[i]!=nil && p.levelsNext[i].key<key{
		 	p = p.levelsNext[i]
		 }
		 if p.levelsNext[i]!=nil && p.levelsNext[i].key==key{
		 	p.levelsNext[i].value = value
			 return
		 }
		update[i] = p
	}
	for i,p:=range update{
		newNode.levelsNext[i] = p.levelsNext[i]
		p.levelsNext[i] = newNode
	}
	if level>sl.level{
		sl.level = level
	}
}

func (sl *SkipList) Find(key int) interface{} {
	p:=sl.head
	for i:=sl.level-1;i>=0;i--{
		for p.levelsNext[i]!=nil && p.levelsNext[i].key<key{
			p = p.levelsNext[i]
		}
		if p.levelsNext[i]!=nil && p.levelsNext[i].key == key{
			return p.levelsNext[i].value
		}
	}
	return nil
}

func (sl *SkipList) Delete(key int)  {
	p:=sl.head

	var node *skipListNode
	update:=make([]*skipListNode,sl.level)
	for i:=sl.level-1;i>=0;i--{
		for p.levelsNext[i]!=nil && p.levelsNext[i].key<key{
			p = p.levelsNext[i]
		}
		if p.levelsNext[i]!=nil && p.levelsNext[i].key == key{
			node = p.levelsNext[i]
		}
		update[i] = p
	}

	if node == nil{
		return
	}
	for i:=0;i<len(node.levelsNext);i++{
		if update[i] == sl.head && node.levelsNext[i] == nil{
			sl.level--
		}
		update[i].levelsNext[i]=node.levelsNext[i]
		node.levelsNext[i] = nil

	}
}

func (sl *SkipList) randomLevel() int {
	level := rand.Intn(MaxLevel)
	return level+1
}