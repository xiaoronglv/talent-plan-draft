package main

import (
	"github.com/pingcap/check"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var _ = check.Suite(&sortTestSuite{})

func TestT(t *testing.T) {
	check.TestingT(t)
}

func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int63()
	}
}

type sortTestSuite struct{}

func (s *sortTestSuite) TestMergeSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for i := range lens {
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		MergeSort(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) TestmergeSort(c *check.C) {
	lens := []int{2, 4, 8, 16, 256, 512, 1024, 1 << 11, 1 << 13, 1 << 19}

	for i := range lens {
		src := make([]int64, lens[i])
		tmp := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		mergeSort(src, tmp)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}
}

func (s *sortTestSuite) Testdivide(c *check.C) {
	nums := []int64{1, 31, 67, 0, 2, 522, 3} // sort a slice that contains 7(odd) elements.
	left, right := divide(nums)
	c.Assert(len(left), check.Equals, 3)
	c.Assert(len(right), check.Equals, 4)

	nums = []int64{2, 67, 0, 2, 522, 3} // sort a slice that contains 6(even) elements.
	left, right = divide(nums)
	c.Assert(len(left), check.Equals, 3)
	c.Assert(len(right), check.Equals, 3)

	nums = []int64{3} // sort a slice that contains one element.
	left, right = divide(nums)
	c.Assert(len(left), check.Equals, 0)
	c.Assert(len(right), check.Equals, 1)
}

func (s *sortTestSuite) Testmerge(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}

	for _, l := range lens {
		// prepare the to be sorted nums
		nums := make([]int64, l)
		prepare(nums)

		// prepare the left and right
		left, right := divide(nums)
		// both left and right should be sorted.
		sort.Slice(left, func(i, j int) bool {
			return left[i] < left[j]
		})
		sort.Slice(right, func(i, j int) bool {
			return right[i] < right[j]
		})

		// prepare the expect
		expect := make([]int64, l)
		copy(expect, nums)
		sort.Slice(expect, func(i, j int) bool {
			return expect[i] < expect[j]
		})

		// prepare the temporary storage to save sorted numbers.
		tmp := make([]int64, l)
		merge(nums, tmp, left, right)
		for i := 0; i < l; i++ {
			c.Assert(nums[i], check.Equals, expect[i])
		}
	}
}
