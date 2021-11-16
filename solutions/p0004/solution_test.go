package p0004

import "testing"

func TestSolution(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
