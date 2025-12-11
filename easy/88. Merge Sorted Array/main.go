package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Print(merge(nums1, 3, nums2, len(nums2)))
	//sort(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for i := m; i < (m + n); i++ {
		nums1[i] = nums2[i-m]
	}
	return sort(nums1)
}

func merge_arr(nums1 []int, m int, nums2 []int, n int) []int {
	res := make([]int, 0, m+n)
	il, ir := 0, 0
	for il < len(nums1) && ir < len(nums2) {
		if nums1[il] <= nums2[ir] {
			res = append(res, nums1[il])
			il++
		} else {
			res = append(res, nums2[ir])
			ir++
		}
	}

	for il < len(nums1) {
		res = append(res, nums1[il])
		il++
	}

	for ir < len(nums2) {
		res = append(res, nums2[ir])
		ir++
	}

	return res
}
func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := sort(arr[:mid])
	right := sort(arr[mid:])
	return merge_arr(left, len(left), right, len(right))
}
