package test

import (
	"fmt"
	"hsulei/leetcode"
	"testing"
)

func TestLeetCodeFunc(t *testing.T) {
	// Example 1:
	// Input: s = "III"
	// Output: 3

	// Example 2:
	// Input: s = "IV"
	// Output: 4

	// Example 3:
	// Input: s = "IX"
	// Output: 9

	// Example 4:
	// Input: s = "LVIII"
	// Output: 58
	// Explanation: L = 50, V= 5, III = 3.

	// Example 5:
	// Input: s = "MCMXCIV"
	// Output: 1994
	// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

	example1 := leetcode.RomanToInt("III")
	fmt.Println(example1)
	example2 := leetcode.RomanToInt("IV")
	fmt.Println(example2)
	example3 := leetcode.RomanToInt("IX")
	fmt.Println(example3)
	example4 := leetcode.RomanToInt("LVII")
	fmt.Println(example4)
	example5 := leetcode.RomanToInt("MCMXCIV")
	fmt.Println(example5)

}
