package main

import (
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	mergeSort(src)
}

func mergeSort(src []int64) {
	if len(src) <= 1 {
		return
	} else {
		// divide the src and tmp into sub-slices.
		left, right := divide(src)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			mergeSort(left)
		}()

		mergeSort(right)
		wg.Wait()

		// merge two sorted slice into single slice.
		merge(src, left, right)
	}

}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func merge(nums, left, right []int64) {
	tmp := make([]int64, len(nums))
	j, k := 0, 0

	for i, _ := range nums {
		switch {
		case j >= len(left): // when left is empty
			tmp[i] = right[k]
			k++
		case k >= len(right): // when right is empty
			tmp[i] = left[j]
			j++
		case left[j] <= right[k]:
			tmp[i] = left[j]
			j++
		case right[k] < left[j]:
			tmp[i] = right[k]
			k++
		}
	}

	for i, _ := range tmp {
		nums[i] = tmp[i]
	}
}
