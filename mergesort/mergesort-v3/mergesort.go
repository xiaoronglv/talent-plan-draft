package main

import (
	"sort"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	// Slice consists a pointer to array, length, capacity. The internal array
	// provide the power to store values. When cutting a slice (e.g. s[0:7]) and
	// assign the value to a new variable, golang does not copy the internal
	// array, instead it will reuse the original arry.

	// for example
	// ```golang
	// tmp := []int64{3, 4, 5, 6}
	// left, right := divide(tmp)
	//```

	// left, right are brand new slice, but they points to identical array,
	// which belongs to tmp. we can use this feature to avoid unnecessary memory
	// allocation at each level of merge sort.
	temp := make([]int64, len(src))
	mergeSort(src, temp)
}

// Implement an internal function to sort a slice of int64.
// 1. divide the slice into two slices, which are of same size.
// 2. Recursive sort both slices.
// 3. Merge the two sorted slices.
func mergeSort(src, tmp []int64) {
	// An 1-element slice could be treated as `already sorted`, which
	// alsosupposed to be the stopping condition for standard merge sort.
	// Unfortunately, this kind of stop condition  brings an disadvantages. It
	// will spawn more and more goroutine with the input size growing.
	//
	// To avoid this kind of performance issue, we change the stopping condition
	// when the input slice is less than a threshold, quicksort will be choosen,
	// to avoid unnecessary goroutine.
	if len(src) < 1000000 {
		sort.Slice(src, func(i, j int) bool {
			return src[i] < src[j]
		})
	} else {
		// divide the src and tmp into sub-slices.
		left, right := divide(src)
		lTmp, rTmp := divide(tmp)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			mergeSort(left, lTmp)
		}()

		mergeSort(right, rTmp)
		wg.Wait()

		// merge two sorted slice into single slice.
		merge(src, tmp, left, right)
	}

}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func merge(nums, tmp, left, right []int64) {
	// j is for left index and k is for right index of sub-slice to be sorted.
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
