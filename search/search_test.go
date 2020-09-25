package search

import "testing"

func TestBinarySearch(t *testing.T) {
	testData := []struct{
		Data []int
		FindValue int
		found bool
	}{
		{
			[]int{1,2,3,4,5,5,7},
			3,
			true,
		}, {
			[]int{1,2,3,4,5,5,6},
			9,
			false,
		}, {
				[]int{1,1,2,3,5,8,13,21,34,55},
				2,
				true,
		}, {
			[]int{1,1,2,3,5,8,13,21,34,55},
			34,
			true,
		},{
			[]int{1,1,2,3,5,8,13,21,34,55},
			-1,
			false,
		},
	}
	for _,test:=range testData{
		if BinarySearch(test.Data,test.FindValue)!=test.found{
			t.Errorf("data %d find %d, expect found %t",test.Data,test.FindValue,test.found)
		}
	}

}
