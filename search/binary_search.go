package search

func BinarySearch(data []int, value int) bool {
	left:=0
	right:=len(data)-1
	for left<=right {
		mid := (left+right)>>1
		if data[mid] == value {
			return true
		}
		if data[mid]>value{
			right = mid-1
		}else{
			left = mid+1
		}
	}
	return false
}