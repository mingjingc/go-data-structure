package hashmap

import (
	"fmt"
)

type entry struct {
	key   interface{}
	value interface{}
	next  *entry
}

type HashMap struct {
	data []*entry

	capacity int
	length int
}

func NewHashMap() *HashMap {
	return &HashMap{
		data:     make([]*entry,256),
		capacity: 256,
		length:   0,
	}
}

func (h *HashMap) Put(key,value interface{})  {
	i :=  h.calculatePosition(key)
	e:=h.data[i]

	var pre *entry
	for e!=nil && e.key !=key {
		pre = e
		e = e.next
	}
	if e==nil {
		 e = &entry{
			key:   key,
			value: value,
			next:  nil,
		}
		if pre == nil{
			h.data[i] = e
		} else {
			pre.next = e
		}
		h.length++
		return
	}
	e.value = value
}
func (h *HashMap) Get(key interface{}) (interface{}, bool) {
	i:=h.calculatePosition(key)
	e:=h.data[i]
	for e!=nil && e.key!=key {
		e = e.next
	}
	if e==nil{
		return nil,false
	}
	return e.value,true
}

func (h *HashMap) Delete(key interface{})  {
	i:=h.calculatePosition(key)
	e := h.data[i]
	var pre *entry
	for e!=nil && e.key!=key {
		pre = e
		e = e.next
	}
	if e!=nil {
		if pre == nil{
			h.data[i] = nil
		}else{
			pre.next = nil
		}
	}
	h.length--
}

func (h *HashMap) Len() int {
	return h.length
}

func (h *HashMap) calculatePosition(key interface{}) uint32  {
	return hash(key)%uint32(h.capacity)
}


func hash(key interface{}) uint32 {
	keyBytes := []byte(fmt.Sprintf("%v",key))
	var h uint32
	for _,b:=range keyBytes {
		h = fnv(h,uint32(b))
	}
	return h
}

// https://tools.ietf.org/html/draft-eastlake-fnv-13
func fnv(a,b uint32) uint32  {
	return a*0x01000193 ^ b
}