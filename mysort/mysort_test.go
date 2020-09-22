package mysort

import (
	"testing"
)

type Student struct {
	Name string
	Age  int
}

func TestQuickSort(t *testing.T) {

	ints := []int{1, 2, 8, 9, 0, -9, 8}
	QuickSortInts(ints)
	if !checkIntsInOrder(ints) {
		t.Error("quick sort failed!")
	}

	students := []Student{
		{
			Name: "李明",
			Age:  13,
		},
		{
			Name: "小花",
			Age:  15,
		},
		{
			Name: "王明",
			Age:  14,
		},
		{
			Name: "小亮",
			Age:  20,
		},
		{
			Name: "志明",
			Age:  17,
		},
	}
	QuickSortSlice(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	if !checkStudentInOrder(students) {
		t.Error("QuickSortSlice failed!")
	}

	// use less function to implement descending sort
	ints = []int{1, 2, 8, 9, 0, -9, 8}
	QuickSortSlice(ints, func(i, j int) bool {
		return ints[i] > ints[j]
	})
}

func TestInsertionSort(t *testing.T) {
	ints := []int{1, 2, 8, 9, 0, -9, 8}
	InsertionSortInts(ints)
	if !checkIntsInOrder(ints) {
		t.Error("insertion sort failed!")
	}
}

func checkStudentInOrder(students []Student) bool {
	i := 1
	for ; i < len(students); i++ {
		if students[i].Age < students[i-1].Age {
			return false
		}
	}
	return true
}
func checkIntsInOrder(data []int) bool {
	i := 1
	for ; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
