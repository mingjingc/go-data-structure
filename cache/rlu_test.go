package cache

import "testing"

func TestRLUCache(t *testing.T) {
	rluCache := NewLRUCache(2)
	rluCache.Put(1, 1)
	rluCache.Put(2, 2)
	if value := rluCache.Get(1); value != 1 {
		t.Errorf("key %d, expect value %d, but got %d", 1, 1, value)
	}
	rluCache.Put(3, 3)
	if value := rluCache.Get(2); value != nil {
		t.Errorf("key %d,expect value nil", 2)
	}
	rluCache.Put(1, 4)
	if value := rluCache.Get(1); value != 4 {
		t.Errorf("key %d after cover,value expect %d", 1, 4)
	}
}
