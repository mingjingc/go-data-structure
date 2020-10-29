package mysort

func MergeSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := len(data) >> 1
	MergeSort(data[:mid])
	MergeSort(data[mid:])

	cpy := make([]int, len(data))
	copy(cpy, data)
	left := 0
	right := mid
	i := 0
	for ; left < mid && right < len(cpy); i++ {
		if cpy[left] > cpy[right] {
			data[i] = cpy[right]
			right++
			continue
		}
		data[i] = cpy[left]
		left++
	}
	for ; left < mid; i++ {
		data[i] = cpy[left]
		left++
	}
	for ; right < len(cpy); i++ {
		data[i] = cpy[right]
		right++
	}
}
