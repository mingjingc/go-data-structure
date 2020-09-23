package bitset

import "testing"

func TestBitSet(t *testing.T)  {
	set := NewBitSet()

	set.Set(8)
	if !set.Get(8){
		t.Errorf("Bit Index %d should be true", 8)
	}

	set.Set(89)
	if !set.Get(89) {
		t.Errorf("Bit Index %d should be true",89)
	}
	if !set.Get(8){
		t.Errorf("Bit Index %d should be true", 8)
	}

	set.Clean(8)
	if set.Get(8){
		t.Errorf("Bit Index %d was clean, but got true",8)
	}

}