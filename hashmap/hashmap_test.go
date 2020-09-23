package hashmap

import "testing"

func TestHashMap(t *testing.T)  {
	hashMap := NewHashMap()

	hashMap.Put("a",1)
	hashMap.Put("b",1)
	if v,ok:=hashMap.Get("a");!ok || v!=1{
		t.Errorf("key %v should get %v, but got %v","a",1,v)
	}
	hashMap.Delete("a")
	if v,ok:=hashMap.Get("a");ok{
		t.Errorf("key %v had been deleted, but got value %v","a",v)
	}

	if hashMap.Len() !=1 {
		t.Errorf("HashMap length should be %d, but got %d",1,hashMap.Len())
	}
}
