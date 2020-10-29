package mysort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, td := range sortTestData {
		if heapSort(td.data); !reflect.DeepEqual(td.data, td.result) {
			t.Errorf("expect %d, but got %d", td.result, td.data)
		}
	}
}
