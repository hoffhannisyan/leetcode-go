package tasks

/*
Longest Palindromic Substring
Difficulty: Medium

Given a string s, return the longest palindromic substring in s.

A palindrome is a string that reads the same forward and backward.

Example 1:
  Input: s = "babad"
  Output: "bab"
  Explanation: "aba" is also a valid answer.

Example 2:
  Input: s = "cbbd"
  Output: "bb"

Constraints:
  - 1 <= s.length <= 1000
  - s consists of only digits and English letters.

Approach:
  Expand Around Center — For each index, we expand outward while the substring remains a palindrome.
  There are 2n - 1 possible centers (each character and each gap between characters).
  For each center, we expand to find the maximum-length palindrome.

  Time Complexity:  O(n²)
  Space Complexity: O(1)
*/

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	transformed := make([]byte, len(s)*2+3)
	transformed[0] = '^'
	for i := 0; i < len(s); i++ {
		transformed[i*2+1] = '#'
		transformed[i*2+2] = s[i]
	}
	transformed[len(transformed)-2] = '#'
	transformed[len(transformed)-1] = '$'

	n := len(transformed)
	P := make([]int, n)
	center := 0
	right := 0

	for i := 1; i < n-1; i++ {

		mirror := 2*center - i

		if i < right {
			P[i] = min(right-i, P[mirror])
		}

		for transformed[i+P[i]+1] == transformed[i-P[i]-1] {
			P[i]++
		}

		if i+P[i] > right {
			center = i
			right = i + P[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}
