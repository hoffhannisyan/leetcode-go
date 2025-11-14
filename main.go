package main

import (
	"fmt"

	"github.com/hoffhannisyan/leetcode/tasks"
)

func main() {

	// --- LeetCode #1: Two Sum ---
	/*
		printTitle(1, "Two Sum")

		nums := []int{2, 7, 11, 15}
		target := 9
		result := tasks.TwoSum(nums, target)
		fmt.Println("Input:", nums, "Target:", target)
		fmt.Println("Output:", result)
	*/

	// --- LeetCode #2: Add Two Numbers ---
	/*
		printTitle(2, "Add Two Numbers")

		l1 := &tasks.ListNode{Val: 2, Next: &tasks.ListNode{Val: 4, Next: &tasks.ListNode{Val: 3}}}
		l2 := &tasks.ListNode{Val: 5, Next: &tasks.ListNode{Val: 6, Next: &tasks.ListNode{Val: 4}}}
		resultList := tasks.AddTwoNumbers(l1, l2)

		fmt.Print("Output: [")
		for node := resultList; node != nil; node = node.Next {
			fmt.Print(node.Val)
			if node.Next != nil {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	*/

	// --- LeetCode #3: Longest Substring Without Repeating Characters ---
	/*
		printTitle(3, "Longest Substring Without Repeating Characters")

		s := "pwwkew"
		length := tasks.LengthOfLongestSubstring(s)
		fmt.Println("Input:", s)
		fmt.Println("Output:", length)
	*/

	// --- LeetCode #4: Median of Two Sorted Arrays ---
	/*
		 printTitle(4, "Median of Two Sorted Arrays")

		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		median := tasks.FindMedianSortedArrays(nums1, nums2)
		fmt.Println("Input:", nums1, nums2)
		fmt.Println("Output:", median)
	*/

	// --- LeetCode #5: Longest Palindromic Substring ---

	// printTitle(5, "Longest Palindromic Substring")

	// s1 := "babad"
	// fmt.Printf("Input: %q -> Output: %q\n", s1, tasks.LongestPalindrome(s1))

	// s2 := "cbbd"
	// fmt.Printf("Input: %q -> Output: %q\n", s2, tasks.LongestPalindrome(s2))

	// --- LeetCode #6: Zigzag Conversion ---

	// printTitle(6, "Zigzag Conversion")

	// s := "PAYPALISHIRING"
	// fmt.Printf("Input: %q, numRows = 3 -> Output: %q\n", s, tasks.Convert(s, 3))
	// fmt.Printf("Input: %q, numRows = 4 -> Output: %q\n", s, tasks.Convert(s, 4))
	// fmt.Printf("Input: %q, numRows = 1 -> Output: %q\n", "A", tasks.Convert("A", 1))

	// --- LeetCode #7: Reverse Integer ---
	// printTitle(7, "Reverse Integer")

	// fmt.Printf("Input: %d -> Output: %d\n", 123, tasks.Reverse(123))
	// fmt.Printf("Input: %d -> Output: %d\n", -123, tasks.Reverse(-123))
	// fmt.Printf("Input: %d -> Output: %d\n", 120, tasks.Reverse(120))
	// fmt.Printf("Input: %d -> Output: %d\n", 1534236469, tasks.Reverse(1534236469)) // overflow -> 0

	// --- LeetCode #8: String to Integer (atoi) ---

	// printTitle(8, "String to Integer (atoi)")

	// fmt.Printf("Input: %q -> Output: %d\n", "42", tasks.MyAtoi("42"))
	// fmt.Printf("Input: %q -> Output: %d\n", "   -042", tasks.MyAtoi("   -042"))
	// fmt.Printf("Input: %q -> Output: %d\n", "1337c0d3", tasks.MyAtoi("1337c0d3"))
	// fmt.Printf("Input: %q -> Output: %d\n", "0-1", tasks.MyAtoi("0-1"))
	// fmt.Printf("Input: %q -> Output: %d\n", "words and 987", tasks.MyAtoi("words and 987"))

	// --- LeetCode #9: Palindrome Number ---
	printTitle(9, "Palindrome Number")

	tests := []int{121, -121, 10, 12321, 0, 1221}

	for _, t := range tests {
		fmt.Printf("Input: %d -> Output: %v\n", t, tasks.IsPalindrome(t))
	}

}

func printTitle(id int, title string) {
	fmt.Printf("\n--- LeetCode #%d: %s ---\n", id, title)
}
