package test

import (
	"fmt"
	"hsulei/leetcode"
	"testing"
)

func TestLeetCodeFunc(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := leetcode.Intersect(nums1, nums2)
	fmt.Println(res)
}
