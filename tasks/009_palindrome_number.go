package tasks

/*
LeetCode #9 — Palindrome Number
Difficulty: Easy

Given an integer x, return true if x is a palindrome, and false otherwise.

A palindrome number reads the same forward and backward.

Examples:

	Input: x = 121
	Output: true
	Explanation: 121 reads as 121 from left to right and from right to left.

	Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-.
	             Therefore it is not a palindrome.

	Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:

	-2^31 <= x <= 2^31 - 1

Notes:
  - Negative numbers are never palindromes due to the '-' sign.
  - To avoid overflow, reverse only half of the digits.
*/
func IsPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0

	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}
