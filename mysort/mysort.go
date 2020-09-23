package mysort

import (
	"reflect"
	"sort"
)

type intSlice []int

func (s intSlice) Len() int {
	return len(s)
}
func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type sliceWrapper struct {
	data interface{}
	less func(i, j int) bool
}

func (s sliceWrapper) Len() int {
	rv := reflect.ValueOf(s.data)
	return rv.Len()
}
func (is sliceWrapper) Less(i, j int) bool {
	return is.less(i, j)
}
func (is sliceWrapper) Swap(i, j int) {
	sw := reflect.Swapper(is.data)
	sw(i, j)
}

func QuickSortSlice(data interface{}, less func(i, j int) bool) {
	sw := sliceWrapper{
		data: data,
		less: less,
	}
	QuickSort(sw)
}

func QuickSortInts(data []int) {
	QuickSort(intSlice(data))
}

func QuickSort(data sort.Interface) {
	quickSort(data, 0, data.Len())
}
func quickSort(data sort.Interface, start, end int) {
	if start >= end-1 {
		return
	}
	pivot := start
	left := start + 1
	right := end - 1
	for left < right {
		for left < right && data.Less(left, pivot) {
			left++
		}
		for left < right && data.Less(pivot, right) {
			right--
		}
		if left < right {
			data.Swap(left, right)
			left++
			right--
		}
	}
	if data.Less(pivot, left) {
		data.Swap(pivot, left-1)
		pivot = left - 1
	} else {
		data.Swap(pivot, left)
		pivot = left
	}
	quickSort(data, start, pivot)
	quickSort(data, pivot+1, end)
}

func InsertionSortInts(data []int) {
	InsertionSort(intSlice(data))
}
func InsertionSort(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(min, i)
	}
}
