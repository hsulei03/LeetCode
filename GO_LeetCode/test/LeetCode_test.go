package test

import (
	"fmt"
	// "hsulei/leetcode"
	"hsulei/leetcode"
	"testing"
)

func TestLeetCodeFunc(t *testing.T) {

	// Input: nums = [3,2,2,3], val = 3
	nums := []int{3, 2, 2, 3}
	val := 3
	res := leetcode.RemoveElement(nums, val)
	fmt.Println(res)
}
