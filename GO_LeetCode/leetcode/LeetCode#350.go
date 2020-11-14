package leetcode

//Given two arrays, write a function to compute their intersection.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]

func Intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums1 {
		m[v]++
	}

	for _, v2 := range nums2 {
		count, ok := m[v2]
		if ok && count > 0 {
			m[v2]--
			res = append(res, v2)
		}
	}
	return res
}
