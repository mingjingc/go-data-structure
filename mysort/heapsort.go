package mysort

// heap sort is similar to heap implements  priority queue
func heapSort(data []int) {
	initHeap(data)
	for i := len(data) - 1; i >= 1; i-- {
		data[0], data[i] = data[i], data[0]
		initHeap(data[:i])
	}
}
func initHeap(data []int) {
	n := len(data)
	for i := n / 2; i >= 0; i-- {
		down(data, i, n)
	}
}
func down(data []int, i, n int) {
	for {
		l := i*2 + 1
		if l >= n {
			break
		}
		if r := l + 1; r < n && data[r] > data[l] {
			l = r
		}
		if data[i] >= data[l] {
			break
		}
		data[i], data[l] = data[l], data[i]
		i = l
	}
}
