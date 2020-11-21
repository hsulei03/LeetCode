package leetcode

import (
	"strings"
)

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

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

func RomanToInt(s string) int {
	symbolMap := make(map[string]int)
	symbolMap["I"] = 1
	symbolMap["V"] = 5
	symbolMap["X"] = 10
	symbolMap["L"] = 50
	symbolMap["C"] = 100
	symbolMap["D"] = 500
	symbolMap["M"] = 1000
	input := strings.Split(s, "")

	var theSymbolValue []int

	for _, theSymbol := range input {
		value, ok := symbolMap[theSymbol]
		if ok {
			theSymbolValue = append(theSymbolValue, value)
		}
	}
	var result int = 0
	for len(theSymbolValue) > 0 {
		if len(theSymbolValue) == 1 {
			result += theSymbolValue[0]
			break
		}
		x := theSymbolValue[0]
		y := theSymbolValue[1]
		value := 0
		if x < y && (y/x == 5 || y/x == 10) {
			value = y - x
			theSymbolValue = theSymbolValue[2:]
		} else {
			value = x
			theSymbolValue = theSymbolValue[1:]
		}
		result += value
	}

	return result
}
