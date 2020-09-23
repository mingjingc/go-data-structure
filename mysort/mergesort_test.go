package mysort

import "testing"

func TestMergeSort(t *testing.T) {
	data := []int{1,2,99,8,0,12,-3}
	MergeSort(data)
	if !checkIntsInOrder(data){
		t.Error("merge sort failed!")
	}
}