package list

import "testing"

func TestSkipList(t *testing.T)  {
	skipList:=NewSkipList()
	skipList.Put(1,1)
	skipList.Put(2,2)
	skipList.Put(3,3)
	if v:=skipList.Find(2);v!=2{
		t.Errorf("SkipList Find key %v, should get %v, but got %v",2,2,v)
	}
	skipList.Delete(3)
	if v:=skipList.Find(3);v!=nil{
		t.Errorf("Key %v had been deleted, find expect nil, but got %v ",3,v)
	}

	skipList.Put(4,4)
	skipList.Put(5,5)
	if v:=skipList.Find(5);v!=5{
		t.Errorf("SkipList Find key %d, should get %v, but got %v",5,5,v)
	}
}