package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 4, 2, 4, 5, 9, 0}
	BubbleSort(nums)
}
