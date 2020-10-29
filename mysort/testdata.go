package mysort

var sortTestData = []struct {
	data   []int
	result []int
}{
	{
		data:   []int{},
		result: []int{},
	}, {
		data:   []int{1},
		result: []int{1},
	}, {
		data:   []int{1, 2, 8, 9, 0, -9, 8},
		result: []int{-9, 0, 1, 2, 8, 8, 9},
	},
}
